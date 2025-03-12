package models

import uuid "github.com/satori/uuid"

type User struct {
	Login        string    `json:"login"`
	PhoneNumber  string    `json:"phone_number"`
	Id           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Description  string    `json:"description"`
	UserPic      string    `json:"path"`
	PasswordHash []byte    `json:"-"`
}
