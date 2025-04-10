package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
)

type CartRepo interface {
	GetCart(ctx context.Context, userID string) (map[string]int, error)
	UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error
}

type CartUsecase interface {
	GetCart(ctx context.Context, userID string) (models.Cart, error)
	UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error
}

type RestaurantRepo interface {
	GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error)
}
