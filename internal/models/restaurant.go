package models

import (
	"html"

	uuid "github.com/satori/uuid"
)

// easyjson:json
type Restaurant struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	ImageURL    string    `json:"image_url"`
}

// easyjson:json
type RestaurantSearch struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	BannerURL   string          `json:"image_url"`
	Address     string          `json:"address"`
	Rating      float64         `json:"rating"`
	RatingCount float64         `json:"rating_count"`
	Description string          `json:"description"`
	Products    []ProductSearch `json:"products"`
}

// easyjson:json
type ProductCategory struct {
	Name     string          `json:"name"`
	Products []ProductSearch `json:"products"`
}

// easyjson:json
type ProductSearch struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	ImageURL     string    `json:"image_url"`
	Weight       int       `json:"weight"`
	Category     string    `json:"category"`
}

func (r *Restaurant) Sanitize() {
	r.Name = html.EscapeString(r.Name)
	r.Description = html.EscapeString(r.Description)
	r.ImageURL = html.EscapeString(r.ImageURL)
}

func (r *RestaurantSearch) Sanitize() {
	r.Name = html.EscapeString(r.Name)
	r.Description = html.EscapeString(r.Description)
	r.BannerURL = html.EscapeString(r.BannerURL)
	r.Address = html.EscapeString(r.Address)
}

func (r *ProductSearch) Sanitize() {
	r.Name = html.EscapeString(r.Name)
	r.Category = html.EscapeString(r.Category)
	r.ImageURL = html.EscapeString(r.ImageURL)
}
