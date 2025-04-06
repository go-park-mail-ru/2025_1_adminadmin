package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	redisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
	pgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
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

func (uc *CartUsecase) GetCart(ctx context.Context, userID string) ([]models.CartItem, error) {
	cartRaw, err := uc.cartRepo.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	productIDs := make([]string, 0, len(cartRaw))
	for id := range cartRaw {
		productIDs = append(productIDs, id)
	}

	items, err := uc.restaurantRepo.GetCartItem(ctx, productIDs, cartRaw)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (uc *CartUsecase) AddItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.AddItem(ctx, userID, productID)
}

func (uc *CartUsecase) RemoveItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.RemoveItem(ctx, userID, productID)
}
