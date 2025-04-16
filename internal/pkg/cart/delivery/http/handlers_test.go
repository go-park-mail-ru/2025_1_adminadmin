package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
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
				r.Header.Set("X-CSRF-Token", "csrf2")
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
			expectedBody:   `{"id":"`,
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

func TestCartHandler_UpdateQuantityInCart(t *testing.T) {
	secret := "test-secret"
	login := "testuser"
	userId := uuid.NewV4()
	validToken := utils.GenerateJWTForTest(t, login, secret, userId)

	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		body           string
		productID      string
		mockSetup      func(mockUC *mocks.MockCartUsecase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Invalid JWT",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "invalid"})
			},
			productID:      uuid.NewV4().String(),
			body:           `{}`,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "некорректный JWT-токен",
		},
		{
			name: "Invalid CSRF Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf1"})
				r.Header.Set("X-CSRF-Token", "csrf2")
			},
			productID:      uuid.NewV4().String(),
			body:           `{}`,
			mockSetup:      func(mockUC *mocks.MockCartUsecase) {},
			expectedStatus: http.StatusForbidden,
			expectedBody:   "некорректный CSRF-токен",
		},
		{
			name: "Invalid Body",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			productID:      uuid.NewV4().String(),
			body:           `not-a-json`,
			mockSetup:      func(mockUC *mocks.MockCartUsecase) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Некорректный формат данных",
		},
		{
			name: "Update Error",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			productID: uuid.NewV4().String(),
			body: `{
				"restaurant_id": "` + uuid.NewV4().String() + `",
				"quantity": 2
			}`,
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().UpdateItemQuantity(gomock.Any(), login, gomock.Any(), gomock.Any(), int64(2)).
					Return(nil)
				mockUC.EXPECT().GetCart(gomock.Any(), login).Return(models.Cart{CartItems: []models.CartItem{}}, nil)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Не удалось обновить количество товара в корзине",
		},
		{
			name: "Success",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			productID: uuid.NewV4().String(),
			body: `{
				"restaurant_id": "` + uuid.NewV4().String() + `",
				"quantity": 3
			}`,
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().UpdateItemQuantity(gomock.Any(), login, gomock.Any(), gomock.Any(), int64(3)).
					Return(nil)
				mockUC.EXPECT().GetCart(gomock.Any(), login).Return(models.Cart{CartItems: []models.CartItem{}}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"id":"`,
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

			req := httptest.NewRequest("PUT", "/api/cart/"+tt.productID, strings.NewReader(tt.body))
			req = mux.SetURLVars(req, map[string]string{"productID": tt.productID})
			tt.cookieSetup(req)
			w := httptest.NewRecorder()

			handler := CartHandler{
				cartUsecase: mockUsecase,
				secret:      secret,
			}
			handler.UpdateQuantityInCart(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestCartHandler_ClearCart(t *testing.T) {
	secret := "test-secret"
	login := "testuser"
	userID := uuid.NewV4()
	validToken := utils.GenerateJWTForTest(t, login, secret, userID)

	tests := []struct {
		name           string
		setupRequest   func(r *http.Request)
		mockSetup      func(mockUC *mocks.MockCartUsecase)
		expectedStatus int
	}{
		{
			name:           "No JWT cookie",
			setupRequest:   func(r *http.Request) {},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Invalid CSRF",
			setupRequest: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf1"})
				r.Header.Set("X-CSRF-Token", "csrf2")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "ClearCart error",
			setupRequest: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().ClearCart(gomock.Any(), login).Return(errors.New("fail"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Success",
			setupRequest: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: validToken})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "csrf-token"})
				r.Header.Set("X-CSRF-Token", "csrf-token")
			},
			mockSetup: func(mockUC *mocks.MockCartUsecase) {
				mockUC.EXPECT().ClearCart(gomock.Any(), login).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUC := mocks.NewMockCartUsecase(ctrl)
			if tt.mockSetup != nil {
				tt.mockSetup(mockUC)
			}

			req := httptest.NewRequest(http.MethodDelete, "/api/cart/clear", nil)
			tt.setupRequest(req)
			w := httptest.NewRecorder()

			handler := CartHandler{
				cartUsecase: mockUC,
				secret:      secret,
			}

			handler.ClearCart(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
