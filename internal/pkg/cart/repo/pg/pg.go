package pg

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	dbUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/db"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/lib/pq"
	"github.com/satori/uuid"
)

const (
	getFieldProduct   = "SELECT id, name, price, image_url, weight FROM products WHERE id = ANY($1)"
	getRestaurantName = "SELECT name FROM restaurants WHERE id = $1"
	insertOrder       = `INSERT INTO orders (id, user_id, status, address_id, order_products,
		apartment_or_office, intercom, entrance, floor,
		courier_comment, leave_at_door, created_at, final_price, order_items) 
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
	countOrdersQuery = `SELECT COUNT(*) FROM orders WHERE user_id = $1;`

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
	updateOrderStatus            = `UPDATE orders SET status = $1 WHERE id = $2;`
	scheduleDeliveryStatusChange = `SELECT cron.schedule_in('20 seconds', $$UPDATE orders SET status = 'in delivery' WHERE id = $1$$);`
)

type RestaurantRepository struct {
	db pgxtype.Querier
}

func NewRestaurantRepository() (*RestaurantRepository, error) {
	db, err := dbUtils.InitDB()
	return &RestaurantRepository{db: db}, err
}

func (r *RestaurantRepository) GetProductPrice(ctx context.Context, productID string) (float64, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var price float64
	query := `SELECT price FROM products WHERE id = $1`
	err := r.db.QueryRow(ctx, query, productID).Scan(&price)
	if err != nil {
		logger.Error("Ошибка получения цены товара: ", slog.String("error", err.Error()))
		return 0, err
	}
	return price, nil
}

func (r *RestaurantRepository) GetRecommendedProducts(ctx context.Context, productIDs []string, restaurantID string) ([]models.CartItem, error) {
	// Преобразуем строковые ID в UUID[]
	var ids []uuid.UUID
	for _, pid := range productIDs {
		id, err := uuid.FromString(pid)
		if err != nil {
			return nil, fmt.Errorf("invalid product ID: %s", pid)
		}
		ids = append(ids, id)
	}

	restID, err := uuid.FromString(restaurantID)
	if err != nil {
		return nil, fmt.Errorf("invalid restaurant ID: %s", restaurantID)
	}

	rows, err := r.db.Query(ctx, `
        WITH current_cart AS (
            SELECT unnest($1::UUID[]) AS product_id
        ),
        orders_with_cart_items AS (
            SELECT o.id AS order_id
            FROM orders o
            JOIN current_cart cc ON cc.product_id = ANY(o.order_items)
        ),
        related_products AS (
            SELECT p.*
            FROM orders_with_cart_items owci
            JOIN orders o ON o.id = owci.order_id
            JOIN products p ON p.id = ANY(o.order_items)
            WHERE NOT EXISTS (
                SELECT 1
                FROM current_cart cc
                WHERE cc.product_id = p.id
            )
            AND p.restaurant_id = $2
        )
        SELECT 
            rp.id,
            rp.name,
            rp.price,
            rp.image_url,
            rp.weight
        FROM related_products rp
        GROUP BY rp.id, rp.name, rp.price, rp.image_url, rp.weight
        ORDER BY COUNT(*) DESC
        LIMIT 5;
    `, pq.Array(ids), restID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.CartItem
	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&item.ImageURL,
			&item.Weight,
		)
		if err != nil {
			return nil, err
		}
		item.Amount = 1 // можно не возвращать
		result = append(result, item)
	}

	return result, nil
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

	var ids []uuid.UUID
	for _, item := range order.OrderProducts.CartItems {
		ids = append(ids, item.Id)
	}

	_, err = r.db.Exec(ctx, insertOrder,
		order.ID, userID, order.Status, order.Address, string(orderProductsStr),
		order.ApartmentOrOffice, order.Intercom, order.Entrance, order.Floor,
		order.CourierComment, order.LeaveAtDoor, order.CreatedAt, order.FinalPrice, ids)

	if err != nil {
		logger.Error("Ошибка при вставке заказа в базу данных", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Заказ успешно сохранен", slog.String("order_id", order.ID.String()))

	return nil
}

func (r *RestaurantRepository) GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, int, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var totalCount int
	err := r.db.QueryRow(ctx, countOrdersQuery, user_id).Scan(&totalCount)
	if err != nil {
		logger.Error("Ошибка при получении общего количества заказов: " + err.Error())
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, getAllOrders, user_id, count, offset)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
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
			return nil, 0, err
		}
		if err := json.Unmarshal([]byte(orderProductsJSON), &order.OrderProducts); err != nil {
			logger.Error("ошибка анмаршалинга JSON: " + err.Error())
			return nil, 0, err
		}
		order.Sanitize()
		orders = append(orders, order)
	}
	logger.Info("Successful")
	return orders, totalCount, rows.Err()
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

	// Безопасный параметризованный запрос
	query := `
        SELECT cron.schedule(
            'delivery_status_' || $1,
            '20 seconds',
            'SELECT set_order_in_delivery($1)'
        )
    `

	_, err := r.db.Exec(ctx, query, orderID)
	if err != nil {
		logger.Error("Failed to schedule delivery status update",
			slog.String("error", err.Error()),
		)
		return fmt.Errorf("failed to schedule delivery update: %w", err)
	}

	logger.Info("Delivery status update scheduled successfully")
	return nil
}
