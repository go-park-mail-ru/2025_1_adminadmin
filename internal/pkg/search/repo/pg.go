package repo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	searchProductsInRestaurant = ` 
		SELECT id, name, price, image_url, weight, category
FROM products
WHERE restaurant_id = $1 AND tsvector_column @@ plainto_tsquery('ru', $2)
ORDER BY category, name;
	`
	searchRestaurantWithProducts1 = `
	WITH ranked AS (
		SELECT 
			r.id AS restaurant_id,
			r.name,
			r.banner_url,
			r.address, 
			r.rating,
			r.rating_count,
			r.description,
			ts_rank(r.tsvector_column, plainto_tsquery('ru', $1)) AS ts1, 
			ts_rank(p.tsvector_column, plainto_tsquery('ru', $1)) AS ts2
		FROM restaurants r 
		JOIN products p ON r.id = p.restaurant_id
	)
	SELECT DISTINCT restaurant_id, ts1, ts2
	FROM ranked 
	WHERE ts1 >= 0.3 OR ts2 >= 0.3
	ORDER BY ts1 DESC, ts2 DESC
	LIMIT $2 OFFSET $3;`

	searchRestaurantWithProducts2 = `
	SELECT id, name, price, image_url, weight, category, ts_rank(tsvector_column, plainto_tsquery('ru', $1)) AS ts FROM products
	WHERE restaurant_id = $2 ORDER BY ts DESC LIMIT 5;`
)

type SearchRepo struct {
	db pgxtype.Querier
}

func NewSearchRepo(db pgxtype.Querier) *SearchRepo {
	return &SearchRepo{
		db: db,
	}
}

func (r *SearchRepo) SearchRestaurantWithProducts(ctx context.Context, query string, count, offset int) ([]models.RestaurantSearch, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))
	rows, err := r.db.Query(ctx, searchRestaurantWithProducts1, query, count, offset)
	if err != nil {
		logger.Error("Ошибка при выполнении запроса", slog.String("error", err.Error()))
		return nil, fmt.Errorf("error in db.Query: %w", err)
	}
	defer rows.Close()

	var restaurants []models.RestaurantSearch

	for rows.Next() {
		var restaurant models.RestaurantSearch
		var ts1 interface{}
		var ts2 interface{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.BannerURL,
			&restaurant.Address,
			&restaurant.Rating,
			&restaurant.RatingCount,
			&restaurant.Description,
			&ts1,
			&ts2,
		)

		if err != nil {
			logger.Error("Ошибка при сканировании", slog.String("error", err.Error()))
			return nil, fmt.Errorf("error in rows.Scan: %w", err)
		}
		
		products, err := r.db.Query(ctx, searchRestaurantWithProducts2, query, restaurant.ID)
		if err != nil {
			logger.Error("Ошибка при выполнении запроса", slog.String("error", err.Error()))
			return nil, fmt.Errorf("error in db.Query: %w", err)
		}
		defer products.Close()

		for products.Next() {
			var product models.ProductSearch
			var ts interface{}
			err = products.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.ImageURL,
				&product.Weight,
				&product.Category,
				&ts,
			)
			if err != nil {
				return nil, fmt.Errorf("error in rows.Scan: %w", err)
			}
			restaurant.Products = append(restaurant.Products, product)
		}
		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func (r *SearchRepo) SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductCategory, error) {
	rows, err := r.db.Query(ctx, searchProductsInRestaurant, restaurantID, query)
	if err != nil {
		return nil, fmt.Errorf("error in db.Query: %w", err)
	}
	defer rows.Close()

	categoryMap := make(map[string][]models.ProductSearch)

	for rows.Next() {
		var product models.ProductSearch
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.ImageURL,
			&product.Weight,
			&product.Category,
		)
		if err != nil {
			return nil, fmt.Errorf("error in rows.Scan: %w", err)
		}

		categoryMap[product.Category] = append(categoryMap[product.Category], product)
	}

	var productCategories []models.ProductCategory
	for category, products := range categoryMap {
		productCategories = append(productCategories, models.ProductCategory{
			Name:     category,
			Products: products,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return productCategories, nil
}
