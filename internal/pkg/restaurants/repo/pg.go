package repo

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	getAllRestaurant  = "SELECT id, name, description, type, rating FROM restaurants LIMIT $1 OFFSET $2;"
	getRestaurantByid = "SELECT id, name, description, type, rating FROM restaurants WHERE id=$1;"
	getProductsByRestaurant = "SELECT id, name, price, image_url, weight, amount FROM products WHERE restaurant_id = $1 LIMIT $2 OFFSET $3;"
)

type RestaurantRepository struct {
	db pgxtype.Querier
}

func NewRestaurantRepository(db pgxtype.Querier) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Product, error) {
	rows, err := r.db.Query(ctx, getProductsByRestaurant, restaurantID, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.ImageURL, &product.Weight, &product.Amount); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, rows.Err()
}

func (r *RestaurantRepository) GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error) {
	rows, err := r.db.Query(ctx, getAllRestaurant, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []models.Restaurant
	for rows.Next() {
		var restaurant models.Restaurant
		if err := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Type, &restaurant.Rating); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	return restaurants, rows.Err()
}

func (r *RestaurantRepository) GetById(ctx context.Context, id uuid.UUID) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.db.QueryRow(ctx, getRestaurantByid, id).
		Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Type, &restaurant.Rating)
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}
