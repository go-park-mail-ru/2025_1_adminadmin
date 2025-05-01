package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type RestaurantRepo interface {
	GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error)
	GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) (*models.RestaurantFull, error)
	GetReviews(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Review, error) 
	CreateReviews(ctx context.Context, req models.Review, id uuid.UUID, restaurantID uuid.UUID) error
	ReviewExists(ctx context.Context, userID, restaurantID uuid.UUID) (bool, error) 

}

type RestaurantUsecase interface {
	GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error)
	GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) (*models.RestaurantFull, error)
	GetReviews(ctx context.Context, restaurantID uuid.UUID, count, offset int) ([]models.Review, error)
	CreateReview(ctx context.Context, req models.ReviewInReq, id uuid.UUID, restaurantID uuid.UUID, login string) (models.Review, error)
	ReviewExists(ctx context.Context, userID, restaurantID uuid.UUID) (bool, error)

}
