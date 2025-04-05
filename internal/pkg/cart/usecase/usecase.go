package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
)

type CartUsecase struct {
	cartRepo *repo.CartRepository
	repo     interfaces.RestaurantRepo
}

func NewCartUsecase(cartRepo *repo.CartRepository, r interfaces.RestaurantRepo) *CartUsecase {
	return &CartUsecase{cartRepo: cartRepo, repo: r}
}

func (uc *CartUsecase) GetCart(ctx context.Context, userID string) (map[string]int, []models.CartItem, error) {
	return uc.cartRepo.GetCart(ctx, userID), uc.repo.GetCartItem(ctx, userID)
}

func (uc *CartUsecase) AddItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.AddItem(ctx, userID, productID)
}

func (uc *CartUsecase) RemoveItem(ctx context.Context, userID, productID string) error {
	return uc.cartRepo.RemoveItem(ctx, userID, productID)
}
