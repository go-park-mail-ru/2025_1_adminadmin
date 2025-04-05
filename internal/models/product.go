package models

import (
	"html"

	"github.com/satori/uuid"
)

// easyjson:json
type Product struct {
	Id           uuid.UUID `json:"id"`
	RestaurantId uuid.UUID `json:"-"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	ImageURL     string    `json:"image_url"`
	Weight       int       `json:"weight"`
	Amount       int       `json:"amount"`
}

func (p *Product) Sanitize() {
	p.Name = html.EscapeString(p.Name)
}
			