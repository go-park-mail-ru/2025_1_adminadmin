package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type CartHandler struct {
	cartUsecase *usecase.CartUsecase
	secret      string
}

func NewCartHandler(cartUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase, secret: os.Getenv("JWT_SECRET")}
}

func (h *CartHandler) getCartData(r *http.Request) (models.Cart, string, error) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			return models.Cart{}, "", fmt.Errorf("токен отсутствует")
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		return models.Cart{}, "", fmt.Errorf("ошибка при чтении куки")
	}

	JWTStr := cookie.Value
	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, fmt.Errorf("невалидный токен"), http.StatusUnauthorized)
		return models.Cart{}, "", fmt.Errorf("невалидный токен")
	}

	ctx := context.Background()
	cart, err := h.cartUsecase.GetCart(ctx, login)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка получения корзины: %w", err), http.StatusInternalServerError)
		return models.Cart{}, "", fmt.Errorf("ошибка получения корзины")
	}

	return cart, login, nil
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	cart, _, err := h.getCartData(r)
	if err != nil {
		utils.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(cart)
	if err != nil {
		log.LogHandlerError(log.GetLoggerFromContext(r.Context()), fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать корзину", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *CartHandler) UpdateQuantityInCart(w http.ResponseWriter, r *http.Request) {
	_, login, err := h.getCartData(r)
	if err != nil {
		utils.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	productID := vars["productID"]

	var requestBody models.CartInReq
	if err := easyjson.UnmarshalFromReader(r.Body, &requestBody); err != nil {
		utils.SendError(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.cartUsecase.UpdateItemQuantity(ctx, login, productID, requestBody.RestaurantId, requestBody.Quantity)
	if err != nil {
		utils.SendError(w, "Не удалось обновить количество товара в корзине", http.StatusInternalServerError)
		return
	}

	cart, _, err := h.getCartData(r)
	if err != nil {
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(cart)
	if err != nil {
		utils.SendError(w, "Ошибка сериализации корзины", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))
	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		return
	}

	JWTStr := cookie.Value
	claims := jwt.MapClaims{}

	login, _ := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	err = h.cartUsecase.ClearCart(r.Context(), login)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при очистке корзины: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
