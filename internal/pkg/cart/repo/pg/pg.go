package pg

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	getFieldProduct = "SELECT id, name, price, image_url, weight FROM products WHERE id = ANY($1)"
)

type RestaurantRepository struct {
	db *pgxpool.Pool
}

func NewRestaurantRepository(db *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int) ([]models.CartItem, error) {
	rows, err := r.db.Query(ctx, getFieldProduct, productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem

	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.ImageURL, &item.Weight)
		if err != nil {
			return nil, err
		}
		item.Amount = productAmounts[item.ID]
		items = append(items, item)
	}

	return items, nil
}
