package models

import "github.com/satori/uuid"

type Order struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	Status        string    `json:"status"`
	AddressId     uuid.UUID `json:"address_id"`
	OrderProducts string    `json:"order_products"`
}
