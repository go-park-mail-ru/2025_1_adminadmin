package repo

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	searchRestaurantWithProducts = `
		SELECT r.id, r.name, r.banner_url, r.address, r.rating, r.rating_count, r.description,
			   p.id, p.name, p.price, p.image_url, p.weight, p.category
		FROM restaurants r
		LEFT JOIN products p ON r.id = p.restaurant_id
		WHERE r.tsvector_column @@ plainto_tsquery('russian', $1) OR p.tsvector_column @@ plainto_tsquery('russian', $1)
		LIMIT 5;
	`

	searchProductsInRestaurant = `
		SELECT id, name, price, image_url, weight, category
		FROM products
		WHERE restaurant_id = $1 AND tsvector_column @@ plainto_tsquery('russian', $2)
		LIMIT 5;
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
	var currentRestaurant *models.RestaurantSearch

	for rows.Next() {
		var restaurant models.RestaurantSearch
		var product models.ProductSearch
		err = rows.Scan(
			&restaurant.ID,
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

		if currentRestaurant == nil || currentRestaurant.ID != restaurant.ID {
			if currentRestaurant != nil {
				restaurants = append(restaurants, *currentRestaurant)
			}
			currentRestaurant = &restaurant
			currentRestaurant.Products = []models.ProductSearch{}
		}
		currentRestaurant.Products = append(currentRestaurant.Products, product)
	}

	if currentRestaurant != nil {
		restaurants = append(restaurants, *currentRestaurant)
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
