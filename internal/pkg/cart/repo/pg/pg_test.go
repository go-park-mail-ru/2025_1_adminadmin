package pg

import (
	"context"
	"fmt"
	"testing"

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
