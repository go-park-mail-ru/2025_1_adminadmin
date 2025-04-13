package pg

import (
	"context"
	"fmt"
	"log"

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

func (r *RestaurantRepository) Save(ctx context.Context, order *models.Order, userLogin string) error {
	log.Printf("Запрос на сохранение заказа. Логин пользователя: %s", userLogin)

	var userID uuid.UUID
	err := r.db.QueryRow(ctx, `SELECT id FROM users WHERE login = $1`, userLogin).Scan(&userID)
	if err != nil {
		log.Printf("Ошибка при поиске пользователя по логину %s: %v", userLogin, err)
		return fmt.Errorf("не удалось найти пользователя по логину %s: %w", userLogin, err)
	}

	log.Printf("Найден user_id для логина %s: %s", userLogin, userID)

	query := `INSERT INTO orders (
		id, user_id, status, address_id, order_products,
		apartment_or_office, intercom, entrance, floor,
		courier_comment, leave_at_door, created_at, final_price, card_number
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`

	orderProductsStr, err := order.OrderProducts.MarshalJSON()
	if err != nil {
		log.Printf("Ошибка при маршалинге заказанных товаров: %v", err)
		return err
	}

	log.Printf("Данные заказа: ID: %s, Статус: %s, Адрес: %s", order.ID, order.Status, order.Address)

	_, err = r.db.Exec(ctx, query,
		order.ID, userID, order.Status, order.Address, string(orderProductsStr),
		order.ApartmentOrOffice, order.Intercom, order.Entrance, order.Floor,
		order.CourierComment, order.LeaveAtDoor, order.CreatedAt, order.FinalPrice, order.CardNumber)

	if err != nil {
		log.Printf("Ошибка при вставке заказа в базу данных: %v", err)
		return err
	}

	log.Printf("Заказ с ID %s успешно сохранен", order.ID)

	return nil
}
