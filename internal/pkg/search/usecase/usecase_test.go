package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/mocks"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/usecase"
)

func TestSearchRestaurantWithProducts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSearchRepo(ctrl)
	usecase := usecase.NewSearchUsecase(mockRepo)

	ctx := context.Background()
	query := "pizza"
	count := 10
	offset := 0

	expectedRestaurants := []models.RestaurantSearch{
		{Name: "Pizza Place", Description: "Best pizza in town"},
	}

	// Setup mock behavior
	mockRepo.EXPECT().SearchRestaurantWithProducts(ctx, query, count, offset).
		Return(expectedRestaurants, nil)

	// Test the function
	restaurants, err := usecase.SearchRestaurantWithProducts(ctx, query, count, offset)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedRestaurants, restaurants)
}

func TestSearchRestaurantWithProducts_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSearchRepo(ctrl)
	usecase := usecase.NewSearchUsecase(mockRepo)

	ctx := context.Background()
	query := "burger"
	count := 10
	offset := 0

	// Setup mock to return an error
	mockRepo.EXPECT().SearchRestaurantWithProducts(ctx, query, count, offset).
		Return(nil, errors.New("search error"))

	// Test the function
	restaurants, err := usecase.SearchRestaurantWithProducts(ctx, query, count, offset)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, restaurants)
	assert.Contains(t, err.Error(), "error in SearchRestaurantWithProducts")
}


func TestSearchProductsInRestaurant_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSearchRepo(ctrl)
	usecase := usecase.NewSearchUsecase(mockRepo)

	ctx := context.Background()
	restaurantID := uuid.NewV4()
	query := "spicy"

	// Setup mock to return an error
	mockRepo.EXPECT().SearchProductsInRestaurant(ctx, restaurantID, query).
		Return(nil, errors.New("product search error"))

	// Test the function
	productCategories, err := usecase.SearchProductsInRestaurant(ctx, restaurantID, query)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, productCategories)
	assert.Contains(t, err.Error(), "error in SearchProductsInRestaurant")
}
