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
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart"

	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	validation "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/validation"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type CartHandler struct {
	cartUsecase cart.CartUsecase
	secret      string
}

func NewCartHandler(cartUsecase cart.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase, secret: os.Getenv("JWT_SECRET")}
}

func (h *CartHandler) getCartData(r *http.Request) (models.Cart, string, error, bool) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			return models.Cart{}, "", fmt.Errorf("токен отсутствует"), false
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		return models.Cart{}, "", fmt.Errorf("ошибка при чтении куки"), false
	}

	JWTStr := cookie.Value
	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, fmt.Errorf("невалидный токен"), http.StatusUnauthorized)
		return models.Cart{}, "", fmt.Errorf("невалидный токен"), false
	}

	ctx := context.Background()
	cart, err, full_cart := h.cartUsecase.GetCart(ctx, login)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка получения корзины: %w", err), http.StatusInternalServerError)
		return models.Cart{}, "", fmt.Errorf("ошибка получения корзины"), false
	}

	return cart, login, nil, full_cart
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context())

	cart, _, err, full_cart := h.getCartData(r)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusUnauthorized)
		utils.SendError(w, "некорректный JWT-токен", http.StatusUnauthorized)
		return
	}
	cart.Sanitize()

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		return
	}

	if !full_cart {
		log.LogHandlerError(logger, fmt.Errorf("корзина пуста"), http.StatusNotFound)
		utils.SendError(w, "корзина пуста", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(cart)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать корзину", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *CartHandler) UpdateQuantityInCart(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context())

	_, login, err, _ := h.getCartData(r)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusUnauthorized)
		utils.SendError(w, "некорректный JWT-токен", http.StatusUnauthorized)
		return
	}

	if login == "" {
		log.LogHandlerError(logger, errors.New("невалидный токен"), http.StatusUnauthorized)
		utils.SendError(w, "невалидный токен", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	productID := vars["productID"]

	var requestBody models.CartInReq
	if err := easyjson.UnmarshalFromReader(r.Body, &requestBody); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка чтения тела запроса: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}

	requestBody.Sanitize()

	ctx := context.Background()
	err = h.cartUsecase.UpdateItemQuantity(ctx, login, productID, requestBody.RestaurantId, requestBody.Quantity)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("не удалось обновить количество: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось обновить количество товара в корзине", http.StatusInternalServerError)
		return
	}

	cart, _, err, full_cart := h.getCartData(r)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusUnauthorized)
		utils.SendError(w, "некорректный JWT-токен", http.StatusUnauthorized)
		return
	}

	if !full_cart {
		log.LogHandlerError(logger, fmt.Errorf("корзина пуста"), http.StatusOK)
		utils.SendError(w, "корзина пуста", http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(cart)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка сериализации корзины: %w", err), http.StatusInternalServerError)
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
			utils.SendError(w, "JWT cookie not found", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Bad request", http.StatusBadRequest) 
		return
	}


	JWTStr := cookie.Value
	claims := jwt.MapClaims{}

	login, _ := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)

	if login == "" {
		log.LogHandlerError(logger, errors.New("пустой login из токена"), http.StatusUnauthorized)
		utils.SendError(w, "некорректный JWT токен", http.StatusForbidden)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
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

func (h *CartHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context())

	cart, userID, err, full_cart := h.getCartData(r)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusUnauthorized)
		utils.SendError(w, "некорректный JWT-токен", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		return
	}

	if !full_cart {
		log.LogHandlerError(logger, fmt.Errorf("корзина пуста"), http.StatusNotFound)
		utils.SendError(w, "корзина пуста", http.StatusNotFound)
		return
	}

	var req models.OrderInReq
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка чтения тела запроса: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}
	cart.Sanitize()

	if err := validation.ValidateOrderInput(&req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("валидация заказа: %w", err), http.StatusBadRequest)
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	order, err := h.cartUsecase.CreateOrder(ctx, userID, req, cart)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("не удалось создать заказ: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка при создании заказа", http.StatusInternalServerError)
		return
	}

	err = h.cartUsecase.ClearCart(r.Context(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при очистке корзины: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(order)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать корзину", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
