package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search/mocks"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
)

func TestSearchHandler_SearchRestaurantWithProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := mocks.NewMockSearchUsecase(ctrl)
	handler := &SearchHandler{uc: mockUC}

	expectedQuery := "pizza"
	expectedCount := 10
	expectedOffset := 0

	expectedRestaurants := []models.RestaurantSearch{
		{
			ID:          uuid.NewV4(),
			Name:        "Pizza Place",
			BannerURL:   "http://example.com/image.jpg",
			Address:     "123 Pizza Street",
			Rating:      4.5,
			RatingCount: 120,
			Description: "Best pizza in town",
			Products: []models.ProductSearch{
				{
					ID:       uuid.NewV4(),
					Name:     "Pepperoni",
					Price:    12.99,
					ImageURL: "http://example.com/pepperoni.jpg",
					Weight:   300,
					Category: "Pizza",
				},
				{
					ID:       uuid.NewV4(),
					Name:     "Margherita",
					Price:    10.99,
					ImageURL: "http://example.com/margherita.jpg",
					Weight:   250,
					Category: "Pizza",
				},
			},
		},
	}

	mockUC.EXPECT().
		SearchRestaurantWithProducts(gomock.Any(), expectedQuery, expectedCount, expectedOffset).
		Return(expectedRestaurants, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/search?query=%s&count=%d&offset=%d", url.QueryEscape(expectedQuery), expectedCount, expectedOffset), nil)
	rec := httptest.NewRecorder()

	handler.SearchRestaurantWithProducts(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var actual []models.RestaurantSearch
	err := json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if !reflect.DeepEqual(expectedRestaurants, actual) {
		t.Errorf("expected %+v, got %+v", expectedRestaurants, actual)
	}
}

func TestSearchHandler_SearchProductsInRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		uc *mocks.MockSearchUsecase
	}

	testRestaurantID := uuid.NewV4()
	testProductID := uuid.NewV4()
	testCategories := []models.ProductCategory{
		{
			Name: "Burgers",
			Products: []models.ProductSearch{
				{
					ID:       testProductID,
					Name:     "Cheeseburger",
					Price:    300,
					ImageURL: "cheeseburger.jpg",
					Weight:   250,
					Category: "Burgers",
				},
			},
		},
	}

	testCases := []struct {
		name           string
		setup          func(f *fields)
		restaurantID   string
		query          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - found products",
			setup: func(f *fields) {
				f.uc.EXPECT().SearchProductsInRestaurant(
					gomock.Any(),
					testRestaurantID,
					"burger",
				).Return(testCategories, nil)
			},
			restaurantID:   testRestaurantID.String(),
			query:          "burger",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"name":"Burgers","products":[{"id":"` + testProductID.String() + `","name":"Cheeseburger","price":300,"image_url":"cheeseburger.jpg","weight":250,"category":"Burgers"}]}]`,
		},
		{
			name: "Empty query",
			setup: func(f *fields) {
				f.uc.EXPECT().SearchProductsInRestaurant(
					gomock.Any(),
					testRestaurantID,
					"",
				).Return([]models.ProductCategory{}, nil)
			},
			restaurantID:   testRestaurantID.String(),
			query:          "",
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
		{
			name: "Search error",
			setup: func(f *fields) {
				f.uc.EXPECT().SearchProductsInRestaurant(
					gomock.Any(),
					testRestaurantID,
					"burger",
				).Return(nil, errors.New("search error"))
			},
			restaurantID:   testRestaurantID.String(),
			query:          "burger",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Ошибка поиска продуктов"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fields := fields{
				uc: mocks.NewMockSearchUsecase(ctrl),
			}
			if tc.setup != nil {
				tc.setup(&fields)
			}

			h := &SearchHandler{
				uc: fields.uc,
			}

			req := httptest.NewRequest("GET", fmt.Sprintf("/restaurants/%s/search?query=%s", tc.restaurantID, url.QueryEscape(tc.query)), nil)
			w := httptest.NewRecorder()

			// Устанавливаем vars для mux
			req = mux.SetURLVars(req, map[string]string{"id": tc.restaurantID})

			h.SearchProductsInRestaurant(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, resp.StatusCode)
			}

			body, _ := io.ReadAll(resp.Body)
			bodyStr := strings.TrimSpace(string(body))

			if tc.expectedStatus == http.StatusOK {
				// Для успешных случаев проверяем точное соответствие JSON
				var expected, actual interface{}
				json.Unmarshal([]byte(tc.expectedBody), &expected)
				json.Unmarshal([]byte(bodyStr), &actual)

				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("expected body %s, got %s", tc.expectedBody, bodyStr)
				}
			} else {
				// Для ошибок проверяем наличие ожидаемого текста
				if !strings.Contains(bodyStr, tc.expectedBody) {
					t.Errorf("expected body to contain %s, got %s", tc.expectedBody, bodyStr)
				}
			}
		})
	}
}
