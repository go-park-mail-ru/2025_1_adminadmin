package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants"
	"github.com/satori/uuid"
)

type RestaurantUsecase struct {
	repo interfaces.RestaurantRepo
}

func NewRestaurantsUsecase(r interfaces.RestaurantRepo) *RestaurantUsecase {
	return &RestaurantUsecase{repo: r}
}

func (u *RestaurantUsecase) GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Product, error) {
	return u.repo.GetProductsByRestaurant(ctx, restaurantID, count, offset)
}

func (u *RestaurantUsecase) GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error) {
	return u.repo.GetAll(ctx, count, offset)
}

func (u *RestaurantUsecase) GetById(ctx context.Context, id uuid.UUID) (*models.Restaurant, error) {
	return u.repo.GetById(ctx, id)
}
