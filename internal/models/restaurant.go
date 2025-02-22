package models

import uuid "github.com/satori/uuid"

type Restaurant struct {
	Id           uuid.UUID `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Rating      float64 `json:"rating"`
}
