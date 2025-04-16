package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCartHandler_GetCart(t *testing.T) {
	secret := "test-secret"
	login := "testuser"
	userId := uuid.NewV4()
	validToken := utils.GenerateJWTForTest(t, login, secret, userId)

	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		mockSetup      func(mockUC *mocks.MockCartUsecase)
		expectedStatus int
		expectedBody   string
		fullCart       bool
		cart           models.Cart
	}{
		{
			name: "Invalid JWT",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "invalid-token"})
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "некорректный JWT-токен",
		},
		{
			name: "Invalid CSRF Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf1"})
				r.Header.Set("X-CSRF-Token", "csrf2") // Несовпадение
			},
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().GetCart(gomock.Any(), login).
					Return(models.Cart{}, nil, true)
			},
			expectedStatus: http.StatusForbidden,
			expectedBody:   "некорректный CSRF-токен",
		},
		{
			name: "Empty Cart",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().GetCart(gomock.Any(), login).
					Return(models.Cart{}, nil, false)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "корзина пуста",
		},

		{
			name: "Success",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().GetCart(gomock.Any(), login).
					Return(models.Cart{
						Id:   uuid.NewV4(),
						Name: "<Cart>",
						CartItems: []models.CartItem{
							{
								Id:       uuid.NewV4(),
								Name:     "<Roll>",
								ImageURL: "http://test.com/image?x=<x>",
								Price:    100,
								Weight:   200,
								Amount:   1,
							},
						},
					}, nil, true)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":"`, // просто проверка что пришел JSON
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mocks.NewMockCartUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			req := httptest.NewRequest("GET", "/api/cart", nil)
			tt.cookieSetup(req)
			w := httptest.NewRecorder()

			handler := CartHandler{
				cartUsecase: mockUsecase,
				secret:      secret,
			}
			handler.GetCart(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}
