package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
)

type CartRepo interface {
	GetCart(ctx context.Context, userID string) (models.CartItem, error)
	AddItem(ctx context.Context, userID, productID string) error
	RemoveItem(ctx context.Context, userID, productID string) error
}

type CartUsecase interface {
	GetCart(ctx context.Context, userID string) (models.CartItem, error)
	AddItem(ctx context.Context, userID, productID string) error
	RemoveItem(ctx context.Context, userID, productID string) error
}

type RestaurantRepo interface {
	GetCartItem(ctx context.Context, userID string) (models.CartItem, error)
}
