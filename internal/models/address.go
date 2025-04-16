package models

import (
	"html"

	"github.com/satori/uuid"
)

// easyjson:json
type Address struct {
	Id      uuid.UUID `json:"id"`
	Address string    `json:"address"`
	UserId  uuid.UUID `json:"user_id"`
}

type DeleteAddressReq struct {
	Id string `json:"id"`
}

func (a *Address) Sanitize() {
	a.Address = html.EscapeString(a.Address)
}
