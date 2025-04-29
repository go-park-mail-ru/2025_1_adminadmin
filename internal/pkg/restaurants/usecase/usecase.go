package usecase

import (
	"context"
	"fmt"

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

func (u *RestaurantUsecase) GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) (*models.RestaurantFull, error) {
	restaurant, err := u.repo.GetProductsByRestaurant(ctx, restaurantID, count, offset)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении данных о ресторане: %w", err)
	}

	reviews, err := u.repo.GetReviews(ctx, restaurantID, 2, 0)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении отзывов: %w", err)
	}

	restaurant.Reviews = reviews

	return restaurant, nil
}

func (u *RestaurantUsecase) GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error) {
	return u.repo.GetAll(ctx, count, offset)
}

func (u *RestaurantUsecase) GetReviews(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Review, error) {
	return u.repo.GetReviews(ctx, restaurantID, count, offset)
}
