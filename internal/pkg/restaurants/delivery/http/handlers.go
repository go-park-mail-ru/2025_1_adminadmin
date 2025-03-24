package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/restaurants/usecase"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/utils/send_error"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

type RestaurantHandler struct {
	restaurantUsecase usecase.RestaurantUsecase
}

func NewRestaurantHandler(ru usecase.RestaurantUsecase) *RestaurantHandler {
	return &RestaurantHandler{restaurantUsecase: ru}
}

func (h *RestaurantHandler) GetProductsByRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *RestaurantHandler) RestaurantList(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(restaurants)
	if err != nil {
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

func (h *RestaurantHandler) RestaurantById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id := uuid.FromStringOrNil(idStr)
	if id == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant, err := h.restaurantUsecase.GetById(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(restaurant)
	if err != nil {
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
