package repo

import (
	"context"
	"errors"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsByRestaurant(t *testing.T) {
	restaurantID := uuid.NewV4()
	columns := []string{
		"id", "name", "banner_url", "address", "description", "rating", "rating_count",
		"working_mode_from", "working_mode_to",
		"delivery_time_from", "delivery_time_to",
	}

	testRestaurant := models.RestaurantFull{
		Id:           restaurantID,
		Name:         "Testaurant",
		BannerURL:    "banner.jpg",
		Address:      "123 Street",
		Description:  "Nice food",
		Rating:       4.5,
		RatingCount:  100,
		WorkingMode:  models.WorkingMode{From: 10, To: 22},
		DeliveryTime: models.DeliveryTime{From: 30, To: 60},
		Tags:         []string{"sushi", "pizza"},
		Categories: []models.Category{
			{
				Name: "Pizza",
				Products: []models.Product{
					{
						Id:       uuid.NewV4(),
						Name:     "Pizza",
						Price:    1000,
						ImageURL: "pizza.jpg",
						Weight:   500,
					},
				},
			},
		},
	}

	tests := []struct {
		name       string
		repoMocker func(*pgxpoolmock.MockPgxPool)
		wantErr    bool
	}{
		{
			name: "Successful",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				restaurantRow := pgxpoolmock.NewRows(columns).
					AddRow(
						testRestaurant.Id,
						testRestaurant.Name,
						testRestaurant.BannerURL,
						testRestaurant.Address,
						testRestaurant.Description,
						testRestaurant.Rating,
						testRestaurant.RatingCount,
						testRestaurant.WorkingMode.From,
						testRestaurant.WorkingMode.To,
						testRestaurant.DeliveryTime.From,
						testRestaurant.DeliveryTime.To,
					).ToPgxRows()
				restaurantRow.Next()

				mockPool.EXPECT().QueryRow(gomock.Any(), getProductsByRestaurant, restaurantID).Return(restaurantRow)

				tagRows := pgxpoolmock.NewRows([]string{"tag"}).
					AddRow("sushi").
					AddRow("pizza").ToPgxRows()

				mockPool.EXPECT().Query(gomock.Any(), getRestaurantTag, restaurantID).Return(tagRows, nil)

				product := testRestaurant.Categories[0].Products[0]
				productRows := pgxpoolmock.NewRows([]string{
					"id", "name", "price", "image_url", "weight", "category",
				}).AddRow(
					product.Id,
					product.Name,
					product.Price,
					product.ImageURL,
					product.Weight,
					"Main",
				).ToPgxRows()

				mockPool.EXPECT().Query(gomock.Any(), getRestaurantProduct, restaurantID, 10, 0).Return(productRows, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			tt.repoMocker(mockPool)

			repo := NewRestaurantRepository(mockPool)
			_, err := repo.GetProductsByRestaurant(context.Background(), restaurantID, 10, 0)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		count  int
		offset int
	}

	columns := []string{"id", "name", "description", "rating"}

	restaurantID := uuid.NewV4()
	expectedRestaurant := models.Restaurant{
		Id:          restaurantID,
		Name:        "Тестовый ресторан",
		Description: "Лучшее место на земле",
		Rating:      4.8,
	}

	tests := []struct {
		name       string
		args       args
		repoMocker  func(mock *pgxpoolmock.MockPgxPool)
		wantResult []models.Restaurant
		wantErr    bool
	}{
		{
			name: "Successful",
			args: args{count: 10, offset: 0},
			repoMocker: func(mock *pgxpoolmock.MockPgxPool) {
				pgxRows := pgxpoolmock.NewRows(columns).
					AddRow(
						expectedRestaurant.Id,
						expectedRestaurant.Name,
						expectedRestaurant.Description,
						expectedRestaurant.Rating,
					).ToPgxRows()

				mock.EXPECT().
					Query(gomock.Any(), getAllRestaurant, 10, 0).
					Return(pgxRows, nil)
			},
			wantResult: []models.Restaurant{expectedRestaurant},
			wantErr:    false,
		},
		{
			name: "Error",
			args: args{count: 10, offset: 0},
			repoMocker: func(mock *pgxpoolmock.MockPgxPool) {
				mock.EXPECT().
					Query(gomock.Any(), getAllRestaurant, 10, 0).
					Return(nil, errors.New("db error"))
			},
			wantResult: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			tt.repoMocker(mockPool)

			repo := NewRestaurantRepository(mockPool)

			got, err := repo.GetAll(context.Background(), tt.args.count, tt.args.offset)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, got)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	columns := []string{"id", "name", "description", "rating"}

	restaurantID := uuid.NewV4()
	expectedRestaurant := models.Restaurant{
		Id:          restaurantID,
		Name:        "Шаурма",
		Description: "Легендарная шаурма",
		Rating:      4.9,
	}

	tests := []struct {
		name        string
		inputID     uuid.UUID
		mockSetup   func(mock *pgxpoolmock.MockPgxPool)
		wantResult  *models.Restaurant
		wantErr     bool
	}{
		{
			name:    "Successful",
			inputID: restaurantID,
			mockSetup: func(mock *pgxpoolmock.MockPgxPool) {
				row := pgxpoolmock.NewRows(columns).
					AddRow(
						expectedRestaurant.Id,
						expectedRestaurant.Name,
						expectedRestaurant.Description,
						expectedRestaurant.Rating,
					).ToPgxRows()
				row.Next()

				mock.EXPECT().
					QueryRow(gomock.Any(), getRestaurantByid, restaurantID).
					Return(row)
			},
			wantResult: &expectedRestaurant,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			tt.mockSetup(mockPool)

			repo := NewRestaurantRepository(mockPool)

			got, err := repo.GetById(context.Background(), tt.inputID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, got)
			}
		})
	}
}

