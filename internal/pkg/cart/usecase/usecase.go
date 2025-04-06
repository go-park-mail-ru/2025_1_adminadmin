package usecase

import (
	"context"

	repo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
)

type CartUsecase struct {
	cartRepo *repo.CartRepository
}

func NewCartUsecase(cartRepo *repo.CartRepository) *CartUsecase {
	return &CartUsecase{cartRepo: cartRepo}
}

func (uc *CartUsecase) GetCart(ctx context.Context, userID string) (map[string]int, error) {
	return uc.cartRepo.GetCart(ctx, userID)
}

func (uc *CartUsecase) AddItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.AddItem(ctx, userID, productID)
}

func (uc *CartUsecase) RemoveItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.RemoveItem(ctx, userID, productID)
}
