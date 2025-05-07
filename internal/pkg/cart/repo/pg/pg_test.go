package pg

import (
	"context"
	"encoding/json"
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

func TestSaveOrder(t *testing.T) {
	testOrderID := uuid.NewV4()
	testUserLogin := "test_user"
	testUserID := uuid.NewV4()
	testOrder := models.Order{
		ID:      testOrderID,
		Status:  "new",
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

func TestGetOrders(t *testing.T) {
    testUserID := uuid.NewV4()
    testOrderID := uuid.NewV4()
    testAddressID := "test_address_123"
    testTime := time.Now().UTC()
    
    // Тестовые данные для корзины
    testCart := models.Cart{
        Id:   uuid.NewV4(),
        Name: "Test Restaurant",
        CartItems: []models.CartItem{
            {
                Id:       uuid.NewV4(),
                Name:     "Burger",
                Price:    499,
                ImageURL: "burger.jpg",
                Weight:   250,
                Amount:   2,
            },
        },
    }
    testCartJSON, _ := json.Marshal(testCart)
    
    testOrder := models.Order{
        ID:            testOrderID,
        UserID:        testUserID.String(),
        Status:        "processing",
        Address:       testAddressID,
        OrderProducts: testCart,
        ApartmentOrOffice: "42",
        Intercom:      "1234",
        Entrance:      "1",
        Floor:         "4",
        CourierComment: "Call before arrival",
        LeaveAtDoor:   false,
        FinalPrice:    999.99,
        CreatedAt:     testTime,
    }
    
    columns := []string{
        "id", "user_id", "status", "address_id", "order_products",
        "apartment_or_office", "intercom", "entrance", "floor", 
        "courier_comment", "leave_at_door", "final_price", "created_at",
    }
    
    tests := []struct {
        name           string
        userID        uuid.UUID
        count         int
        offset        int
        repoMocker    func(*pgxpoolmock.MockPgxPool)
        expectedResult []models.Order
        expectError   bool
    }{
        {
            name:    "Success - single order",
            userID:  testUserID,
            count:   10,
            offset:  0,
            repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
                rows := pgxpoolmock.NewRows(columns).
                    AddRow(
                        testOrder.ID,
                        testOrder.UserID,
                        testOrder.Status,
                        testOrder.Address,
                        string(testCartJSON),
                        testOrder.ApartmentOrOffice,
                        testOrder.Intercom,
                        testOrder.Entrance,
                        testOrder.Floor,
                        testOrder.CourierComment,
                        testOrder.LeaveAtDoor,
                        testOrder.FinalPrice,
                        testTime,
                    ).ToPgxRows()
                
                mockPool.EXPECT().
                    Query(gomock.Any(), getAllOrders, testUserID, 10, 0).
                    Return(rows, nil)
            },
            expectedResult: []models.Order{testOrder},
            expectError:   false,
        },
        {
            name:    "Error - database query fails",
            userID:  testUserID,
            count:   10,
            offset:  0,
            repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
                mockPool.EXPECT().
                    Query(gomock.Any(), getAllOrders, testUserID, 10, 0).
                    Return(nil, fmt.Errorf("database error"))
            },
            expectedResult: nil,
            expectError:   true,
        },
        {
            name:    "Error - invalid JSON in order_products",
            userID:  testUserID,
            count:   10,
            offset:  0,
            repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
                rows := pgxpoolmock.NewRows(columns).
                    AddRow(
                        testOrder.ID,
                        testOrder.UserID,
                        testOrder.Status,
                        testOrder.Address,
                        "invalid json",
                        testOrder.ApartmentOrOffice,
                        testOrder.Intercom,
                        testOrder.Entrance,
                        testOrder.Floor,
                        testOrder.CourierComment,
                        testOrder.LeaveAtDoor,
                        testOrder.FinalPrice,
                        testTime,
                    ).ToPgxRows()
                
                mockPool.EXPECT().
                    Query(gomock.Any(), getAllOrders, testUserID, 10, 0).
                    Return(rows, nil)
            },
            expectedResult: nil,
            expectError:   true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()
            
            mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
            tt.repoMocker(mockPool)
            
            repo := &RestaurantRepository{db: mockPool}
            
            orders, err := repo.GetOrders(context.Background(), tt.userID, tt.count, tt.offset)
            
            if tt.expectError {
                assert.Error(t, err)
                if tt.expectedResult == nil {
                    assert.Nil(t, orders)
                }
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expectedResult, orders)
                
                // Проверка sanitize
                if len(orders) > 0 {
                    assert.NotEqual(t, "", orders[0].Status)
                    assert.NotEqual(t, "", orders[0].Address)
                    if len(orders[0].OrderProducts.CartItems) > 0 {
                        assert.NotEqual(t, "", orders[0].OrderProducts.CartItems[0].Name)
                    }
                }
            }
        })
    }
}

