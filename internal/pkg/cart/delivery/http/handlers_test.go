package http

import (
	"errors"
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

func TestGetCartData(t *testing.T) {
    secret := "test-secret"
    login := "testuser"
    userId := uuid.NewV4()

    tests := []struct {
        name           string
        cookieSetup    func(r *http.Request)
        mockUsecase    func(mockUsecase *mocks.MockCartUsecase)
        expectedStatus int
        expectedErrStr string
        expectSuccess  bool
        expectedCart   models.Cart
    }{
        {
            name: "Usecase returns error",
            cookieSetup: func(r *http.Request) {
                token := utils.GenerateJWTForTest(t, login, secret, userId)
                r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: token})
            },
            mockUsecase: func(mockUsecase *mocks.MockCartUsecase) {
                mockUsecase.EXPECT().GetCart(gomock.Any(), login).
                    Return(models.Cart{}, errors.New("db error"), false)
            },
            expectedStatus: http.StatusInternalServerError,
            expectedErrStr: "ошибка получения корзины",
        },
        {
            name: "Successful cart fetch",
            cookieSetup: func(r *http.Request) {
                token := utils.GenerateJWTForTest(t, login, secret, userId)
                r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: token})
            },
            mockUsecase: func(mockUsecase *mocks.MockCartUsecase) {
                mockUsecase.EXPECT().GetCart(gomock.Any(), login).
                    Return(models.Cart{
                        Id:   uuid.NewV4(),
                        Name: "<Test & Restaurant>",
                        CartItems: []models.CartItem{
                            {
                                Id:       uuid.NewV4(),
                                Name:     "<Sushi & Roll>",
                                Price:    499.99,
                                ImageURL: "http://example.com/image.jpg?test=<bad>",
                                Weight:   300,
                                Amount:   2,
                            },
                        },
                    }, nil, true)
            },
            expectSuccess: true,
            expectedCart: models.Cart{
                Name: "Test &amp; Restaurant", // Экранированная строка
                CartItems: []models.CartItem{
                    {
                        Name:     "Sushi &amp; Roll",
                        ImageURL: "http://example.com/image.jpg?test=&lt;bad&gt;",
                    },
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            mockUsecase := mocks.NewMockCartUsecase(ctrl)
            defer ctrl.Finish()

            if tt.mockUsecase != nil {
                tt.mockUsecase(mockUsecase)
            }

            req := httptest.NewRequest("GET", "/api/cart", nil)
            tt.cookieSetup(req)
            h := CartHandler{cartUsecase: mockUsecase, secret: secret}

            cart, loginOut, err, full := h.getCartData(req)

            if tt.expectSuccess {
                assert.NoError(t, err)
                assert.Equal(t, login, loginOut)
                assert.True(t, full)

                // Проверим, что строки экранируются через Sanitize()
                cart.Sanitize()
                assert.Equal(t, tt.expectedCart.Name, cart.Name)
                assert.Equal(t, tt.expectedCart.CartItems[0].Name, cart.CartItems[0].Name)
                assert.Equal(t, tt.expectedCart.CartItems[0].ImageURL, cart.CartItems[0].ImageURL)
            } else {
                assert.Error(t, err)
                if tt.expectedErrStr != "" {
                    assert.Contains(t, err.Error(), tt.expectedErrStr)
                }
            }
        })
    }
}
