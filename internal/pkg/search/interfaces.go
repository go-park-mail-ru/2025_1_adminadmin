package search

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type SearchRepo interface {
	SearchRestaurantWithProducts(ctx context.Context, query string, count, offset int) ([]models.RestaurantSearch, int, error)
	SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductCategory, error) 
	}

type SearchUsecase interface {
	SearchRestaurantWithProducts(ctx context.Context, query string, count, offset int) ([]models.RestaurantSearch, int, error)
	SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductCategory, error) 
	}