func TestGetOrderById(t *testing.T) {
    testUserID := uuid.NewV4()
    testOrderID := uuid.NewV4()
    testAddressID := "test_address_123"
    testTime := time.Now().UTC()
    
    // Тестовые данные для корзины
    testCart := models.Cart{
        Id:   uuid.NewV4(),
        Name: "Test Restaurant",
        CartItems: []models.CartItem{
            {
                Id:       uuid.NewV4(),
                Name:     "Burger",
                Price:    499,
                ImageURL: "burger.jpg",
                Weight:   250,
                Amount:   2,
            },
        },
    }
    testCartJSON, _ := json.Marshal(testCart)
    
    testOrder := models.Order{
        ID:            testOrderID,
        UserID:        testUserID.String(),
        Status:        "processing",
        Address:       testAddressID,
        OrderProducts: testCart,
        ApartmentOrOffice: "42",
        Intercom:      "1234",
        Entrance:      "1",
        Floor:         "4",
        CourierComment: "Call before arrival",
        LeaveAtDoor:   false,
        FinalPrice:    999.99,
        CreatedAt:     testTime,
    }
    
    columns := []string{
        "id", "user_id", "status", "address_id", "order_products",
        "apartment_or_office", "intercom", "entrance", "floor", 
        "courier_comment", "leave_at_door", "final_price", "created_at",
    }

    tests := []struct {
        name           string
        orderID       uuid.UUID
        userID        uuid.UUID
        repoMocker    func(*pgxpoolmock.MockPgxPool)
        expectedResult models.Order
        expectError   bool
        errorMessage  string
    }{
        {
            name:     "Success",
            orderID:  testOrderID,
            userID:   testUserID,
            repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
                row := pgxpoolmock.NewRows(columns).
                    AddRow(
                        testOrder.ID,
                        testOrder.UserID,
                        testOrder.Status,
                        testOrder.Address,
                        string(testCartJSON),
                        testOrder.ApartmentOrOffice,
                        testOrder.Intercom,
                        testOrder.Entrance,
                        testOrder.Floor,
                        testOrder.CourierComment,
                        testOrder.LeaveAtDoor,
                        testOrder.FinalPrice,
                        testTime,
                    ).ToPgxRows()
                row.Next()
                
                mockPool.EXPECT().
                    QueryRow(gomock.Any(), getOrderById, testOrderID, testUserID).
                    Return(row)
            },
            expectedResult: testOrder,
            expectError:   false,
        },
        {
            name:     "Error - invalid JSON",
            orderID:  testOrderID,
            userID:   testUserID,
            repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
                row := pgxpoolmock.NewRows(columns).
                    AddRow(
                        testOrder.ID,
                        testOrder.UserID,
                        testOrder.Status,
                        testOrder.Address,
                        "invalid json",
                        testOrder.ApartmentOrOffice,
                        testOrder.Intercom,
                        testOrder.Entrance,
                        testOrder.Floor,
                        testOrder.CourierComment,
                        testOrder.LeaveAtDoor,
                        testOrder.FinalPrice,
                        testTime,
                    ).ToPgxRows()
                row.Next()
                
                mockPool.EXPECT().
                    QueryRow(gomock.Any(), getOrderById, testOrderID, testUserID).
                    Return(row)
            },
            expectedResult: models.Order{},
            expectError:   true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()
            
            mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
            tt.repoMocker(mockPool)
            
            repo := &RestaurantRepository{db: mockPool}
            
            order, err := repo.GetOrderById(context.Background(), tt.orderID, tt.userID)
            
            if tt.expectError {
                assert.Error(t, err)
                if tt.errorMessage != "" {
                    assert.Contains(t, err.Error(), tt.errorMessage)
                }
                assert.Equal(t, models.Order{}, order)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expectedResult, order)
                
                // Проверка sanitize
                assert.NotEqual(t, "", order.Status)
                assert.NotEqual(t, "", order.Address)
                if len(order.OrderProducts.CartItems) > 0 {
                    assert.NotEqual(t, "", order.OrderProducts.CartItems[0].Name)
                }
            }
        })
    }
}