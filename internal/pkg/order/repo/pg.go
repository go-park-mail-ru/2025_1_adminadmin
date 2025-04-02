package repo

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

const (
	getUserOrders = "SELECT id, status, address_id, order_products FROM orders WHERE user_id = $1"
	insertUserOrder = "INSERT INTO orders (id, user_id, status, address_id, order_products) VALUES ($1, $2, $3, $4, $5)"
)

type OrderRepo struct {
	db pgxtype.Querier
}

func CreateOrderRepo(db pgxtype.Querier) *OrderRepo {
	return &OrderRepo{db: db}
}

func (repo *OrderRepo) GetUserOrders(ctx context.Context, user_id uuid.UUID) ([]models.Order, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := repo.db.Query(ctx, getUserOrders, user_id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.Id, &order.Status, &order.AddressId, &order.OrderProducts); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		order.UserId = user_id
		orders = append(orders, order)
	}

	logger.Info("Successful")
	return orders, rows.Err()
}

func (repo *OrderRepo) InsertUserOrder(ctx context.Context, order models.Order) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, insertUserOrder, order.Id, order.UserId, order.Status, order.AddressId, order.OrderProducts)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}
