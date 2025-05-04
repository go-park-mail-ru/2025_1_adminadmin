package repo

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	searchProductsInRestaurant = ` 
		SELECT id, name, price, image_url, weight, category
		FROM products
		WHERE restaurant_id = $1 AND tsvector_column @@ plainto_tsquery('ru', $2)
		LIMIT 5;
	`
	searchRestaurantWithProducts = ` 
	WITH matched_restaurants AS (
    SELECT r.id, 1 AS priority
    FROM restaurants r
    WHERE r.tsvector_column @@ plainto_tsquery('ru', $1)
    
    UNION

    SELECT r.id, 2 AS priority
    FROM restaurants r
    JOIN products p ON r.id = p.restaurant_id
    WHERE p.tsvector_column @@ plainto_tsquery('ru', $1)
)
SELECT 
    r.id, r.name, r.banner_url, r.address, r.rating, r.rating_count, r.description,
    p.id, p.name, p.price, p.image_url, p.weight, p.category
FROM matched_restaurants mr
JOIN restaurants r ON r.id = mr.id
LEFT JOIN products p ON r.id = p.restaurant_id
WHERE r.tsvector_column @@ plainto_tsquery('ru', $1)
   OR p.tsvector_column @@ plainto_tsquery('ru', $1)
ORDER BY mr.priority ASC, r.rating DESC
LIMIT 40;
`
)

type SearchRepo struct {
	db pgxtype.Querier
}

func NewSearchRepo(db pgxtype.Querier) *SearchRepo {
	return &SearchRepo{
		db: db,
	}
}

func (r *SearchRepo) SearchRestaurantWithProducts(ctx context.Context, query string) ([]models.RestaurantSearch, error) {
	rows, err := r.db.Query(ctx, searchRestaurantWithProducts, query)
	if err != nil {
		return nil, fmt.Errorf("error in db.Query: %w", err)
	}
	defer rows.Close()

	var restaurants []models.RestaurantSearch
	restaurantMap := make(map[uuid.UUID]*models.RestaurantSearch)

	for rows.Next() {
		var restaurantID uuid.UUID
		var restaurant models.RestaurantSearch
		var product models.ProductSearch

		err = rows.Scan(
			&restaurantID,
			&restaurant.Name,
			&restaurant.BannerURL,
			&restaurant.Address,
			&restaurant.Rating,
			&restaurant.RatingCount,
			&restaurant.Description,
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

		if _, ok := restaurantMap[restaurantID]; !ok {
			restaurant.ID = restaurantID
			restaurant.Products = []models.ProductSearch{}
			restaurantMap[restaurantID] = &restaurant
		}

		if product.ID != uuid.Nil {
			restaurantMap[restaurantID].Products = append(restaurantMap[restaurantID].Products, product)
		}
	}

	for _, rest := range restaurantMap {
		restaurants = append(restaurants, *rest)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return restaurants, nil
}

func (r *SearchRepo) SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductSearch, error) {
	rows, err := r.db.Query(ctx, searchProductsInRestaurant, restaurantID, query)
	if err != nil {
		return nil, fmt.Errorf("error in db.Query: %w", err)
	}
	defer rows.Close()

	var products []models.ProductSearch
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
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return products, nil
}
