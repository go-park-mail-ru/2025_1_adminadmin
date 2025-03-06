package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

//TODO: correct tests

func TestRestaurantList(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		count          string
		offset         string
		expectedCode   int
		expectedLength int
	}{
		{
			name: "Default parameters",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list", nil),
				w: httptest.NewRecorder(),
			},
			count:          "",
			offset:         "",
			expectedCode:   http.StatusOK,
			expectedLength: 10,
		},
		{
			name: "Valid count and offset",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list?count=5&offset=2", nil),
				w: httptest.NewRecorder(),
			},
			count:          "5",
			offset:         "2",
			expectedCode:   http.StatusOK,
			expectedLength: 5,
		},
		{
			name: "Count exceeds total restaurants",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list?count=15", nil),
				w: httptest.NewRecorder(),
			},
			count:          "15",
			offset:         "",
			expectedCode:   http.StatusOK,
			expectedLength: 10,
		},
		{
			name: "Invalid count (negative)",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list?count=-1", nil),
				w: httptest.NewRecorder(),
			},
			count:          "-1",
			offset:         "",
			expectedCode:   http.StatusOK,
			expectedLength: 0,
		},
		{
			name: "Invalid offset (negative)",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list?count=5&offset=-1", nil),
				w: httptest.NewRecorder(),
			},
			count:          "5",
			offset:         "-1",
			expectedCode:   http.StatusOK,
			expectedLength: 0,
		},
		{
			name: "Offset exceeds total restaurants",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/list?count=5&offset=15", nil),
				w: httptest.NewRecorder(),
			},
			count:          "5",
			offset:         "15",
			expectedCode:   http.StatusOK,
			expectedLength: 0,
		},
	}

	ctrl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	defer ctrl.Finish()
	pgxRows := pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).AddRow(uuid.NewV4(), "restaurant #1", "bla bla bla", "some kitchen", 4.5).AddRow(uuid.NewV4(), "restaurant #1", "bla bla bla", "some kitchen", 4.5).ToPgxRows()
	pgxRows.Next()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil).AnyTimes()
	h := Handler{db: mockPool}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			h.RestaurantList(test.args.w, test.args.r)

			if test.args.w.Code != test.expectedCode {
				t.Errorf("unexpected status code: expected %d got %d", test.expectedCode, test.args.w.Code)
			}

			if test.expectedCode == http.StatusOK {
				contentType := test.args.w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Unexpected 'Content-Type' header value: expected 'application/json' got %s", contentType)
				}
			}

			var restaurants []models.Restaurant
			err := json.Unmarshal(test.args.w.Body.Bytes(), &restaurants)
			if err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if len(restaurants) != test.expectedLength {
				t.Errorf("Unexpected number of restaurants: expected %d got %d", test.expectedLength, len(restaurants))
			}
		})
	}
}

func TestRestaurantByID(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	testRestaurant := models.Restaurant{
		Id:          uuid.NewV4(),
		Name:        "Test Restaurant",
		Description: "Test Description",
		Type:        "Test Type",
		Rating:      4.5,
	}
	restaurants = append(restaurants, testRestaurant)

	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedBody models.Restaurant
	}{
		{
			name: "Valid ID",
			args: args{
				r: httptest.NewRequest("GET", fmt.Sprintf("http://localhost:5458/api/restaurants/%s", testRestaurant.Id.String()), nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusOK,
			expectedBody: testRestaurant,
		},
		{
			name: "Non-existent ID",
			args: args{
				r: httptest.NewRequest("GET", fmt.Sprintf("http://localhost:5458/api/restaurants/%s", uuid.NewV4().String()), nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusNotFound,
			expectedBody: models.Restaurant{},
		},
		{
			name: "Invalid ID",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/just_a_random_mess", nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: models.Restaurant{},
		},
	}

	ctrl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	defer ctrl.Finish()
	pgxRows := pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).AddRow(uuid.NewV4(), "restaurant #1", "bla bla bla", "some kitchen", 4.5).ToPgxRows()
	pgxRows.Next()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), 0, 100).Return(pgxRows, nil)
	h := Handler{db: mockPool}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := mux.NewRouter()
			router.HandleFunc("/api/restaurants/{id}", h.RestaurantByID).Methods(http.MethodGet)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, test.args.r)

			if w.Code != test.expectedCode {
				t.Errorf("Unexpected status code: expected %d got %d", test.expectedCode, w.Code)
			}

			if test.expectedCode == http.StatusOK {
				var responseRestaurant models.Restaurant
				err := json.Unmarshal(w.Body.Bytes(), &responseRestaurant)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if responseRestaurant != test.expectedBody {
					t.Errorf("Unexpected restaurant: expected %+v  got %+v", test.expectedBody, responseRestaurant)
				}
			}
		})
	}
}
