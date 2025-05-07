package repo

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/mocks"
)

func TestSearchRestaurantWithProducts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockQuerier(ctrl)
	repo := repo.NewSearchRepo(mockDB)

	ctx := context.Background()
	query := "pizza"
	count := 10
	offset := 0
	restaurantID := uuid.NewV4()

	expectedRestaurants := []models.RestaurantSearch{
		{
			ID:          restaurantID,
			Name:        "Pizza Place",
			BannerURL:   "http://example.com/banner",
			Address:     "123 Pizza St",
			Rating:      4.5,
			RatingCount: 100,
			Description: "Best pizza in town",
		},
	}

	// Setup mock behavior for restaurant search query
	mockDB.EXPECT().Query(ctx, repo.SearchRestaurantWithProducts1, query, count, offset).
		Return([]models.RestaurantSearch{expectedRestaurants[0]}, nil).Times(1)

	// Test the function
	restaurants, err := repo.SearchRestaurantWithProducts(ctx, query, count, offset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedRestaurants, restaurants)
}

func TestSearchRestaurantWithProducts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockQuerier(ctrl)
	repo := repo.NewSearchRepo(mockDB)

	ctx := context.Background()
	query := "burger"
	count := 10
	offset := 0

	// Setup mock to return an error
	mockDB.EXPECT().Query(ctx, repo.SearchRestaurantWithProducts1, query, count, offset).
		Return(nil, errors.New("database error")).Times(1)

	// Test the function
	restaurants, err := repo.SearchRestaurantWithProducts(ctx, query, count, offset)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, restaurants)
	assert.Contains(t, err.Error(), "error in db.Query")
}

func TestSearchProductsInRestaurant_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockQuerier(ctrl)
	repo := repo.NewSearchRepo(mockDB)

	ctx := context.Background()
	restaurantID := uuid.NewV4()
	query := "cheese"

	expectedProductCategories := []models.ProductCategory{
		{
			Name: "Pizza",
			Products: []models.ProductSearch{
				{Name: "Cheese Pizza", Price: 10, ImageURL: "http://example.com/pizza", Weight: 200, Category: "Pizza"},
			},
		},
	}

	// Setup mock behavior for product search query
	mockDB.EXPECT().Query(ctx, repo.SearchProductsInRestaurant, restaurantID, query).
		Return([]models.ProductCategory{expectedProductCategories[0]}, nil).Times(1)

	// Test the function
	productCategories, err := repo.SearchProductsInRestaurant(ctx, restaurantID, query)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedProductCategories, productCategories)
}

func TestSearchProductsInRestaurant_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockQuerier(ctrl)
	repo := repo.NewSearchRepo(mockDB)

	ctx := context.Background()
	restaurantID := uuid.NewV4()
	query := "spicy"

	// Setup mock to return an error
	mockDB.EXPECT().Query(ctx, repo.SearchProductsInRestaurant, restaurantID, query).
		Return(nil, errors.New("database error")).Times(1)

	// Test the function
	productCategories, err := repo.SearchProductsInRestaurant(ctx, restaurantID, query)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, productCategories)
	assert.Contains(t, err.Error(), "error in db.Query")
}
