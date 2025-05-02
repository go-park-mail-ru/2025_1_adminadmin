package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/satori/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/mailru/easyjson"
)

const maxRequestBodySize = 10 << 20

var allowedMimeTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

type AuthHandler struct {
	client gen.AuthServiceClient
	secret string
}

func CreateAuthHandler(client gen.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: client, secret: os.Getenv("JWT_SECRET")}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignInReq
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	req.Sanitize()

	user, err := h.client.SignIn(r.Context(), &gen.SignInRequest{
		Login:    req.Login,
		Password: req.Password})

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrUserNotFound:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidCredentials:
			log.LogHandlerError(logger, err, http.StatusUnauthorized)
			utils.SendError(w, err.Error(), http.StatusUnauthorized)
		default:
			log.LogHandlerError(logger, fmt.Errorf("Неизвестная ошибка: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "Неизвестная ошибка", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    user.Token,
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    user.CsrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", user.CsrfToken)
	w.Header().Set("Content-Type", "application/json")

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
    	log.LogHandlerError(logger, fmt.Errorf("Некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "Некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignUpReq
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	req.Sanitize()

	user, err := h.client.SignUp(r.Context(), &gen.SignUpRequest{
		Login:       req.Login,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrInvalidPassword:
			log.LogHandlerError(logger, fmt.Errorf("неправильный логин или пароль: %w", err), http.StatusBadRequest)
			utils.SendError(w, "Неправильный логин или пароль", http.StatusBadRequest)
		case auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrCreatingUser:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		default:
			log.LogHandlerError(logger, fmt.Errorf("неизвестная ошибка: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "Неизвестная ошибка", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    user.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    user.CsrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", user.CsrfToken)
	w.Header().Set("Content-Type", "application/json")

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
    	log.LogHandlerError(logger, fmt.Errorf("Некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "Некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) Check(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

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

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := h.client.Check(r.Context(), &gen.CheckRequest{
		Login: login,
	})

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
    	log.LogHandlerError(logger, fmt.Errorf("Некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "Некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}
