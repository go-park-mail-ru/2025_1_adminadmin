package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type CartHandler struct {
	cartUsecase *usecase.CartUsecase
	secret      string
}

func NewCartHandler(cartUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase, secret: os.Getenv("JWT_SECRET")}
}

func (h *CartHandler) getLoginFromCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookieJWT, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	JWTStr := cookieJWT.Value
	claims := jwt.MapClaims{}
	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return "", nil
	}
	return login, nil
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))
	login, err := h.getLoginFromCookie(w, r)
	if err != nil || login == "" {
		return
	}

	ctx := context.Background()
	items, err := h.cartUsecase.GetCart(ctx, login)
	if err != nil {
		utils.SendError(w, "Не удалось получить корзину", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(items)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
    login, err := h.getLoginFromCookie(w, r)
    if err != nil || login == "" {
        return
    }

    vars := mux.Vars(r)
    productID := vars["productID"]

    ctx := context.Background()
    err = h.cartUsecase.AddItem(ctx, login, productID)
    if err != nil {
        utils.SendError(w, "Не удалось добавить товар в корзину", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) UpdateQuantityInCart(w http.ResponseWriter, r *http.Request) {
    login, err := h.getLoginFromCookie(w, r)
    if err != nil || login == "" {
        return
    }

    vars := mux.Vars(r)
    productID := vars["productID"]

    var requestBody struct {
        Quantity int `json:"quantity"`
		RestaurantId string `json:"restaurantId"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        utils.SendError(w, "Некорректный формат данных", http.StatusBadRequest)
        return
    }

    ctx := context.Background()
    err = h.cartUsecase.UpdateItemQuantity(ctx, login, productID, requestBody.RestaurantId, requestBody.Quantity)
    if err != nil {
        utils.SendError(w, "Не удалось обновить количество товара в корзине", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	login, err := h.getLoginFromCookie(w, r)
	if err != nil || login == "" {
		return
	}

	vars := mux.Vars(r)
	productID := vars["productID"]

	ctx := context.Background()
	err = h.cartUsecase.RemoveItem(ctx, login, productID)
	if err != nil {
		utils.SendError(w, "Не удалось удалить товар из корзины", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
