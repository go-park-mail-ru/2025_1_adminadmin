package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/utils/send_error"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct {
	uc auth.AuthUsecase
}

func CreateAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignIn(r.Context(), req)

	if err != nil {
		if errors.Is(err, auth.ErrInvalidLogin) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrUserNotFound) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrInvalidCredentials) {
			utils.SendError(w, err.Error(), http.StatusUnauthorized)
		}
		if errors.Is(err, auth.ErrGeneratingToken) {
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", csrfToken)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignUp(r.Context(), req)

	if err != nil {
		if errors.Is(err, auth.ErrInvalidLogin) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrInvalidPassword) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrInvalidName) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrInvalidPhone) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrCreatingUser) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, auth.ErrGeneratingToken) {
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", csrfToken)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Check(w http.ResponseWriter, r *http.Request) {
	cookieCSRF, err := r.Cookie("CSRF-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headerCSRF := r.Header.Get("X-CSRF-Token")
	if cookieCSRF.Value == "" || headerCSRF == "" || cookieCSRF.Value != headerCSRF {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cookieJWT, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	JWTStr := cookieJWT.Value

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(JWTStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	login, ok := claims["login"].(string)
	if !ok || login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := h.uc.Check(r.Context(), login)

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil || cookie.Value == "" {
		utils.SendError(w, "Пользователь уже разлогинен", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}
