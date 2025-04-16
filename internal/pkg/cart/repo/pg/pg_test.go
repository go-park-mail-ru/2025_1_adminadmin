package pg

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetCartItem(t *testing.T) {
	testProductID := uuid.NewV4()
	testRestaurantID := uuid.NewV4()
	testProductIDs := []string{testProductID.String()}
	testProductAmounts := map[string]int{testProductID.String(): 2}
	testRestaurantName := "Test Restaurant"

	product := models.CartItem{
		Id:       testProductID,
		Name:     "Burger",
		Price:    499,
		ImageURL: "default.png",
		Weight:   250,
		Amount:   2,
	}

	productColumns := []string{"id", "name", "price", "image_url", "weight"}
	restaurantColumn := []string{"name"}

	tests := []struct {
		name           string
		repoMocker     func(*pgxpoolmock.MockPgxPool)
		expectedResult models.Cart
		expectError    bool
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				productRows := pgxpoolmock.NewRows(productColumns).
					AddRow(product.Id, product.Name, product.Price, product.ImageURL, product.Weight).
					ToPgxRows()

				mockPool.EXPECT().
					Query(gomock.Any(), getFieldProduct, testProductIDs).
					Return(productRows, nil)

				restaurantRow := pgxpoolmock.NewRows(restaurantColumn).
					AddRow(testRestaurantName).ToPgxRows()
				restaurantRow.Next()

				mockPool.EXPECT().
					QueryRow(gomock.Any(), getRestaurantName, testRestaurantID.String()).
					Return(restaurantRow)
			},
			expectedResult: models.Cart{
				Id:   testRestaurantID,
				Name: testRestaurantName,
				CartItems: []models.CartItem{
					product,
				},
			},
			expectError: false,
		},
		{
			name: "Product query error",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().
					Query(gomock.Any(), getFieldProduct, testProductIDs).
					Return(nil, fmt.Errorf("db error"))
			},
			expectedResult: models.Cart{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			tt.repoMocker(mockPool)

			repo := &RestaurantRepository{db: mockPool}

			result, err := repo.GetCartItem(context.Background(), testProductIDs, testProductAmounts, testRestaurantID.String())
			t.Logf("%+v", result)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

type mockRow struct {
	scanFunc func(dest ...interface{}) error
}

func (m *mockRow) Scan(dest ...interface{}) error {
	return m.scanFunc(dest...)
}

func TestSaveOrder(t *testing.T) {
	testOrderID := uuid.NewV4()
	testUserLogin := "test_user"
	testUserID := uuid.NewV4()
	testOrder := models.Order{
		ID:     testOrderID,
		Status: "new",
		Address: "123 Test St",
		OrderProducts: models.Cart{
			Id:   uuid.NewV4(),
			Name: "Test Restaurant",
			CartItems: []models.CartItem{
				{
					Id:       uuid.NewV4(),
					Name:     "Test Burger",
					Price:    499.99,
					ImageURL: "burger.png",
					Weight:   250,
					Amount:   2,
				},
				{
					Id:       uuid.NewV4(),
					Name:     "Fries",
					Price:    199.49,
					ImageURL: "fries.png",
					Weight:   150,
					Amount:   1,
				},
			},
		},
		ApartmentOrOffice: "12B",
		Intercom:          "123",
		Entrance:          "A",
		Floor:             "3",
		CourierComment:    "Позвоните, когда будете на месте",
		LeaveAtDoor:       true,
		CreatedAt:         time.Now(),
		FinalPrice:        1199.47,
	}
	

	tests := []struct {
		name        string
		mock        func(mockPool *pgxpoolmock.MockPgxPool)
		expectError bool
	}{
		{
			name: "Success",
			mock: func(mockPool *pgxpoolmock.MockPgxPool) {
				userRow := pgxpoolmock.NewRows([]string{"id"}).AddRow(testUserID).ToPgxRows()
				userRow.Next()

				mockPool.EXPECT().
					QueryRow(gomock.Any(), `SELECT id FROM users WHERE login = $1`, testUserLogin).
					Return(userRow)

				mockPool.EXPECT().
					Exec(gomock.Any(), insertOrder,
						testOrder.ID, testUserID, testOrder.Status, testOrder.Address,
						gomock.Any(),
						testOrder.ApartmentOrOffice, testOrder.Intercom, testOrder.Entrance, testOrder.Floor,
						testOrder.CourierComment, testOrder.LeaveAtDoor, testOrder.CreatedAt, testOrder.FinalPrice,
					).
					Return(nil, nil)
			},
			expectError: false,
		},
		{
			name: "User not found",
			mock: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().
					QueryRow(gomock.Any(), `SELECT id FROM users WHERE login = $1`, testUserLogin).
					Return(&mockRow{
						scanFunc: func(dest ...interface{}) error {
							return errors.New("no rows")
						},
					})
			},
			expectError: true,
		},
		{
			name: "Insert fails",
			mock: func(mockPool *pgxpoolmock.MockPgxPool) {
				userRow := pgxpoolmock.NewRows([]string{"id"}).AddRow(testUserID).ToPgxRows()
				userRow.Next()

				mockPool.EXPECT().
					QueryRow(gomock.Any(), `SELECT id FROM users WHERE login = $1`, testUserLogin).
					Return(userRow)

				mockPool.EXPECT().
					Exec(gomock.Any(), insertOrder,
						testOrder.ID, testUserID, testOrder.Status, testOrder.Address,
						gomock.Any(), testOrder.ApartmentOrOffice, testOrder.Intercom,
						testOrder.Entrance, testOrder.Floor, testOrder.CourierComment,
						testOrder.LeaveAtDoor, testOrder.CreatedAt, testOrder.FinalPrice).
					Return(nil, errors.New("insert error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			tt.mock(mockPool)

			repo := &RestaurantRepository{db: mockPool}

			err := repo.Save(context.Background(), testOrder, testUserLogin)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
