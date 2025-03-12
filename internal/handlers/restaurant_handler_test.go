package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)


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
		mockRows       *pgxpoolmock.Rows
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
			mockRows: pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).AddRow(uuid.NewV4(), "La Piazza", "Итальянская кухня", "Итальянский", 4.5).
				AddRow(uuid.NewV4(), "Sakura", "Японская кухня", "Японский", 4.7).
				AddRow(uuid.NewV4(), "Steak House", "Лучшие стейки в городе", "Американский", 4.6).
				AddRow(uuid.NewV4(), "Bistro Parisien", "Французская кухня", "Французский", 4.3).
				AddRow(uuid.NewV4(), "Taco Loco", "Мексиканская кухня", "Мексиканский", 4.2).
				AddRow(uuid.NewV4(), "Dragon Wok", "Китайская кухня", "Китайский", 4.4).
				AddRow(uuid.NewV4(), "Berlin Döner", "Настоящий немецкий донер", "Немецкий", 4.1).
				AddRow(uuid.NewV4(), "Kebab King", "Лучший кебаб в городе", "Турецкий", 4.0).
				AddRow(uuid.NewV4(), "Green Garden", "Вегетарианская кухня", "Вегетарианский", 4.8).
				AddRow(uuid.NewV4(), "Sea Breeze", "Свежие морепродукты", "Морепродукты", 4.9),
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
			mockRows: pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).AddRow(uuid.NewV4(), "La Piazza", "Итальянская кухня", "Итальянский", 4.5).
				AddRow(uuid.NewV4(), "Sakura", "Японская кухня", "Японский", 4.7).
				AddRow(uuid.NewV4(), "Steak House", "Лучшие стейки в городе", "Американский", 4.6).
				AddRow(uuid.NewV4(), "Bistro Parisien", "Французская кухня", "Французский", 4.3).
				AddRow(uuid.NewV4(), "Taco Loco", "Мексиканская кухня", "Мексиканский", 4.2),
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
			mockRows: pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).AddRow(uuid.NewV4(), "La Piazza", "Итальянская кухня", "Итальянский", 4.5).
				AddRow(uuid.NewV4(), "Sakura", "Японская кухня", "Японский", 4.7).
				AddRow(uuid.NewV4(), "Steak House", "Лучшие стейки в городе", "Американский", 4.6).
				AddRow(uuid.NewV4(), "Bistro Parisien", "Французская кухня", "Французский", 4.3).
				AddRow(uuid.NewV4(), "Taco Loco", "Мексиканская кухня", "Мексиканский", 4.2).
				AddRow(uuid.NewV4(), "Dragon Wok", "Китайская кухня", "Китайский", 4.4).
				AddRow(uuid.NewV4(), "Berlin Döner", "Настоящий немецкий донер", "Немецкий", 4.1).
				AddRow(uuid.NewV4(), "Kebab King", "Лучший кебаб в городе", "Турецкий", 4.0).
				AddRow(uuid.NewV4(), "Green Garden", "Вегетарианская кухня", "Вегетарианский", 4.8).
				AddRow(uuid.NewV4(), "Sea Breeze", "Свежие морепродукты", "Морепродукты", 4.9),
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
			mockRows:       pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}),
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
			mockRows:       pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}),
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
			mockRows:       pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

			count, err := strconv.Atoi(test.count)
			if err != nil {
				count = 100
			}
			offset, err := strconv.Atoi(test.offset)
			if err != nil {
				offset = 0
			}

			mockPool.EXPECT().
				Query(gomock.Any(), selectAll, count, offset).
				Return(test.mockRows.ToPgxRows(), nil)

			h := Handler{db: mockPool}

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
			err = json.Unmarshal(test.args.w.Body.Bytes(), &restaurants)
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

	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedBody models.Restaurant
		mockRows     *pgxpoolmock.Rows
	}{
		{
			name: "Valid ID",
			args: args{
				r: httptest.NewRequest("GET", fmt.Sprintf("http://localhost:5458/api/restaurants/%s", testRestaurant.Id.String()), nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusOK,
			expectedBody: testRestaurant,
			mockRows: pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}).
				AddRow(testRestaurant.Id, testRestaurant.Name, testRestaurant.Description, testRestaurant.Type, testRestaurant.Rating),
		},
		{
			name: "Non-existent ID",
			args: args{
				r: httptest.NewRequest("GET", fmt.Sprintf("http://localhost:5458/api/restaurants/%s", uuid.NewV4().String()), nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusNotFound,
			expectedBody: models.Restaurant{},
			mockRows:     pgxpoolmock.NewRows([]string{"id", "name", "description", "type", "rating"}),
		},
		{
			name: "Invalid ID",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/restaurants/just_a_random_mess", nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: models.Restaurant{},
			mockRows:     nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

			if test.mockRows != nil {
				mockPool.EXPECT().
					Query(gomock.Any(), selectById, gomock.Any()).
					Return(test.mockRows.ToPgxRows(), nil)
			}

			h := Handler{db: mockPool}

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
					t.Errorf("Unexpected restaurant: expected %+v got %+v", test.expectedBody, responseRestaurant)
				}
			}
		})
	}
}
