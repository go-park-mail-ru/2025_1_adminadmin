package repo

import "github.com/jackc/pgtype/pgxtype"

const (
	getAllRestaurant        = "SELECT id, name, description, type, rating FROM restaurants LIMIT $1 OFFSET $2;"
	getRestaurantByid       = "SELECT id, name, description, type, rating FROM restaurants WHERE id=$1;"
	getProductsByRestaurant = "SELECT id, name, price, image_url, weight, amount FROM products WHERE restaurant_id = $1 LIMIT $2 OFFSET $3;"
)

type RestaurantRepository struct {
	db pgxtype.Querier
}

func NewRestaurantRepository(db pgxtype.Querier) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}
