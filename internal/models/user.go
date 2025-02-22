package models

import uuid "github.com/satori/uuid"

type User struct {
	Login        string    `json:"login"`
	Id           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	UserPic      string    `json:"path"`
	PasswordHash string    `json:"-"`
}
