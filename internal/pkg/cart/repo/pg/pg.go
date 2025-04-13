package pg

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/satori/uuid"
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

func (r *RestaurantRepository) Save(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders (
		id, user_id, status, address_id, order_products,
		apartment_or_office, intercom, entrance, floor,
		courier_comment, leave_at_door, created_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

	orderProductsStr, err := order.OrderProducts.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query,
		order.ID, order.UserID, order.Status, order.Address, string(orderProductsStr),
		order.ApartmentOrOffice, order.Intercom, order.Entrance, order.Floor,
		order.CourierComment, order.LeaveAtDoor, order.CreatedAt)

	return err
}
