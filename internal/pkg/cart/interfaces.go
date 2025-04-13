package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type CartRepo interface {
	GetCart(ctx context.Context, userID string) (map[string]int, error)
	UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error
	ClearCart(ctx context.Context, userID string) error

	Save(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id uuid.UUID) (models.Order, error)
	GetAllByUser(ctx context.Context, userID uuid.UUID) ([]models.Order, error)
	UpdateStatus(ctx context.Context, orderID uuid.UUID, status string) error
}

type CartUsecase interface {
	GetCart(ctx context.Context, userID string) (models.Cart, error, bool)
	UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error
	ClearCart(ctx context.Context, userID string) error

	CreateOrder(ctx context.Context, userID string, addressID uuid.UUID, cart models.Cart, details models.Order) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (models.Order, error)
	GetAllOrdersByUser(ctx context.Context, userID uuid.UUID) ([]models.Order, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status string) error
}

type RestaurantRepo interface {
	GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error)
}
