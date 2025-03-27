package models

import (
	"html"

	"github.com/satori/uuid"
)

// easyjson:json
type Address struct {
	Id           uuid.UUID `json:"id"`
	Address string    `json:"address"`
	UserId       uuid.UUID `json:"user_id"`
}

func (a *Address) Sanitize() {
	a.Address = html.EscapeString(a.Address)
}
