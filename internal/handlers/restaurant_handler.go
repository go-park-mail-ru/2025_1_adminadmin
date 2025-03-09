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
		{Id: uuid.NewV4(), Name: "Pasta Paradise", Description: "Паста и пицца", Type: "Итальянский", Rating: 4.5},
		{Id: uuid.NewV4(), Name: "Sushi Master", Description: "Суши и сашими", Type: "Японский", Rating: 4.7},
		{Id: uuid.NewV4(), Name: "Burger Joint", Description: "Бургеры и картофель фри", Type: "Американский", Rating: 4.6},
		{Id: uuid.NewV4(), Name: "Le Petit Bistro", Description: "Французские деликатесы", Type: "Французский", Rating: 4.3},
		{Id: uuid.NewV4(), Name: "Taco Fiesta", Description: "Мексиканские тако", Type: "Мексиканский", Rating: 4.2},
		{Id: uuid.NewV4(), Name: "Golden Wok", Description: "Китайские блюда", Type: "Китайский", Rating: 4.4},
		{Id: uuid.NewV4(), Name: "Munich Haus", Description: "Немецкие колбаски", Type: "Немецкий", Rating: 4.1},
		{Id: uuid.NewV4(), Name: "Kebab Palace", Description: "Турецкие кебабы", Type: "Турецкий", Rating: 4.0},
		{Id: uuid.NewV4(), Name: "Veggie Delight", Description: "Вегетарианские блюда", Type: "Вегетарианский", Rating: 4.8},
		{Id: uuid.NewV4(), Name: "Ocean's Catch", Description: "Морепродукты и рыба", Type: "Морепродукты", Rating: 4.9},
		{Id: uuid.NewV4(), Name: "Ristorante Roma", Description: "Итальянские деликатесы", Type: "Итальянский", Rating: 4.5},
		{Id: uuid.NewV4(), Name: "Tokyo Sushi", Description: "Японские суши", Type: "Японский", Rating: 4.7},
		{Id: uuid.NewV4(), Name: "BBQ Pit", Description: "Американский барбекю", Type: "Американский", Rating: 4.6},
		{Id: uuid.NewV4(), Name: "Café de Paris", Description: "Французская выпечка", Type: "Французский", Rating: 4.3},
		{Id: uuid.NewV4(), Name: "Taco Express", Description: "Мексиканские закуски", Type: "Мексиканский", Rating: 4.2},
		{Id: uuid.NewV4(), Name: "Peking Garden", Description: "Китайские деликатесы", Type: "Китайский", Rating: 4.4},
		{Id: uuid.NewV4(), Name: "Bavarian Inn", Description: "Немецкие блюда", Type: "Немецкий", Rating: 4.1},
		{Id: uuid.NewV4(), Name: "Kebab Express", Description: "Турецкие закуски", Type: "Турецкий", Rating: 4.0},
		{Id: uuid.NewV4(), Name: "Veggie Heaven", Description: "Вегетарианские деликатесы", Type: "Вегетарианский", Rating: 4.8},
		{Id: uuid.NewV4(), Name: "Seafood Haven", Description: "Морепродукты и устрицы", Type: "Морепродукты", Rating: 4.9},
		{Id: uuid.NewV4(), Name: "Pasta Factory", Description: "Итальянские паста и соусы", Type: "Итальянский", Rating: 4.5},
		{Id: uuid.NewV4(), Name: "Sakura Sushi", Description: "Японские суши и роллы", Type: "Японский", Rating: 4.7},
		{Id: uuid.NewV4(), Name: "Steak & Ale", Description: "Американские стейки", Type: "Американский", Rating: 4.6},
		{Id: uuid.NewV4(), Name: "Parisian Café", Description: "Французские десерты", Type: "Французский", Rating: 4.3},
		{Id: uuid.NewV4(), Name: "Taco Time", Description: "Мексиканские тако и буррито", Type: "Мексиканский", Rating: 4.2},
		{Id: uuid.NewV4(), Name: "Dragon Palace", Description: "Китайские деликатесы", Type: "Китайский", Rating: 4.4},
		{Id: uuid.NewV4(), Name: "Berliner Haus", Description: "Немецкие блюда", Type: "Немецкий", Rating: 4.1},
		{Id: uuid.NewV4(), Name: "Kebab House", Description: "Турецкие кебабы", Type: "Турецкий", Rating: 4.0},
		{Id: uuid.NewV4(), Name: "Veggie World", Description: "Вегетарианские блюда", Type: "Вегетарианский", Rating: 4.8},
		{Id: uuid.NewV4(), Name: "Seafood Cove", Description: "Морепродукты и рыба", Type: "Морепродукты", Rating: 4.9},
		{Id: uuid.NewV4(), Name: "Pasta Palace", Description: "Итальянские паста и пицца", Type: "Итальянский", Rating: 4.5},
		{Id: uuid.NewV4(), Name: "Sushi World", Description: "Японские суши и сашими", Type: "Японский", Rating: 4.7},
		{Id: uuid.NewV4(), Name: "Burger Barn", Description: "Американские бургеры", Type: "Американский", Rating: 4.6},
		{Id: uuid.NewV4(), Name: "Parisian Bistro", Description: "Французские деликатесы", Type: "Французский", Rating: 4.3},
		{Id: uuid.NewV4(), Name: "Taco Town", Description: "Мексиканские тако", Type: "Мексиканский", Rating: 4.2},
		{Id: uuid.NewV4(), Name: "Dragon Express", Description: "Китайские блюда", Type: "Китайский", Rating: 4.4},
		{Id: uuid.NewV4(), Name: "Bavarian House", Description: "Немецкие колбаски", Type: "Немецкий", Rating: 4.1},
		{Id: uuid.NewV4(), Name: "Kebab Corner", Description: "Турецкие кебабы", Type: "Турецкий", Rating: 4.0},
		{Id: uuid.NewV4(), Name: "Veggie Spot", Description: "Вегетарианские блюда", Type: "Вегетарианский", Rating: 4.8},
		{Id: uuid.NewV4(), Name: "Seafood Shack", Description: "Морепродукты и рыба", Type: "Морепродукты", Rating: 4.9},
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

	rows, err := h.db.Query(r.Context(), selectById, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var restaurant models.Restaurant
	if rows.Next() {
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Description, &restaurant.Type, &restaurant.Rating)
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
