package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	restaurantID := uuid.NewV4()
	productID := uuid.NewV4()
	login := "testuser"

	tests := []struct {
		name           string
		login          string
		mockSetup      func()
		expected       *gen.CartResponse
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name:  "Success",
			login: login,
			mockSetup: func() {
				mockUsecase.EXPECT().GetCart(gomock.Any(), login).
					Return(models.Cart{
						Id:   restaurantID,
						Name: "Test Restaurant",
						CartItems: []models.CartItem{
							{
								Id:       productID,
								Name:     "Product 1",
								Price:    10.5,
								ImageURL: "image1.jpg",
								Weight:   100,
								Amount:   2,
							},
						},
					}, nil, true)
			},
			expected: &gen.CartResponse{
				RestaurantId:   restaurantID.String(),
				RestaurantName: "Test Restaurant",
				Products: []*gen.CartItem{
					{
						Id:       productID.String(),
						Name:     "Product 1",
						Price:    10.5,
						ImageUrl: "image1.jpg",
						Weight:   100,
						Amount:   2,
					},
				},
				FullCart: true,
			},
			expectedErr: nil,
		},
		{
			name:  "Error",
			login: login,
			mockSetup: func() {
				mockUsecase.EXPECT().GetCart(gomock.Any(), login).
					Return(models.Cart{}, errors.New("some error"), false)
			},
			expected:       &gen.CartResponse{},
			expectedErr:    status.Errorf(codes.Internal, "ошибка получения корзины"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := &gen.GetCartRequest{Login: tt.login}
			resp, err := h.GetCart(context.Background(), req)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			}
		})
	}
}

func TestUpdateItemQuantity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	tests := []struct {
		name           string
		input          *gen.UpdateQuantityRequest
		mockSetup      func()
		expected       *emptypb.Empty
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.UpdateQuantityRequest{
				Login:        "testuser",
				ProductId:    "product123",
				RestaurantId: "restaurant456",
				Quantity:     2,
			},
			mockSetup: func() {
				mockUsecase.EXPECT().UpdateItemQuantity(
					gomock.Any(),
					"testuser",
					"product123",
					"restaurant456",
					2,
				).Return(nil)
			},
			expected:    &emptypb.Empty{},
			expectedErr: nil,
		},
		{
			name: "ErrorFromUsecase",
			input: &gen.UpdateQuantityRequest{
				Login:        "testuser",
				ProductId:    "product123",
				RestaurantId: "restaurant456",
				Quantity:     2,
			},
			mockSetup: func() {
				mockUsecase.EXPECT().UpdateItemQuantity(
					gomock.Any(),
					"testuser",
					"product123",
					"restaurant456",
					2,
				).Return(errors.New("some error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "some error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.UpdateItemQuantity(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			}
		})
	}
}

func TestClearCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	tests := []struct {
		name           string
		input          *gen.ClearCartRequest
		mockSetup      func()
		expected       *emptypb.Empty
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.ClearCartRequest{
				Login: "testuser",
			},
			mockSetup: func() {
				mockUsecase.EXPECT().ClearCart(
					gomock.Any(),
					"testuser",
				).Return(nil)
			},
			expected:    &emptypb.Empty{},
			expectedErr: nil,
		},
		{
			name: "ErrorFromUsecase",
			input: &gen.ClearCartRequest{
				Login: "testuser",
			},
			mockSetup: func() {
				mockUsecase.EXPECT().ClearCart(
					gomock.Any(),
					"testuser",
				).Return(errors.New("database error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "database error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.ClearCart(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			}
		})
	}
}

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	restaurantID := uuid.NewV4()
	productID := uuid.NewV4()
	login := "testuser"

	tests := []struct {
		name           string
		input          *gen.CreateOrderRequest
		mockSetup      func()
		expected       *gen.OrderResponse
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.CreateOrderRequest{
				Status:            "pending",
				Address:           "Test Address",
				ApartmentOrOffice: "42",
				Intercom:          "1234",
				Entrance:          "1",
				Floor:             "4",
				CourierComment:    "Call me",
				LeaveAtDoor:       true,
				FinalPrice:        100.50,
				Login:             login,
				Cart: &gen.CartResponse{
					RestaurantId:   restaurantID.String(),
					RestaurantName: "Test Restaurant",
					Products: []*gen.CartItem{
						{
							Id:       productID.String(),
							Name:     "Product 1",
							Price:    10.5,
							ImageUrl: "image1.jpg",
							Weight:   100,
							Amount:   2,
						},
					},
				},
			},
			mockSetup: func() {
				mockUsecase.EXPECT().CreateOrder(
					gomock.Any(),
					login,
					models.OrderInReq{
						Status:            "pending",
						Address:           "Test Address",
						ApartmentOrOffice: "42",
						Intercom:          "1234",
						Entrance:          "1",
						Floor:             "4",
						CourierComment:    "Call me",
						LeaveAtDoor:       true,
						FinalPrice:        100.50,
					},
					models.Cart{
						Id:   restaurantID,
						Name: "Test Restaurant",
						CartItems: []models.CartItem{
							{
								Id:       productID,
								Name:     "Product 1",
								Price:    10.5,
								ImageURL: "image1.jpg",
								Weight:   100,
								Amount:   2,
							},
						},
					},
				).Return(models.Order{
					ID:                uuid.NewV4(),
					Status:            "pending",
					Address:           "Test Address",
					ApartmentOrOffice: "42",
					Intercom:          "1234",
					Entrance:          "1",
					Floor:             "4",
					CourierComment:    "Call me",
					LeaveAtDoor:       true,
					FinalPrice:        100.50,
				}, nil)
			},
			expected: &gen.OrderResponse{
				Status:            "pending",
				Address:           "Test Address",
				ApartmentOrOffice: "42",
				Intercom:          "1234",
				Entrance:          "1",
				Floor:             "4",
				CourierComment:    "Call me",
				LeaveAtDoor:       true,
				FinalPrice:        100.50,
			},
			expectedErr: nil,
		},
		{
			name: "InvalidRestaurantID",
			input: &gen.CreateOrderRequest{
				Cart: &gen.CartResponse{
					RestaurantId: "invalid-uuid",
				},
			},
			mockSetup:      func() {},
			expected:       nil,
			expectedErr:    status.Errorf(codes.InvalidArgument, "invalid restaurant ID"),
			expectedStatus: codes.InvalidArgument,
		},
		{
			name: "UsecaseError",
			input: &gen.CreateOrderRequest{
				Cart: &gen.CartResponse{
					RestaurantId: restaurantID.String(),
					Products: []*gen.CartItem{
						{
							Id: productID.String(),
						},
					},
				},
			},
			mockSetup: func() {
				mockUsecase.EXPECT().CreateOrder(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Return(models.Order{}, errors.New("database error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "database error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.CreateOrder(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				// Проверяем основные поля, так как некоторые (ID, CreatedAt) могут генерироваться автоматически
				assert.Equal(t, tt.expected.Status, resp.Status)
				assert.Equal(t, tt.expected.Address, resp.Address)
				assert.Equal(t, tt.expected.FinalPrice, resp.FinalPrice)
			}
		})
	}
}

func TestGetOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	userID := uuid.NewV4()
	orderID := uuid.NewV4()

	tests := []struct {
		name           string
		input          *gen.GetOrdersRequest
		mockSetup      func()
		expected       *gen.OrderListResponse
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.GetOrdersRequest{
				UserId: userID.String(),
				Count:  10,
				Offset: 0,
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrders(
					gomock.Any(),
					userID,
					10,
					0,
				).Return([]models.Order{
					{
						ID:                orderID,
						Status:            "delivered",
						Address:           "Test Address",
						ApartmentOrOffice: "42",
						FinalPrice:        100.50,
						CreatedAt:         time.Now(),
					},
				}, nil)
			},
			expected: &gen.OrderListResponse{
				Orders: []*gen.OrderResponse{
					{
						Id:                orderID.String(),
						Status:            "delivered",
						Address:           "Test Address",
						ApartmentOrOffice: "42",
						FinalPrice:        100.50,
						UserId:            userID.String(),
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "InvalidUserID",
			input: &gen.GetOrdersRequest{
				UserId: "invalid-uuid",
			},
			mockSetup:      func() {},
			expected:       nil,
			expectedErr:    status.Errorf(codes.InvalidArgument, "invalid restaurant ID"),
			expectedStatus: codes.InvalidArgument,
		},
		{
			name: "EmptyResult",
			input: &gen.GetOrdersRequest{
				UserId: userID.String(),
				Count:  10,
				Offset: 0,
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrders(
					gomock.Any(),
					userID,
					10,
					0,
				).Return([]models.Order{}, nil)
			},
			expected: &gen.OrderListResponse{
				Orders: []*gen.OrderResponse{},
			},
			expectedErr: nil,
		},
		{
			name: "UsecaseError",
			input: &gen.GetOrdersRequest{
				UserId: userID.String(),
				Count:  10,
				Offset: 0,
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrders(
					gomock.Any(),
					userID,
					10,
					0,
				).Return(nil, errors.New("database error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "database error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.GetOrders(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				require.Equal(t, len(tt.expected.Orders), len(resp.Orders))

				if len(tt.expected.Orders) > 0 {
					expectedOrder := tt.expected.Orders[0]
					actualOrder := resp.Orders[0]

					assert.Equal(t, expectedOrder.Id, actualOrder.Id)
					assert.Equal(t, expectedOrder.Status, actualOrder.Status)
					assert.Equal(t, expectedOrder.Address, actualOrder.Address)
					assert.Equal(t, expectedOrder.FinalPrice, actualOrder.FinalPrice)
					assert.Equal(t, expectedOrder.Id, actualOrder.Id)
					assert.Equal(t, expectedOrder.UserId, actualOrder.UserId)
				}
			}
		})
	}
}

func TestGetOrderById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	userID := uuid.NewV4()
	orderID := uuid.NewV4()

	tests := []struct {
		name           string
		input          *gen.GetOrderByIdRequest
		mockSetup      func()
		expected       *gen.OrderResponse
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.GetOrderByIdRequest{
				UserId:  userID.String(),
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrderById(
					gomock.Any(),
					orderID,
					userID,
				).Return(models.Order{
					ID:                orderID,
					Status:            "delivered",
					Address:           "Test Address",
					ApartmentOrOffice: "42",
					Intercom:          "1234",
					Entrance:          "1",
					Floor:             "4",
					CourierComment:    "Call me",
					LeaveAtDoor:       true,
					FinalPrice:        100.50,
					CreatedAt:         time.Now(),
				}, nil)
			},
			expected: &gen.OrderResponse{
				Id:                orderID.String(),
				Status:            "delivered",
				Address:           "Test Address",
				ApartmentOrOffice: "42",
				Intercom:          "1234",
				Entrance:          "1",
				Floor:             "4",
				CourierComment:    "Call me",
				LeaveAtDoor:       true,
				FinalPrice:        100.50,
				UserId:           userID.String(),
			},
			expectedErr: nil,
		},
		{
			name: "InvalidUserID",
			input: &gen.GetOrderByIdRequest{
				UserId:  "invalid-uuid",
				OrderId: orderID.String(),
			},
			mockSetup:      func() {},
			expected:       nil,
			expectedErr:    status.Errorf(codes.InvalidArgument, "invalid user ID"),
			expectedStatus: codes.InvalidArgument,
		},
		{
			name: "InvalidOrderID",
			input: &gen.GetOrderByIdRequest{
				UserId:  userID.String(),
				OrderId: "invalid-uuid",
			},
			mockSetup:      func() {},
			expected:       nil,
			expectedErr:    status.Errorf(codes.InvalidArgument, "invalid order ID"),
			expectedStatus: codes.InvalidArgument,
		},
		{
			name: "OrderNotFound",
			input: &gen.GetOrderByIdRequest{
				UserId:  userID.String(),
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrderById(
					gomock.Any(),
					orderID,
					userID,
				).Return(models.Order{}, errors.New("order not found"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "failed to get order: order not found"),
			expectedStatus: codes.Internal,
		},
		{
			name: "UsecaseError",
			input: &gen.GetOrderByIdRequest{
				UserId:  userID.String(),
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().GetOrderById(
					gomock.Any(),
					orderID,
					userID,
				).Return(models.Order{}, errors.New("database error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "failed to get order: database error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.GetOrderById(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Id, resp.Id)
				assert.Equal(t, tt.expected.Status, resp.Status)
				assert.Equal(t, tt.expected.Address, resp.Address)
				assert.Equal(t, tt.expected.ApartmentOrOffice, resp.ApartmentOrOffice)
				assert.Equal(t, tt.expected.FinalPrice, resp.FinalPrice)
				assert.Equal(t, tt.expected.Id, resp.Id)
				assert.Equal(t, tt.expected.UserId, resp.UserId)
			}
		})
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockCartUsecase(ctrl)
	h := CreateCartHandler(mockUsecase)

	orderID := uuid.NewV4()

	tests := []struct {
		name           string
		input          *gen.UpdateOrderStatusRequest
		mockSetup      func()
		expected       *emptypb.Empty
		expectedErr    error
		expectedStatus codes.Code
	}{
		{
			name: "Success",
			input: &gen.UpdateOrderStatusRequest{
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().UpdateOrderStatus(
					gomock.Any(),
					orderID,
				).Return(nil)
			},
			expected:       &emptypb.Empty{},
			expectedErr:    nil,
			expectedStatus: codes.OK,
		},
		{
			name: "InvalidOrderID",
			input: &gen.UpdateOrderStatusRequest{
				OrderId: "invalid-uuid",
			},
			mockSetup:      func() {},
			expected:       nil,
			expectedErr:    status.Errorf(codes.InvalidArgument, "invalid order ID"),
			expectedStatus: codes.InvalidArgument,
		},
		{
			name: "OrderNotFound",
			input: &gen.UpdateOrderStatusRequest{
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().UpdateOrderStatus(
					gomock.Any(),
					orderID,
				).Return(errors.New("order not found"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "failed to update order status: order not found"),
			expectedStatus: codes.Internal,
		},
		{
			name: "UsecaseError",
			input: &gen.UpdateOrderStatusRequest{
				OrderId: orderID.String(),
			},
			mockSetup: func() {
				mockUsecase.EXPECT().UpdateOrderStatus(
					gomock.Any(),
					orderID,
				).Return(errors.New("database error"))
			},
			expected:       nil,
			expectedErr:    status.Errorf(codes.Internal, "failed to update order status: database error"),
			expectedStatus: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := h.UpdateOrderStatus(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedStatus, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			}
		})
	}
}