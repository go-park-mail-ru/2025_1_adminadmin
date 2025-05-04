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
ORDER BY category, name;
	`

	searchRestaurantWithProducts = ` 
	WITH matched_restaurants AS (
    -- Шаг 1: находим рестораны, которые соответствуют запросу по названию или продуктам
    SELECT r.id, 1 AS priority
    FROM restaurants r
    WHERE r.tsvector_column @@ plainto_tsquery('ru', $1)
    
    UNION
    
    SELECT r.id, 2 AS priority
    FROM restaurants r
    JOIN products p ON r.id = p.restaurant_id
    WHERE p.tsvector_column @@ plainto_tsquery('ru', $1)
),
products_limited AS (
    -- Шаг 2: находим продукты, которые соответствуют запросу для каждого ресторана
    SELECT *,
           ROW_NUMBER() OVER (PARTITION BY restaurant_id ORDER BY id) AS rn
    FROM products
    WHERE restaurant_id IN (SELECT id FROM matched_restaurants)
      AND tsvector_column @@ plainto_tsquery('ru', $1)
),
fallback_products AS (
    -- Шаг 3: если у ресторана нет подходящих продуктов, возвращаем 5 любых продуктов
    SELECT * FROM products
    WHERE restaurant_id IN (SELECT id FROM matched_restaurants)
    LIMIT 5
),
restaurants_limited AS (
    -- Шаг 4: лимитируем количество ресторанов по запросу
    SELECT id, priority
    FROM matched_restaurants
    LIMIT $2 OFFSET $3  
)
-- Основной запрос: выбираем рестораны и их продукты (или 5 любых продуктов, если нет подходящих)
SELECT 
    r.id, r.name, r.banner_url, r.address, r.rating, r.rating_count, r.description,
    COALESCE(p.id, f.id) AS product_id,
    COALESCE(p.name, f.name) AS product_name,
    COALESCE(p.price, f.price) AS price,
    COALESCE(p.image_url, f.image_url) AS image_url,
    COALESCE(p.weight, f.weight) AS weight,
    COALESCE(p.category, f.category) AS category
FROM restaurants_limited mr
JOIN restaurants r ON r.id = mr.id
LEFT JOIN products_limited p ON r.id = p.restaurant_id AND p.rn <= 5
LEFT JOIN fallback_products f ON r.id = f.restaurant_id
ORDER BY mr.priority ASC, r.rating DESC;

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

func (r *SearchRepo) SearchRestaurantWithProducts(ctx context.Context, query string, count, offset int) ([]models.RestaurantSearch, int, error) {
    rows, err := r.db.Query(ctx, searchRestaurantWithProducts, query, count, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("error in db.Query: %w", err)
    }
    defer rows.Close()

    var restaurants []models.RestaurantSearch
    restaurantMap := make(map[uuid.UUID]*models.RestaurantSearch)

    for rows.Next() {
        var restaurantID uuid.UUID
        var productID uuid.UUID
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
            &productID,
            &product.Name,
            &product.Price,
            &product.ImageURL,
            &product.Weight,
            &product.Category,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("error in rows.Scan: %w", err)
        }

        if _, exists := restaurantMap[restaurantID]; !exists {
            restaurant.ID = restaurantID
            restaurant.Products = []models.ProductSearch{}
            restaurantMap[restaurantID] = &restaurant
        }

        if productID != uuid.Nil {
            product.ID = productID
            restaurantMap[restaurantID].Products = append(restaurantMap[restaurantID].Products, product)
        }
    }

    for _, rest := range restaurantMap {
        restaurants = append(restaurants, *rest)
    }

    var totalCount int
    err = r.db.QueryRow(ctx, "SELECT COUNT(*) FROM restaurants r WHERE r.tsvector_column @@ plainto_tsquery('ru', $1)", query).Scan(&totalCount)
    if err != nil {
        return nil, 0, fmt.Errorf("error in count query: %w", err)
    }

    if err := rows.Err(); err != nil {
        return nil, 0, fmt.Errorf("rows iteration error: %w", err)
    }

    return restaurants, totalCount, nil
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
