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
}

func (r *Restaurant) Sanitize() {
	r.Name = html.EscapeString(r.Name)
	r.Description = html.EscapeString(r.Description)
}
