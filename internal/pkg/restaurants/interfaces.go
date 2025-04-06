package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type RestaurantRepo interface {
	GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Restaurant, error)
	GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) (*models.RestaurantFull, error)
}

type RestaurantUsecase interface {
	GetAll(ctx context.Context, count, offset int) ([]models.Restaurant, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Restaurant, error)
	GetProductsByRestaurant(ctx context.Context, restaurantID uuid.UUID, count, offset int) (*models.RestaurantFull, error)
}
