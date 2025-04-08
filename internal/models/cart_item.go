package models

import (
	"html"

	"github.com/satori/uuid"
)

// easyjson:json
type CartItem struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
	Amount   int       `json:"amount"`
}

// easyjson:json
type Cart struct {
	Id       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	CartItems []CartItem `json:"products"`
}

func (a *CartItem) Sanitize() {
	a.Name = html.EscapeString(a.Name)
	a.ImageURL = html.EscapeString(a.ImageURL)
}

func (a *Cart) Sanitize() {
	a.Name = html.EscapeString(a.Name)
}