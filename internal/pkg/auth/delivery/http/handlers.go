package http

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
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

// SignIn godoc
// @Summary Авторизация пользователя
// @Description Вход пользователя по логину и паролю
// @Tags auth
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body models.SignInReq true "Данные для входа"
// @Success 200 {object} models.User "Успешный ответ с данными пользователя"
// @Failure 400 {object} utils.ErrorResponse "Ошибка парсинга или неправильные данные"
// @Failure 500 {object} utils.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/signin [post]
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
		Value:    user.Token2,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", user.Token2)
	w.Header().Set("Content-Type", "application/json")

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
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
