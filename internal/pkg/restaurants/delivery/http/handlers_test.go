package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsByRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockRestaurantUsecase(ctrl)
	handler := NewRestaurantHandler(mockUsecase)

	restId := uuid.NewV4()
	restIdStr := restId.String()

	expectedData := &models.RestaurantFull{
		Id:   restId,
		Name: "Grill",
		Tags: []string{"Мясо"},
		Categories: []models.Category{
			{
				Name: "Горячее",
				Products: []models.Product{
					{Id: uuid.NewV4(), Name: "Стейк", Price: 590},
				},
			},
		},
	}

	tests := []struct {
		name           string
		url            string
		varsID         string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:   "Success",
			url:    fmt.Sprintf("/restaurants/%s?count=10&offset=0", restIdStr),
			varsID: restIdStr,
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetProductsByRestaurant(gomock.Any(), restId, 10, 0).
					Return(expectedData, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "Invalid ID",
			url:    "/restaurants/not-a-uuid",
			varsID: "not-a-uuid",
			mockSetup: func() {
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Usecase error",
			url:    fmt.Sprintf("/restaurants/%s", restIdStr),
			varsID: restIdStr,
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetProductsByRestaurant(gomock.Any(), restId, 100, 0).
					Return(nil, errors.New("Usecase error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rec := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/restaurants/{id}", handler.GetProductsByRestaurant)
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if rec.Code == http.StatusOK {
				var decoded models.RestaurantFull
				err := json.Unmarshal(rec.Body.Bytes(), &decoded)
				assert.NoError(t, err)
				assert.Equal(t, expectedData.Name, decoded.Name)
			}
		})
	}
}

func TestRestaurantList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockRestaurantUsecase(ctrl)

	expectedData := []models.Restaurant{
		{
			Id:          uuid.NewV4(),
			Name:        "Grill",
			Description: "Some good nice grill",
			Rating:      5.0,
		},
		{
			Id:          uuid.NewV4(),
			Name:        "Grill",
			Description: "Some good nice grill",
			Rating:      5.0,
		},
	}

	tests := []struct {
		name           string
		url            string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "Success",
			url:  "/restaurants/list?count=10&offset=0",
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetAll(gomock.Any(), 10, 0).
					Return(expectedData, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Usecase error",
			url:  "/restaurants/list?count=10&offset=0",
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetAll(gomock.Any(), 10, 0).
					Return(nil, errors.New("Usecase error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			handler := NewRestaurantHandler(mockUsecase)
			handler.RestaurantList(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if w.Code == http.StatusOK {
				var decoded []models.Restaurant
				err := json.Unmarshal(w.Body.Bytes(), &decoded)
				assert.NoError(t, err)
				for i, _ := range decoded {
					assert.Equal(t, expectedData[i].Name, decoded[i].Name)
				}

			}
		})
	}
}

func TestReviewsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockRestaurantUsecase(ctrl)

	restaurantID := uuid.NewV4()
	expectedReviews := []models.Review{
		{
			Id:         uuid.NewV4(),
			User:       "testuser",
			UserPic:    "/pic.jpg",
			ReviewText: "Great food!",
			Rating:     5,
			CreatedAt:  time.Now(),
		},
	}

	tests := []struct {
		name           string
		url            string
		vars           map[string]string
		mockSetup      func()
		expectedStatus int
		expectBody     bool
	}{
		{
			name: "Success",
			url:  "/restaurants/" + restaurantID.String() + "/reviews?count=5&offset=0",
			vars: map[string]string{
				"id": restaurantID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetReviews(gomock.Any(), restaurantID, 5, 0).
					Return(expectedReviews, nil)
			},
			expectedStatus: http.StatusOK,
			expectBody:     true,
		},
		{
			name: "Invalid restaurant ID",
			url:  "/restaurants/invalid-id/reviews?count=5&offset=0",
			vars: map[string]string{
				"id": "invalid-id",
			},
			mockSetup:      func() {}, // won't be called
			expectedStatus: http.StatusBadRequest,
			expectBody:     false,
		},
		{
			name: "Usecase error",
			url:  "/restaurants/" + restaurantID.String() + "/reviews?count=5&offset=0",
			vars: map[string]string{
				"id": restaurantID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetReviews(gomock.Any(), restaurantID, 5, 0).
					Return(nil, errors.New("something went wrong"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectBody:     false,
		},
		{
			name: "Reviews not found",
			url:  "/restaurants/" + restaurantID.String() + "/reviews?count=5&offset=0",
			vars: map[string]string{
				"id": restaurantID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetReviews(gomock.Any(), restaurantID, 5, 0).
					Return(nil, nil)
			},
			expectedStatus: http.StatusNotFound,
			expectBody:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)
			r = mux.SetURLVars(r, tt.vars)
			w := httptest.NewRecorder()

			handler := NewRestaurantHandler(mockUsecase)
			handler.ReviewsList(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectBody {
				var decoded []models.Review
				err := json.Unmarshal(w.Body.Bytes(), &decoded)
				assert.NoError(t, err)
				assert.Equal(t, len(expectedReviews), len(decoded))
				assert.Equal(t, expectedReviews[0].User, decoded[0].User)
			}
		})
	}
}

func TestRestaurantHandler_CreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockRestaurantUsecase(ctrl)

	tests := []struct {
		name           string
		url            string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:   "Success",
			url:    "/restaurants/1234/reviews", 
			mockSetup: func() {
				mockUsecase.EXPECT().
					CreateReview(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(models.Review{
						Id:         uuid.NewV4(),
						User:       "user1",
						UserPic:    "/path/to/pic",
						ReviewText: "Great food!",
						Rating:     5,
						CreatedAt:  time.Now(),
					}, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "Invalid restaurant ID",
			url:    "/restaurants/invalid/reviews", 
			mockSetup: func() {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Invalid rating",
			url:    "/restaurants/1234/reviews",
			mockSetup: func() {
				mockUsecase.EXPECT().CreateReview(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(models.Review{}, errors.New("invalid rating"))
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "JWT not provided",
			url:    "/restaurants/1234/reviews",
			mockSetup: func() {},
			expectedStatus: http.StatusUnauthorized, 
		},
		{
			name:   "User already reviewed",
			url:    "/restaurants/1234/reviews",
			mockSetup: func() {
				mockUsecase.EXPECT().ReviewExists(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			r := httptest.NewRequest(http.MethodPost, tt.url, nil)
			r.Header.Add("Content-Type", "application/json")
			
			if tt.name != "JWT not provided" {
				r.Header.Add("Authorization", "Bearer valid-jwt-token")
			}

			w := httptest.NewRecorder()

			handler := NewRestaurantHandler(mockUsecase)
			handler.CreateReview(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}



func TestRestaurantHandler_CheckReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockRestaurantUsecase(ctrl)

	tests := []struct {
		name           string
		url            string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name: "Success",
			url:  "/restaurants/1234/reviews",
			mockSetup: func() {
				mockUsecase.EXPECT().ReviewExistsReturn(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(models.ReviewUser{Id: uuid.NewV4()}, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid restaurant ID",
			url:            "/restaurants/invalid/reviews",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "JWT not provided",
			url:            "/restaurants/1234/reviews",
			mockSetup:      func() {},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Error in review check",
			url:  "/restaurants/1234/reviews",
			mockSetup: func() {
				mockUsecase.EXPECT().ReviewExistsReturn(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(models.ReviewUser{}, errors.New("check error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			handler := NewRestaurantHandler(mockUsecase)
			handler.CheckReviews(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
