package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/golang/mock/gomock"
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

