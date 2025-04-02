package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	authHandler "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/http"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/usecase"
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
	login, ok := authHandler.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return "", nil
	}
	return login, nil
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	login, err := h.getLoginFromCookie(w, r)
	if err != nil || login == "" {
		return
	}

	ctx := context.Background()
	cart, err := h.cartUsecase.GetCart(ctx, login)
	if err != nil {
		utils.SendError(w, err.Error(), http.StatusNotFound)
		utils.SendError(w, "Не удалось получить корзину", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
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
