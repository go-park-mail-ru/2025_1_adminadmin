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

//go:embed scripts/select.sql
var selectAll string

type Handler struct {
	db pgxtype.Querier
}

func CreateHandler(p *pgxpool.Pool) *Handler {
	return &Handler{db: p}
}

var restaurants = []models.Restaurant{
	{Id: uuid.NewV4(), Name: "La Piazza", Description: "Итальянская кухня", Type: "Итальянский", Rating: 4.5},
	{Id: uuid.NewV4(), Name: "Sakura", Description: "Японская кухня", Type: "Японский", Rating: 4.7},
	{Id: uuid.NewV4(), Name: "Steak House", Description: "Лучшие стейки в городе", Type: "Американский", Rating: 4.6},
	{Id: uuid.NewV4(), Name: "Bistro Parisien", Description: "Французская кухня", Type: "Французский", Rating: 4.3},
	{Id: uuid.NewV4(), Name: "Taco Loco", Description: "Мексиканская кухня", Type: "Мексиканский", Rating: 4.2},
	{Id: uuid.NewV4(), Name: "Dragon Wok", Description: "Китайская кухня", Type: "Китайский", Rating: 4.4},
	{Id: uuid.NewV4(), Name: "Berlin Döner", Description: "Настоящий немецкий донер", Type: "Немецкий", Rating: 4.1},
	{Id: uuid.NewV4(), Name: "Kebab King", Description: "Лучший кебаб в городе", Type: "Турецкий", Rating: 4.0},
	{Id: uuid.NewV4(), Name: "Green Garden", Description: "Вегетарианская кухня", Type: "Вегетарианский", Rating: 4.8},
	{Id: uuid.NewV4(), Name: "Sea Breeze", Description: "Свежие морепродукты", Type: "Морепродукты", Rating: 4.9},
	{Id: uuid.NewV4(), Name: "La Piazza", Description: "Итальянская кухня", Type: "Итальянский", Rating: 4.5},
	{Id: uuid.NewV4(), Name: "Sakura", Description: "Японская кухня", Type: "Японский", Rating: 4.7},
	{Id: uuid.NewV4(), Name: "Steak House", Description: "Лучшие стейки в городе", Type: "Американский", Rating: 4.6},
	{Id: uuid.NewV4(), Name: "Bistro Parisien", Description: "Французская кухня", Type: "Французский", Rating: 4.3},
	{Id: uuid.NewV4(), Name: "Taco Loco", Description: "Мексиканская кухня", Type: "Мексиканский", Rating: 4.2},
	{Id: uuid.NewV4(), Name: "Dragon Wok", Description: "Китайская кухня", Type: "Китайский", Rating: 4.4},
	{Id: uuid.NewV4(), Name: "Berlin Döner", Description: "Настоящий немецкий донер", Type: "Немецкий", Rating: 4.1},
	{Id: uuid.NewV4(), Name: "Kebab King", Description: "Лучший кебаб в городе", Type: "Турецкий", Rating: 4.0},
	{Id: uuid.NewV4(), Name: "Green Garden", Description: "Вегетарианская кухня", Type: "Вегетарианский", Rating: 4.8},
	{Id: uuid.NewV4(), Name: "Sea Breeze", Description: "Свежие морепродукты", Type: "Морепродукты", Rating: 4.9},
	{Id: uuid.NewV4(), Name: "La Piazza", Description: "Итальянская кухня", Type: "Итальянский", Rating: 4.5},
	{Id: uuid.NewV4(), Name: "Sakura", Description: "Японская кухня", Type: "Японский", Rating: 4.7},
	{Id: uuid.NewV4(), Name: "Steak House", Description: "Лучшие стейки в городе", Type: "Американский", Rating: 4.6},
	{Id: uuid.NewV4(), Name: "Bistro Parisien", Description: "Французская кухня", Type: "Французский", Rating: 4.3},
	{Id: uuid.NewV4(), Name: "Taco Loco", Description: "Мексиканская кухня", Type: "Мексиканский", Rating: 4.2},
	{Id: uuid.NewV4(), Name: "Dragon Wok", Description: "Китайская кухня", Type: "Китайский", Rating: 4.4},
	{Id: uuid.NewV4(), Name: "Berlin Döner", Description: "Настоящий немецкий донер", Type: "Немецкий", Rating: 4.1},
	{Id: uuid.NewV4(), Name: "Kebab King", Description: "Лучший кебаб в городе", Type: "Турецкий", Rating: 4.0},
	{Id: uuid.NewV4(), Name: "Green Garden", Description: "Вегетарианская кухня", Type: "Вегетарианский", Rating: 4.8},
	{Id: uuid.NewV4(), Name: "Sea Breeze", Description: "Свежие морепродукты", Type: "Морепродукты", Rating: 4.9},
	{Id: uuid.NewV4(), Name: "La Piazza", Description: "Итальянская кухня", Type: "Итальянский", Rating: 4.5},
	{Id: uuid.NewV4(), Name: "Sakura", Description: "Японская кухня", Type: "Японский", Rating: 4.7},
	{Id: uuid.NewV4(), Name: "Steak House", Description: "Лучшие стейки в городе", Type: "Американский", Rating: 4.6},
	{Id: uuid.NewV4(), Name: "Bistro Parisien", Description: "Французская кухня", Type: "Французский", Rating: 4.3},
	{Id: uuid.NewV4(), Name: "Taco Loco", Description: "Мексиканская кухня", Type: "Мексиканский", Rating: 4.2},
	{Id: uuid.NewV4(), Name: "Dragon Wok", Description: "Китайская кухня", Type: "Китайский", Rating: 4.4},
	{Id: uuid.NewV4(), Name: "Berlin Döner", Description: "Настоящий немецкий донер", Type: "Немецкий", Rating: 4.1},
	{Id: uuid.NewV4(), Name: "Kebab King", Description: "Лучший кебаб в городе", Type: "Турецкий", Rating: 4.0},
	{Id: uuid.NewV4(), Name: "Green Garden", Description: "Вегетарианская кухня", Type: "Вегетарианский", Rating: 4.8},
	{Id: uuid.NewV4(), Name: "Sea Breeze", Description: "Свежие морепродукты", Type: "Морепродукты", Rating: 4.9},
	{Id: uuid.NewV4(), Name: "La Piazza", Description: "Итальянская кухня", Type: "Итальянский", Rating: 4.5},
	{Id: uuid.NewV4(), Name: "Sakura", Description: "Японская кухня", Type: "Японский", Rating: 4.7},
	{Id: uuid.NewV4(), Name: "Steak House", Description: "Лучшие стейки в городе", Type: "Американский", Rating: 4.6},
	{Id: uuid.NewV4(), Name: "Bistro Parisien", Description: "Французская кухня", Type: "Французский", Rating: 4.3},
	{Id: uuid.NewV4(), Name: "Taco Loco", Description: "Мексиканская кухня", Type: "Мексиканский", Rating: 4.2},
	{Id: uuid.NewV4(), Name: "Dragon Wok", Description: "Китайская кухня", Type: "Китайский", Rating: 4.4},
	{Id: uuid.NewV4(), Name: "Berlin Döner", Description: "Настоящий немецкий донер", Type: "Немецкий", Rating: 4.1},
	{Id: uuid.NewV4(), Name: "Kebab King", Description: "Лучший кебаб в городе", Type: "Турецкий", Rating: 4.0},
	{Id: uuid.NewV4(), Name: "Green Garden", Description: "Вегетарианская кухня", Type: "Вегетарианский", Rating: 4.8},
	{Id: uuid.NewV4(), Name: "Sea Breeze", Description: "Свежие морепродукты", Type: "Морепродукты", Rating: 4.9},
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

	restaurants = []models.Restaurant{}

	for rows.Next() {
		var restaurant models.Restaurant
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Type, &restaurant.Rating)
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

	var restaurant models.Restaurant
	err := h.db.QueryRow(r.Context(), "SELECT id, name, description, type, rating FROM restaurants WHERE id = $1", id).Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Type, &restaurant.Rating)
	if err != nil {
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
