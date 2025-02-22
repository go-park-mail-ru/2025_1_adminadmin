package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	uuid "github.com/satori/uuid"
)

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
}

func RestaurantList(w http.ResponseWriter, r *http.Request) {

	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 0 || count > len(restaurants) {
		count = len(restaurants)
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	w.Header().Set("total", strconv.Itoa(len(restaurants)))

	if offset > len(restaurants){
		err = json.NewEncoder(w).Encode([]models.Restaurant{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	
	end := offset + count
	if end > len(restaurants){
		end = len(restaurants)
	}

	err = json.NewEncoder(w).Encode(restaurants[offset:end])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
