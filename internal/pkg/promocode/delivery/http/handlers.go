package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/promocode"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/mailru/easyjson"
	"github.com/satori/uuid"
)

type PromocodeHandler struct {
	uc     interfaces.PromocodeUsecase
	secret string
}

func NewPromocodeHandler(uc interfaces.PromocodeUsecase) *PromocodeHandler {
	return &PromocodeHandler{uc: uc, secret: os.Getenv("JWT_SECRET")}
}

func (h *PromocodeHandler) GetPromocodes(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Ошибка авторизации", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка авторизации", http.StatusBadRequest)
		return
	}
	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "Ошибка авторизации", http.StatusForbidden)
		return
	}

	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	userIdStr, ok := jwtUtils.GetIdFromJWT(JWTStr, claims, h.secret)
	if !ok || userIdStr == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: id отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Ошибка авторизации", http.StatusUnauthorized)
		return
	}
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusInternalServerError)
		utils.SendError(w, "Ошибка авторизации", http.StatusInternalServerError)
		return
	}

	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 15
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	promocodes, err := h.uc.GetPromocodes(r.Context(), userId, count, offset)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(promocodes)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось получить промокоды", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *PromocodeHandler) CheckPromocode(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Ошибка авторизации", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка авторизации", http.StatusBadRequest)
		return
	}
	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		utils.SendError(w, "Ошибка авторизации", http.StatusForbidden)
		return
	}

	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	userIdStr, ok := jwtUtils.GetIdFromJWT(JWTStr, claims, h.secret)
	if !ok || userIdStr == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: id отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Ошибка авторизации", http.StatusUnauthorized)
		return
	}
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusInternalServerError)
		utils.SendError(w, "Ошибка авторизации", http.StatusInternalServerError)
		return
	}

	req := models.Promocode{}
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка чтения тела запроса: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	discount, err := h.uc.CheckPromocode(r.Context(), userId, req.Promocode)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	resp := models.PromocodeResp{Discount: discount}
	data, err := json.Marshal(resp)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось применить промокод", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}
