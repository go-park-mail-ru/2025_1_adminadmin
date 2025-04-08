package pg

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	getFieldProduct   = "SELECT id, name, price, image_url, weight FROM products WHERE id = ANY($1)"
	getRestaurantName = "SELECT name FROM restaurants WHERE id = $1"
)

type RestaurantRepository struct {
	db *pgxpool.Pool
}

func NewRestaurantRepository(db *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error) {
	rows, err := r.db.Query(ctx, getFieldProduct, productIDs)
	if err != nil {
		return models.Cart{}, err
	}
	defer rows.Close()

	var items []models.CartItem

	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.ImageURL, &item.Weight)
		if err != nil {
			return models.Cart{}, err
		}
		item.Amount = productAmounts[item.Id.String()]
		items = append(items, item)
	}

	var restaurantName string
	err = r.db.QueryRow(ctx, getRestaurantName, restaurantID).Scan(&restaurantName)
	if err != nil {
		return models.Cart{}, fmt.Errorf("не удалось получить имя ресторана: %w", err)
	}

	uid, err := uuid.FromString(restaurantID)
	if err != nil {
		return models.Cart{}, err
	}

	return models.Cart{
		Id:        uid,
		Name:      restaurantName,
		CartItems: items,
	}, nil
}

