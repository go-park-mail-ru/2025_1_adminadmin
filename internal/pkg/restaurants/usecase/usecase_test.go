package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/mocks"
	"github.com/satori/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsByRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	restaurantID := uuid.NewV4()
	count := 5
	offset := 0

	expected := &models.RestaurantFull{
		Id:     restaurantID,
		Name:   "Грильница",
		Rating: 4.8,
	}

	tests := []struct {
		name        string
		setupMock   func()
		expected    *models.RestaurantFull
		expectError bool
	}{
		{
			name: "Success",
			setupMock: func() {
				mockRepo.EXPECT().
					GetProductsByRestaurant(gomock.Any(), restaurantID, count, offset).
					Return(expected, nil)
			},
			expected:    expected,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					GetProductsByRestaurant(gomock.Any(), restaurantID, count, offset).
					Return(nil, errors.New("fail"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.GetProductsByRestaurant(context.Background(), restaurantID, count, offset)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestRestaurantUsecase_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	count := 5
	offset := 0

	expected := []models.Restaurant{
		{Id: uuid.NewV4(), Name: "Тануки", Rating: 4.3},
		{Id: uuid.NewV4(), Name: "ЯкиТория", Rating: 4.1},
	}

	tests := []struct {
		name        string
		setupMock   func()
		expected    []models.Restaurant
		expectError bool
	}{
		{
			name: "Success",
			setupMock: func() {
				mockRepo.EXPECT().
					GetAll(gomock.Any(), count, offset).
					Return(expected, nil)
			},
			expected:    expected,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					GetAll(gomock.Any(), count, offset).
					Return(nil, errors.New("db error"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.GetAll(context.Background(), count, offset)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}


func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	restaurantID := uuid.NewV4()
	expectedRestaurant := &models.Restaurant{
		Id:          restaurantID,
		Name:        "KFC",
		Description: "Жарим курочку",
		Rating:      4.2,
	}

	tests := []struct {
		name        string
		setupMock   func()
		expected    *models.Restaurant
		expectError bool
	}{
		{
			name: "Success",
			setupMock: func() {
				mockRepo.EXPECT().
					GetById(gomock.Any(), restaurantID).
					Return(expectedRestaurant, nil)
			},
			expected:    expectedRestaurant,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					GetById(gomock.Any(), restaurantID).
					Return(nil, errors.New("not found"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.GetById(context.Background(), restaurantID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
