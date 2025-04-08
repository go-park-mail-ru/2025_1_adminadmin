package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
)

type CartRepo interface {
	GetCart(ctx context.Context, userID string) (map[string]int, error)
	AddItem(ctx context.Context, userID, productID string) error
    UpdateItemQuantity(ctx context.Context, userID, productID string, quantity int) error
	RemoveItem(ctx context.Context, userID, productID string) error
	
}

type CartUsecase interface {
	GetCart(ctx context.Context, userID string) ([]models.CartItem, error)
	AddItem(ctx context.Context, userID, productID string) error
    UpdateItemQuantity(ctx context.Context, userID, productID string, quantity int) error
	RemoveItem(ctx context.Context, userID, productID string) error
}

type RestaurantRepo interface {
	GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int) ([]models.CartItem, error)
}