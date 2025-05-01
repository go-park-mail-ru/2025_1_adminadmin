package repo

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	getAllRestaurant        = "SELECT id, name, description, rating, banner_url FROM restaurants ORDER BY id ASC LIMIT $1 OFFSET $2;"
	getRestaurantByid       = "SELECT id, name, description, rating FROM restaurants WHERE id = $1;"
	getProductsByRestaurant = "SELECT id, name, banner_url, address, description, rating, rating_count, working_mode_from, working_mode_to, delivery_time_from, delivery_time_to FROM restaurants WHERE id = $1 ORDER BY id ASC;"
	getRestaurantTag        = "SELECT rt.name FROM restaurant_tags rt JOIN restaurant_tags_relations rtr ON rtr.tag_id = rt.id WHERE rtr.restaurant_id = $1 ORDER BY rt.name ASC;"
	getRestaurantProduct    = "SELECT id, name, price, image_url, weight, category FROM products WHERE restaurant_id = $1 ORDER BY category ASC, id ASC LIMIT $2 OFFSET $3;"
	getAllReview            = `SELECT r.id, u.login, u.user_pic, COALESCE(r.review_text, '') as review_text, r.rating, r.created_at
								FROM reviews r
								INNER JOIN users u ON r.user_id = u.id
								WHERE r.restaurant_id = $1 ORDER BY r.created_at DESC, r.id ASC
								LIMIT $2 OFFSET $3;`
	insertReview           = "INSERT INTO reviews (id, user_id, restaurant_id, review_text, rating, created_at) VALUES ($1, $2, $3, NULLIF($4, ''), $5, $6);"
	checkReviewExistsQuery = `SELECT EXISTS(SELECT 1 FROM reviews WHERE user_id = $1 AND restaurant_id = $2);`
	getIdByLogin           = "SELECT id FROM reviews WHERE user_id = $1 AND restaurant_id = $2;"
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
	var categoriesOrder []string

	for prodRows.Next() {
		var p models.Product
		var category string
		err := prodRows.Scan(&p.Id, &p.Name, &p.Price, &p.ImageURL, &p.Weight, &category)
		if err != nil {
			logger.Error("failed to scan product: " + err.Error())
			return nil, err
		}
		p.Sanitize()
		if _, exists := categoryMap[category]; !exists {
			categoriesOrder = append(categoriesOrder, category)
		}
		categoryMap[category] = append(categoryMap[category], p)
	}
	for _, categoryName := range categoriesOrder {
		rest.Categories = append(rest.Categories, models.Category{
			Name:     categoryName,
			Products: categoryMap[categoryName],
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
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "42P01" {
			logger.Warn("Таблица reviews не существует, возвращаем пустой массив.")
			return []models.Review{}, nil
		}
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.Id, &review.User, &review.UserPic, &review.ReviewText, &review.Rating, &review.CreatedAt); err != nil {
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

func (repo *RestaurantRepository) CreateReviews(ctx context.Context, req models.Review, id uuid.UUID, restaurantID uuid.UUID) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, insertReview, req.Id, id, restaurantID, req.ReviewText, req.Rating, req.CreatedAt)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("Successful")
	return nil
}

func (repo *RestaurantRepository) ReviewExists(ctx context.Context, userID, restaurantID uuid.UUID) (bool, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var exists bool
	err := repo.db.QueryRow(ctx, checkReviewExistsQuery, userID, restaurantID).Scan(&exists)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	return exists, nil
}

func (repo *RestaurantRepository) ReviewExistsReturn(ctx context.Context, userID, restaurantID uuid.UUID) (models.ReviewUser, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var exists bool
	err := repo.db.QueryRow(ctx, checkReviewExistsQuery, userID, restaurantID).Scan(&exists)
	if err != nil {
		logger.Error(err.Error())
		return models.ReviewUser{}, err
	}
	if !exists {
		return models.ReviewUser{}, nil
	}

	row := repo.db.QueryRow(ctx, getIdByLogin, userID, restaurantID)

	var rev models.ReviewUser
	err = row.Scan(&rev.Id)
	if err != nil {
		logger.Error("failed to scan user: " + err.Error())
		return models.ReviewUser{}, err
	}

	return rev, nil
}
