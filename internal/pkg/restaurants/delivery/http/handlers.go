package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
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

func (h *RestaurantHandler) GetProductsByRestaurant(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
		log.LogHandlerError(logger, errors.New("Неверный формат id ресторана"), http.StatusBadRequest)
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
		log.LogHandlerError(logger, fmt.Errorf("Ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result []byte
	for _, product := range products {
		productData, err := easyjson.Marshal(product)
		if err != nil {
			log.LogHandlerError(logger, fmt.Errorf("Ошибка маршалинга продуктов: %w", err), http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result = append(result, productData...)
	}

	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

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
		log.LogHandlerError(logger, fmt.Errorf("Ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result []byte
	for _, restaurant := range restaurants {
		restaurantData, err := easyjson.Marshal(restaurant)
		if err != nil {
			log.LogHandlerError(logger, fmt.Errorf("Ошибка маршалинга ресторана: %w", err), http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result = append(result, restaurantData...)
	}

	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *RestaurantHandler) RestaurantById(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	vars := mux.Vars(r)
	idStr := vars["id"]

	id := uuid.FromStringOrNil(idStr)
	if id == uuid.Nil {
		log.LogHandlerError(logger, errors.New("Неверный формат id ресторана"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant, err := h.restaurantUsecase.GetById(r.Context(), id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка уровнем ниже (usecase): %w", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := easyjson.Marshal(restaurant)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Не удалось сериализовать данные: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}
