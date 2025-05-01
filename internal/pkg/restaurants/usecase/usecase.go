package usecase

import (
	"context"
	"fmt"
	"time"

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

func (u *RestaurantUsecase) GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count int, offset int) (*models.RestaurantFull, error) {
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

func (u *RestaurantUsecase) GetAll(ctx context.Context, count int, offset int) ([]models.Restaurant, error) {
	return u.repo.GetAll(ctx, count, offset)
}

func (u *RestaurantUsecase) GetReviews(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Review, error) {
	return u.repo.GetReviews(ctx, restaurantID, count, offset)
}

func (u *RestaurantUsecase) CreateReview(ctx context.Context, req models.ReviewInReq, id uuid.UUID, restaurantID uuid.UUID, login string) (models.Review, error) {
	newReview := models.Review{
		Id:         uuid.NewV4(),
		User:       login,
		ReviewText: req.ReviewText,
		Rating:     req.Rating,
		CreatedAt:  time.Now(),
	}
	err := u.repo.CreateReviews(ctx, newReview, id, restaurantID)
	if err != nil {
		return models.Review{}, fmt.Errorf("ошибка при получении данных о ресторане: %w", err)
	}

	return newReview, nil
}

func (u *RestaurantUsecase) ReviewExists(ctx context.Context, userID, restaurantID uuid.UUID) (bool, error) {
    return u.repo.ReviewExists(ctx, userID, restaurantID)
}

