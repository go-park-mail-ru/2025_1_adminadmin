package pg

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	getFieldProduct   = "SELECT id, name, price, image_url, weight FROM products WHERE id = ANY($1)"
	getRestaurantName = "SELECT name FROM restaurants WHERE id = $1"
	insertOrder       = `INSERT INTO orders (id, user_id, status, address_id, order_products,
		apartment_or_office, intercom, entrance, floor,
		courier_comment, leave_at_door, created_at, final_price) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	getAllOrders = `SELECT
    id,
    user_id,
    status,
    address_id,
    order_products,
    apartment_or_office,
    intercom,
    entrance,
    floor,
    courier_comment,
    leave_at_door,
    final_price,
    created_at
FROM orders WHERE user_id = $1 LIMIT $2 OFFSET $3;`
	getOrderById = `SELECT
    id,
    user_id,
    status,
    address_id,
    order_products,
    apartment_or_office,
    intercom,
    entrance,
    floor,
    courier_comment,
    leave_at_door,
    final_price,
    created_at
FROM orders WHERE id = $1 AND user_id = $2;`
	updateOrderStatus = `UPDATE orders SET status = $1 WHERE id = $2;`
	scheduleDeliveryStatusChange = `SELECT cron.schedule_in('20 seconds', $$UPDATE orders SET status = 'in delivery' WHERE id = $1$$);`
)

type RestaurantRepository struct {
	db pgxtype.Querier
}

func NewRestaurantRepository(db pgxtype.Querier) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := r.db.Query(ctx, getFieldProduct, productIDs)
	if err != nil {
		logger.Error("Ошибка при выполнении запроса", slog.String("error", err.Error()))
		return models.Cart{}, err
	}
	defer rows.Close()

	var items []models.CartItem

	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.ImageURL, &item.Weight)
		if err != nil {
			logger.Error("Ошибка при сканировании строки", slog.String("error", err.Error()))
			return models.Cart{}, err
		}
		item.Amount = productAmounts[item.Id.String()]
		items = append(items, item)
		item.Sanitize()
	}

	var restaurantName string
	err = r.db.QueryRow(ctx, getRestaurantName, restaurantID).Scan(&restaurantName)
	if err != nil {
		logger.Error("Ошибка при получении имени ресторана", slog.String("error", err.Error()))
		return models.Cart{}, fmt.Errorf("не удалось получить имя ресторана: %w %s %s", err, restaurantName, restaurantID)
	}

	uid, err := uuid.FromString(restaurantID)
	if err != nil {
		logger.Error("Ошибка при преобразовании restaurantID в UUID", slog.String("error", err.Error()))
		return models.Cart{}, err
	}

	cart := models.Cart{
		Id:        uid,
		Name:      restaurantName,
		CartItems: items,
	}
	cart.Sanitize()

	logger.Info("Успешно получена корзина", slog.String("restaurant_name", restaurantName), slog.Int("items_count", len(items)))
	return cart, nil
}

func (r *RestaurantRepository) Save(ctx context.Context, order models.Order, userLogin string) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()), slog.String("user_login", userLogin))

	var userID uuid.UUID
	err := r.db.QueryRow(ctx, `SELECT id FROM users WHERE login = $1`, userLogin).Scan(&userID)
	if err != nil {
		logger.Error("Ошибка при поиске пользователя по логину", slog.String("error", err.Error()))
		return fmt.Errorf("не удалось найти пользователя по логину %s: %w", userLogin, err)
	}

	orderProductsStr, err := order.OrderProducts.MarshalJSON()
	if err != nil {
		logger.Error("Ошибка при маршалинге заказанных товаров", slog.String("error", err.Error()))
		return err
	}
	order.Sanitize()

	_, err = r.db.Exec(ctx, insertOrder,
		order.ID, userID, order.Status, order.Address, string(orderProductsStr),
		order.ApartmentOrOffice, order.Intercom, order.Entrance, order.Floor,
		order.CourierComment, order.LeaveAtDoor, order.CreatedAt, order.FinalPrice)

	if err != nil {
		logger.Error("Ошибка при вставке заказа в базу данных", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Заказ успешно сохранен", slog.String("order_id", order.ID.String()))

	return nil
}

func (r *RestaurantRepository) GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := r.db.Query(ctx, getAllOrders, user_id, count, offset)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var orderProductsJSON string
		if err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.Address, &orderProductsJSON,
			&order.ApartmentOrOffice, &order.Intercom, &order.Entrance, &order.Floor, &order.CourierComment,
			&order.LeaveAtDoor, &order.FinalPrice, &order.CreatedAt); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		if err := json.Unmarshal([]byte(orderProductsJSON), &order.OrderProducts); err != nil {
			logger.Error("ошибка анмаршалинга JSON: " + err.Error())
			return nil, err
		}
		order.Sanitize()
		orders = append(orders, order)
	}
	logger.Info("Successful")
	return orders, rows.Err()
}

func (r *RestaurantRepository) GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var order models.Order
	var orderProductsJSON string

	err := r.db.QueryRow(ctx, getOrderById, order_id, user_id).Scan(&order.ID, &order.UserID, &order.Status, &order.Address, &orderProductsJSON,
		&order.ApartmentOrOffice, &order.Intercom, &order.Entrance, &order.Floor, &order.CourierComment,
		&order.LeaveAtDoor, &order.FinalPrice, &order.CreatedAt)
	if err != nil {
		logger.Error("Ошибка при получении заказа", slog.String("error", err.Error()))
		return models.Order{}, fmt.Errorf("не удалось получить заказ: %w", err)
	}

	if err = json.Unmarshal([]byte(orderProductsJSON), &order.OrderProducts); err != nil {
		logger.Error("ошибка анмаршалинга JSON: " + err.Error())
		return models.Order{}, err
	}
	logger.Info("Successful")
	return order, nil
}

func (r *RestaurantRepository) UpdateOrderStatus(ctx context.Context, order_id uuid.UUID, status string) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	res, err := r.db.Exec(ctx, updateOrderStatus, status, order_id)
	if err != nil {
		logger.Error("Ошибка при обновлении статуса заказа", slog.String("error", err.Error()))
		return err
	}
	if rows := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("заказ с id %s не найден", order_id)
	}

	logger.Info("Successful")
	return nil
}

func (r *RestaurantRepository) ScheduleDeliveryStatusChange(ctx context.Context, orderID uuid.UUID) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	query := `SELECT cron.schedule_in('20 seconds', $$SELECT set_order_in_delivery('` + orderID.String() + `')$$);`
	_, err := r.db.Exec(ctx, query)
	if err != nil {
		logger.Error("Ошибка при обновлении статуса заказа", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Successful")
	return nil
}