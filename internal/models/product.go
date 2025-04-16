package models

import (
	"html"

	"github.com/satori/uuid"
)

// easyjson:json
type WorkingMode struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// easyjson:json
type DeliveryTime struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// easyjson:json
type Product struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
}

// easyjson:json
type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

// easyjson:json
type RestaurantFull struct {
	Id           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	BannerURL    string       `json:"banner_url"`
	Address      string       `json:"address"`
	Description  string       `json:"description"`
	Rating       float64      `json:"rating"`
	RatingCount  int          `json:"rating_count"`
	WorkingMode  WorkingMode  `json:"working_mode"`
	DeliveryTime DeliveryTime `json:"delivery_time"`
	Tags         []string     `json:"tags"`
	Categories   []Category   `json:"categories"`
}

func (p *Product) Sanitize() {
	p.Name = html.EscapeString(p.Name)
	p.ImageURL = html.EscapeString(p.ImageURL)
}

func (c *Category) Sanitize() {
	c.Name = html.EscapeString(c.Name)
	for i := range c.Products {
		c.Products[i].Sanitize()
	}
}

func (r *RestaurantFull) Sanitize() {
	r.Name = html.EscapeString(r.Name)
	r.BannerURL = html.EscapeString(r.BannerURL)
	r.Address = html.EscapeString(r.Address)
	r.Description = html.EscapeString(r.Description)
	for i := range r.Tags {
		r.Tags[i] = html.EscapeString(r.Tags[i]) 
	}
	for i := range r.Categories {
		r.Categories[i].Sanitize() 
	}
}
