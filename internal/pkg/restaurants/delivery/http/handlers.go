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
	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/satori/uuid"
)

type RestaurantHandler struct {
	restaurantUsecase interfaces.RestaurantUsecase
}

func NewRestaurantHandler(ru interfaces.RestaurantUsecase) *RestaurantHandler {
	return &RestaurantHandler{restaurantUsecase: ru}
}

// GetProductsByRestaurant godoc
// @Summary Получить продукты ресторана
// @Description Получение списка продуктов ресторана с пагинацией
// @Tags restaurants
// @Param id path string true "ID ресторана"
// @Param count query int false "Количество элементов (по умолчанию 100)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Produce json
// @Success 200 {array} models.Product "Успешное получение продуктов ресторана"
// @Failure 400 {object} utils.ErrorResponse "Неверный формат ID ресторана"
// @Failure 500 {object} utils.ErrorResponse "Внутренняя ошибка сервера"
// @Router /restaurants/{id} [get]
func (h *RestaurantHandler) GetProductsByRestaurant(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
		log.LogHandlerError(logger, errors.New("неверный формат id ресторана"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	products, err := h.restaurantUsecase.GetProductsByRestaurant(r.Context(), restaurantID, count, offset)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

// RestaurantList godoc
// @Summary Список ресторанов
// @Description Получение списка ресторанов с пагинацией
// @Tags restaurants
// @Param count query int false "Количество элементов"
// @Param offset query int false "Смещение"
// @Produce json
// @Success 200 {array} models.Restaurant "Успешное получение списка ресторанов"
// @Failure 500 {object} utils.ErrorResponse "Внутренняя ошибка сервера"
// @Router /restaurants/list [get]
func (h *RestaurantHandler) RestaurantList(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	restaurants, err := h.restaurantUsecase.GetAll(r.Context(), count, offset)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if restaurants == nil {
		log.LogHandlerError(logger, fmt.Errorf("рестораны не найдены: %w", err), http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(restaurants)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *RestaurantHandler) ReviewsList(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 1
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
		log.LogHandlerError(logger, errors.New("неверный формат id ресторана"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reviews, err := h.restaurantUsecase.GetReviews(r.Context(), restaurantID, count, offset)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if reviews == nil {
		log.LogHandlerError(logger, fmt.Errorf("отзывы не найдены: %w", err), http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(reviews)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *RestaurantHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
		log.LogHandlerError(logger, errors.New("неверный формат id ресторана"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.ReviewInReq
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	req.Sanitize()
	if req.Rating < 1 || req.Rating > 5 {
		log.LogHandlerError(logger, fmt.Errorf("рейтинг должен быть от 1 до 5"), http.StatusBadRequest)
		utils.SendError(w, "рейтинг должен быть от 1 до 5", http.StatusBadRequest)
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

	idS, ok := jwtUtils.GetIdFromJWT(JWTStr, claims, os.Getenv("JWT_SECRET"))
	if !ok || idS == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := uuid.FromString(idS)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, os.Getenv("JWT_SECRET"))
	if !ok || idS == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	review, err := h.restaurantUsecase.CreateReview(r.Context(), req, id, restaurantID, login)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(review); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}
