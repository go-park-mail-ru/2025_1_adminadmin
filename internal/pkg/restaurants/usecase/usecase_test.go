package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/mocks"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
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

func TestGetReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	restaurantID := uuid.NewV4()
	count := 5
	offset := 0
	expectedReviews := []models.Review{
		{Id: uuid.NewV4(), User: "user1", ReviewText: "Excellent!", Rating: 5.0, CreatedAt: time.Now()},
		{Id: uuid.NewV4(), User: "user2", ReviewText: "Good", Rating: 4.0, CreatedAt: time.Now()},
	}

	tests := []struct {
		name        string
		setupMock   func()
		expected    []models.Review
		expectError bool
	}{
		{
			name: "Success",
			setupMock: func() {
				mockRepo.EXPECT().
					GetReviews(gomock.Any(), restaurantID, count, offset).
					Return(expectedReviews, nil)
			},
			expected:    expectedReviews,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					GetReviews(gomock.Any(), restaurantID, count, offset).
					Return(nil, errors.New("fail"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.GetReviews(context.Background(), restaurantID, count, offset)

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

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	restaurantID := uuid.NewV4()
	userID := uuid.NewV4()
	login := "user1"
	req := models.ReviewInReq{
		ReviewText: "Great place!",
		Rating:     5.0,
	}
	expectedReview := models.Review{
		User:       login,
		ReviewText: req.ReviewText,
		Rating:     req.Rating,
	}

	tests := []struct {
		name        string
		setupMock   func()
		expected    models.Review
		expectError bool
	}{
		{
			name: "Success",
			setupMock: func() {
				mockRepo.EXPECT().
					CreateReviews(gomock.Any(), gomock.Any(), userID, restaurantID).
					Return(nil)
			},
			expected:    expectedReview,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					CreateReviews(gomock.Any(), gomock.Any(), userID, restaurantID).
					Return(errors.New("fail"))
			},
			expected:    models.Review{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.CreateReview(context.Background(), req, userID, restaurantID, login)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, models.Review{}, result)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tt.expected.User, result.User)
				assert.Equal(t, tt.expected.ReviewText, result.ReviewText)
				assert.Equal(t, tt.expected.Rating, result.Rating)
			}
		})
	}
}



func TestReviewExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRestaurantRepo(ctrl)
	usecase := NewRestaurantsUsecase(mockRepo)

	restaurantID := uuid.NewV4()
	userID := uuid.NewV4()

	tests := []struct {
		name        string
		setupMock   func()
		expected    bool
		expectError bool
	}{
		{
			name: "Review exists",
			setupMock: func() {
				mockRepo.EXPECT().
					ReviewExists(gomock.Any(), userID, restaurantID).
					Return(true, nil)
			},
			expected:    true,
			expectError: false,
		},
		{
			name: "Review does not exist",
			setupMock: func() {
				mockRepo.EXPECT().
					ReviewExists(gomock.Any(), userID, restaurantID).
					Return(false, nil)
			},
			expected:    false,
			expectError: false,
		},
		{
			name: "Error",
			setupMock: func() {
				mockRepo.EXPECT().
					ReviewExists(gomock.Any(), userID, restaurantID).
					Return(false, errors.New("fail"))
			},
			expected:    false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := usecase.ReviewExists(context.Background(), userID, restaurantID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
