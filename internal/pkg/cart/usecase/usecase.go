package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	pgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	redisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
)

type CartUsecase struct {
	cartRepo       *redisRepo.CartRepository
	restaurantRepo *pgRepo.RestaurantRepository
}

func NewCartUsecase(cartRepo *redisRepo.CartRepository, restaurantRepo *pgRepo.RestaurantRepository) *CartUsecase {
	return &CartUsecase{
		cartRepo:       cartRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *CartUsecase) GetCart(ctx context.Context, userID string) (models.Cart, error) {
	cartRaw, restaurantID, err := uc.cartRepo.GetCart(ctx, userID)
	if err != nil {
		return models.Cart{}, err
	}

	productIDs := make([]string, 0, len(cartRaw))
	for id := range cartRaw {
		productIDs = append(productIDs, id)
	}

	items, err := uc.restaurantRepo.GetCartItem(ctx, productIDs, cartRaw, restaurantID)
	if err != nil {
		return models.Cart{}, err
	}

	return items, nil
}

func (uc *CartUsecase) AddItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.AddItem(ctx, userID, productID)
}

func (uc *CartUsecase) UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error {
	return uc.cartRepo.UpdateItemQuantity(ctx, userID, productID, restaurantId, quantity)
}
func (uc *CartUsecase) RemoveItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.RemoveItem(ctx, userID, productID)
}
