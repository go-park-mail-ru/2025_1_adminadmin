package cart

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type CartUsecase interface {
	GetCart(ctx context.Context, userID string) (models.Cart, error, bool)
	UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error
	ClearCart(ctx context.Context, userID string) error

	CreateOrder(ctx context.Context, userID string, details models.OrderInReq, cart models.Cart) (models.Order, error)
	GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, int, error)
	GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error)
	UpdateOrderStatus(ctx context.Context, order_id uuid.UUID) error
}

type CartRepo interface {
	GetCart(ctx context.Context, userID string) (map[string]int, string, float64, error)
	UpdateItemQuantity(ctx context.Context, userID, productID, restaurantID string, quantity int, price float64) error
	ClearCart(ctx context.Context, userID string) error
}

type RestaurantRepo interface {
	GetProductPrice(ctx context.Context, productID string) (float64, error)
	GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error)

	Save(ctx context.Context, order models.Order, userLogin string) error
	GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, int, error)
	GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error)
	UpdateOrderStatus(ctx context.Context, order_id uuid.UUID, status string) error
	ScheduleDeliveryStatusChange(ctx context.Context, orderID uuid.UUID) error
}
