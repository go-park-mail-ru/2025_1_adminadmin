package models

import uuid "github.com/satori/uuid"

type SignUpReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	Login        string    `json:"login"`
	Id           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	UserPic      string    `json:"path"`
	PasswordHash string    `json:"-"`
}
