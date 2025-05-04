package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/search"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

type SearchHandler struct {
	uc search.SearchUsecase
}

func NewSearchHandler(uc search.SearchUsecase) *SearchHandler {
	return &SearchHandler{
		uc: uc,
	}
}

func (h *SearchHandler) SearchRestaurantWithProducts(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	query, _ := url.QueryUnescape(r.URL.Query().Get("query"))
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	countInt, offsetInt := 30, 0
	if count != "" {
		fmt.Sscanf(count, "%d", &countInt)
	}
	if offset != "" {
		fmt.Sscanf(offset, "%d", &offsetInt)
	}

	restaurants, err := h.uc.SearchRestaurantWithProducts(r.Context(), query, countInt, offsetInt)
	if err != nil {
		utils.SendError(w, "Ошибка поиска ресторанов с продуктами", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(restaurants)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать результат", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *SearchHandler) SearchProductsInRestaurant(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))
	vars := mux.Vars(r)
	restaurantIDStr := vars["id"]
	restaurantID := uuid.FromStringOrNil(restaurantIDStr)
	if restaurantID == uuid.Nil {
		log.LogHandlerError(logger, errors.New("неверный формат id ресторана"), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := r.URL.Query().Get("query")
	productCategories, err := h.uc.SearchProductsInRestaurant(r.Context(), restaurantID, query)
	if err != nil {
		utils.SendError(w, "Ошибка поиска продуктов", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(productCategories)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Не удалось сериализовать корзину", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
