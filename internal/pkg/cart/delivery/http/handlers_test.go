package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	productId := uuid.NewV4().String()
	restaurantId := uuid.NewV4().String()
	tests := []struct {
		name           string
		login          string
		grpcResponse   *gen.CartResponse
		grpcErr        error
		setupRequest   func() *http.Request
		expectedStatus int
	}{
		{
			name:  "GetCart_Success",
			login: "testuser",
			grpcResponse: &gen.CartResponse{
				RestaurantId:   restaurantId,
				RestaurantName: "Test Restaurant",
				Products: []*gen.CartItem{
					{
						Id:       productId,
						Name:     "Product 1",
						Price:    10.5,
						ImageUrl: "image1.jpg",
						Weight:   100,
						Amount:   2,
					},
				},
				FullCart: true,
			},
			grpcErr: nil,
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("GET", "/cart", nil)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
				return r
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:  "GetCart_EmptyCart",
			login: "testuser",
			grpcResponse: &gen.CartResponse{
				RestaurantId:   "",
				RestaurantName: "",
				Products:       []*gen.CartItem{},
				FullCart:       false,
			},
			grpcErr: nil,
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("GET", "/cart", nil)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
				return r
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockCartServiceClient(ctrl)

			if tt.login != "" && tt.grpcResponse != nil {
				mockClient.EXPECT().GetCart(
					gomock.Any(),
					&gen.GetCartRequest{Login: tt.login},
				).Return(tt.grpcResponse, tt.grpcErr)
			}

			handler := CartHandler{
				client: mockClient,
				secret: secret,
			}

			req := tt.setupRequest()
			w := httptest.NewRecorder()

			handler.GetCart(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.name == "GetCart_Success" {
				var response models.Cart
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Len(t, response.CartItems, 1)
			}
		})
	}
}

func TestUpdateQuantityInCart(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrfToken := "test-csrf"
	userID := uuid.NewV4()
	productID := uuid.NewV4().String()
	restaurantID := uuid.NewV4().String()

	validCart := &gen.CartResponse{
		RestaurantId:   restaurantID,
		RestaurantName: "Test Restaurant",
		Products: []*gen.CartItem{
			{
				Id:       productID,
				Name:     "Product 1",
				Price:    10.5,
				ImageUrl: "image1.jpg",
				Weight:   100,
				Amount:   2,
			},
		},
		FullCart: true,
	}

	tests := []struct {
		name             string
		requestBody      string
		setupRequest     func() *http.Request
		expectStatus     int
		mockGrpcBehavior func(mockClient *mocks.MockCartServiceClient)
	}{
		{
			name:         "UpdateQuantity_Success",
			requestBody:  fmt.Sprintf(`{"quantity": 3, "restaurant_id": "%s"}`, restaurantID),
			expectStatus: http.StatusOK,
			setupRequest: func() *http.Request {
				body := strings.NewReader(fmt.Sprintf(`{"quantity": 3, "restaurant_id": "%s"}`, restaurantID))
				r := httptest.NewRequest("PUT", fmt.Sprintf("/cart/%s", productID), body)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})
				r.Header.Set("X-CSRF-Token", csrfToken)
				return r
			},
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {
				mockClient.EXPECT().GetCart(gomock.Any(), &gen.GetCartRequest{Login: login}).Return(validCart, nil).Times(2)
				mockClient.EXPECT().UpdateItemQuantity(gomock.Any(), &gen.UpdateQuantityRequest{
					Login:        login,
					ProductId:    productID,
					RestaurantId: restaurantID,
					Quantity:     3,
				}).Return(&empty.Empty{}, nil)
			},
		},
		{
			name:         "UpdateQuantity_InvalidCSRF",
			requestBody:  fmt.Sprintf(`{"quantity": 1, "restaurant_id": "%s"}`, restaurantID),
			expectStatus: http.StatusUnauthorized,
			setupRequest: func() *http.Request {
				body := strings.NewReader(fmt.Sprintf(`{"quantity": 1, "restaurant_id": "%s"}`, restaurantID))
				r := httptest.NewRequest("PUT", fmt.Sprintf("/cart/%s", productID), body)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				return r
			},
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {
				mockClient.EXPECT().GetCart(gomock.Any(), &gen.GetCartRequest{Login: login}).Return(validCart, nil).AnyTimes()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockCartServiceClient(ctrl)
			tt.mockGrpcBehavior(mockClient)

			handler := CartHandler{
				client: mockClient,
				secret: secret,
			}

			req := tt.setupRequest()
			req = mux.SetURLVars(req, map[string]string{"productID": productID})
			w := httptest.NewRecorder()

			handler.UpdateQuantityInCart(w, req)

			assert.Equal(t, tt.expectStatus, w.Code)
		})
	}
}

func TestClearCart(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrfToken := "test-csrf"
	userID := uuid.NewV4()

	tests := []struct {
		name             string
		setupRequest     func() *http.Request
		expectStatus     int
		mockGrpcBehavior func(mockClient *mocks.MockCartServiceClient)
	}{
		{
			name: "ClearCart_Success",
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("DELETE", "/cart", nil)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})
				r.Header.Set("X-CSRF-Token", csrfToken)
				return r
			},
			expectStatus: http.StatusOK,
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {
				mockClient.EXPECT().ClearCart(gomock.Any(), &gen.ClearCartRequest{
					Login: login,
				}).Return(&empty.Empty{}, nil)
			},
		},
		{
			name: "ClearCart_NoToken",
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("DELETE", "/cart", nil)
				// No AdminJWT cookie
				return r
			},
			expectStatus:     http.StatusUnauthorized,
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {},
		},
		{
			name: "ClearCart_InvalidLogin",
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("DELETE", "/cart", nil)
				// Генерация токена с некорректным логином
				tokenStr := utils.GenerateJWTForTest(t, "", secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})
				r.Header.Set("X-CSRF-Token", csrfToken)
				return r
			},
			expectStatus:     http.StatusForbidden,
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {},
		},
		{
			name: "ClearCart_InvalidCSRF",
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("DELETE", "/cart", nil)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				// No CSRF token
				return r
			},
			expectStatus: http.StatusUnauthorized,
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {
				// CSRF error happens before ClearCart
			},
		},
		{
			name: "ClearCart_GRPCError",
			setupRequest: func() *http.Request {
				r := httptest.NewRequest("DELETE", "/cart", nil)
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userID)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})
				r.Header.Set("X-CSRF-Token", csrfToken)
				return r
			},
			expectStatus: http.StatusInternalServerError,
			mockGrpcBehavior: func(mockClient *mocks.MockCartServiceClient) {
				mockClient.EXPECT().ClearCart(gomock.Any(), &gen.ClearCartRequest{
					Login: login,
				}).Return(nil, fmt.Errorf("gRPC error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockCartServiceClient(ctrl)
			tt.mockGrpcBehavior(mockClient)

			handler := CartHandler{
				client: mockClient,
				secret: secret,
			}

			req := tt.setupRequest()
			w := httptest.NewRecorder()

			handler.ClearCart(w, req)

			assert.Equal(t, tt.expectStatus, w.Code)
		})
	}
}
