package repo

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/lib/pq"
	"github.com/satori/uuid"
)

const (
	getAllRestaurant        = "SELECT id, name, description, rating, banner_url FROM restaurants LIMIT $1 OFFSET $2;"
	getRestaurantByid       = "SELECT id, name, description, rating FROM restaurants WHERE id = $1;"
	getProductsByRestaurant = "SELECT id, name, banner_url, address, description, rating, rating_count, working_mode_from, working_mode_to, delivery_time_from, delivery_time_to FROM restaurants WHERE id = $1;"
	getRestaurantTag        = "SELECT rt.name FROM restaurant_tags rt JOIN restaurant_tags_relations rtr ON rtr.tag_id = rt.id WHERE rtr.restaurant_id = $1"
	getRestaurantProduct    = "SELECT id, name, price, image_url, weight, category FROM products WHERE restaurant_id = $1 ORDER BY category LIMIT $2 OFFSET $3"
	getAllReview            = `SELECT r.id, u.login, r.review_text, r.rating, r.created_at
								FROM reviews r
								INNER JOIN users u ON r.user_id = u.id
								WHERE r.restaurant_id = $1
								LIMIT $2 OFFSET $3;`
)

type RestaurantRepository struct {
	db pgxtype.Querier
}

func NewRestaurantRepository(db pgxtype.Querier) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count int, offset int) (*models.RestaurantFull, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	row := r.db.QueryRow(ctx, getProductsByRestaurant, restaurantID)

	var rest models.RestaurantFull
	err := row.Scan(
		&rest.Id, &rest.Name, &rest.BannerURL, &rest.Address, &rest.Description, &rest.Rating, &rest.RatingCount,
		&rest.WorkingMode.From, &rest.WorkingMode.To,
		&rest.DeliveryTime.From, &rest.DeliveryTime.To,
	)
	if err != nil {
		logger.Error("failed to scan restaurant: " + err.Error())
		return nil, err
	}
	rest.Sanitize()

	tagRows, err := r.db.Query(ctx, getRestaurantTag, restaurantID)
	if err != nil {
		logger.Error("failed to query tags: " + err.Error())
		return nil, err
	}
	defer tagRows.Close()

	for tagRows.Next() {
		var tag string
		if err := tagRows.Scan(&tag); err != nil {
			logger.Error("failed to scan tag: " + err.Error())
			return nil, err
		}
		rest.Tags = append(rest.Tags, tag)
	}

	prodRows, err := r.db.Query(ctx, getRestaurantProduct, restaurantID, count, offset)
	if err != nil {
		logger.Error("failed to query products: " + err.Error())
		return nil, err
	}
	defer prodRows.Close()

	categoryMap := make(map[string][]models.Product)

	for prodRows.Next() {
		var p models.Product
		var category string
		err := prodRows.Scan(&p.Id, &p.Name, &p.Price, &p.ImageURL, &p.Weight, &category)
		if err != nil {
			logger.Error("failed to scan product: " + err.Error())
			return nil, err
		}
		p.Sanitize()
		categoryMap[category] = append(categoryMap[category], p)
	}

	for categoryName, products := range categoryMap {
		rest.Categories = append(rest.Categories, models.Category{
			Name:     categoryName,
			Products: products,
		})
	}

	logger.Info("Successfully built RestaurantFull model")
	return &rest, nil
}

func (r *RestaurantRepository) GetAll(ctx context.Context, count int, offset int) ([]models.Restaurant, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := r.db.Query(ctx, getAllRestaurant, count, offset)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant
	for rows.Next() {
		var restaurant models.Restaurant
		if err := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Rating, &restaurant.ImageURL); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
		restaurant.Sanitize()
	}

	logger.Info("Successful")
	return restaurants, rows.Err()
}

func (r *RestaurantRepository) GetReviews(ctx context.Context, restaurantID uuid.UUID, count int, offset int) ([]models.Review, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := r.db.Query(ctx, getAllReview, restaurantID, count, offset)
	if err != nil {
		if isTableNotExistError(err) {
			logger.Error("Таблица reviews не существует, возвращаем пустой массив.")
			return []models.Review{}, nil
		}
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.Id, &review.User, &review.ReviewText, &review.Rating, &review.Rating, &review.CreatedAt); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		reviews = append(reviews, review)
		review.Sanitize()
	}

	if len(reviews) == 0 {
		return []models.Review{}, nil
	}

	logger.Info("Successful")
	return reviews, rows.Err()
}

func isTableNotExistError(err error) bool {
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "42P01" {
			return true
		}
	}
	return false
}
