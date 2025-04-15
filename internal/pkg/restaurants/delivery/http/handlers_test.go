package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/satori/uuid"
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
			Id:   uuid.NewV4(),
			Name: "Grill",
			Description: "Some good nice grill",
			Rating: 5.0,
		},
		{
			Id:   uuid.NewV4(),
			Name: "Grill",
			Description: "Some good nice grill",
			Rating: 5.0,
		},
	}
	

	tests := []struct {
		name           string
		url            string
		mockSetup      func()
		expectedStatus int
	}{
		{
			name:   "Success",
			url:    "/restaurants/list?count=10&offset=0",
			mockSetup: func() {
				mockUsecase.EXPECT().
					GetAll(gomock.Any(), 10, 0).
					Return(expectedData, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "Usecase error",
			url:    "/restaurants/list?count=10&offset=0",
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
				for i, _ := range decoded{
					assert.Equal(t, expectedData[i].Name, decoded[i].Name)
				}
				
			}
		})
	}
}
