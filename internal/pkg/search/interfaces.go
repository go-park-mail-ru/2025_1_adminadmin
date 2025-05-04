package search

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type SearchRepo interface {
	SearchRestaurantWithProducts(ctx context.Context, query string) ([]models.RestaurantSearch, error)
	SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductSearch, error) 


}

type SearchUsecase interface {
	SearchRestaurantWithProducts(ctx context.Context, query string) ([]models.RestaurantSearch, error)
	SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductSearch, error)

}