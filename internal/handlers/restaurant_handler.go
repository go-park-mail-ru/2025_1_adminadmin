package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/gorilla/mux"
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

var defaultRestaurantOptions = RestaurantOptions{
	count:  10,
	offset: 0,
}

type RestaurantOptions struct {
	count  int
	offset int
}

type applyRestaurantOption interface {
	apply(*RestaurantOptions)
}

type funcRestaurantOption struct {
	f func(option *RestaurantOptions)
}

func (fdo *funcRestaurantOption) apply(opt *RestaurantOptions) {
	fdo.f(opt)
}

func newFuncRestaurantOption(f func(option *RestaurantOptions)) *funcRestaurantOption {
	return &funcRestaurantOption{
		f: f,
	}
}

func WithCustomCount(count int) applyRestaurantOption {
	return newFuncRestaurantOption(func(o *RestaurantOptions) {
		if count >= 0 && count <= len(restaurants) {
			o.count = count
		} 
		if count > len(restaurants) {
			o.count = len(restaurants)
		}
		if count < 0 {
			o.count = 0
		}
	})
}

func WithCustomOffset(offset int) applyRestaurantOption {
	return newFuncRestaurantOption(func(o *RestaurantOptions) {
		if offset >= 0 && offset < len(restaurants) {
			o.offset = offset
		} else {
			o.count = 0
		}
	})
}

type Options struct {
	opts RestaurantOptions
}

func NewOptions(opts ...applyRestaurantOption) *Options {
	options := defaultRestaurantOptions
	for _, option := range opts {
		option.apply(&options)
	}
	return &Options{opts: options}

}

func RestaurantList(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	offsetStr := r.URL.Query().Get("offset")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = defaultRestaurantOptions.count
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = defaultRestaurantOptions.offset
	}

	params := NewOptions(WithCustomCount(count), WithCustomOffset(offset))

	w.Header().Set("total", strconv.Itoa(len(restaurants)))
	end := params.opts.offset + params.opts.count
	if end > len(restaurants){
		end = len(restaurants)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(restaurants[params.opts.offset:end])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func RestaurantByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uuid.FromStringOrNil(idStr)
	if id == uuid.Nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found := false
	for _, restaurant := range restaurants {
		if id == restaurant.Id {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(restaurant)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			found = true
			break
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
}
