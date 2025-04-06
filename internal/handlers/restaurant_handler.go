package handlers

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/uuid"
)

//go:embed scripts/select_all_restaurants.sql
var selectAll string

//go:embed scripts/select_restaurant_by_id.sql
var selectById string

type Handler struct {
	db pgxtype.Querier
}

func CreateHandler(p *pgxpool.Pool) *Handler {
	return &Handler{db: p}
}

func (h *Handler) RestaurantList(w http.ResponseWriter, r *http.Request) {
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

	rows, err := h.db.Query(r.Context(), selectAll, count, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	restaurants := []models.Restaurant{}

	for rows.Next() {
		var restaurant models.Restaurant
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Rating)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		restaurants = append(restaurants, restaurant)
	}

	if err = rows.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

func (h *Handler) RestaurantByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.FromStringOrNil(idStr)
	if id == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := h.db.Query(r.Context(), selectById, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var restaurant models.Restaurant
	if rows.Next() {
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Rating)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
