package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search"
	"github.com/satori/uuid"
)

type SearchUsecase struct {
	repoSearch search.SearchRepo
}

func NewSearchUsecase(repoSearch search.SearchRepo) *SearchUsecase {
	return &SearchUsecase{
		repoSearch: repoSearch,
	}
}

func (uc *SearchUsecase) SearchRestaurantWithProducts(ctx context.Context, query string) ([]models.RestaurantSearch, error) {
	restaurants, err := uc.repoSearch.SearchRestaurantWithProducts(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error in SearchRestaurantWithProducts: %w", err)
	}
	return restaurants, nil
}

func (uc *SearchUsecase) SearchProductsInRestaurant(ctx context.Context, restaurantID uuid.UUID, query string) ([]models.ProductSearch, error) {
	products, err := uc.repoSearch.SearchProductsInRestaurant(ctx, restaurantID, query)
	if err != nil {
		return nil, fmt.Errorf("error in SearchProductsInRestaurant: %w", err)
	}
	return products, nil
}
