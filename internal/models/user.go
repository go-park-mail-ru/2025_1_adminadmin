package models

import (
	"html"

	uuid "github.com/satori/uuid"
)

// easyjson:json
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

func (u *User) Sanitize() {
	u.Login = html.EscapeString(u.Login)
	u.PhoneNumber = html.EscapeString(u.PhoneNumber)
	u.FirstName = html.EscapeString(u.FirstName)
	u.LastName = html.EscapeString(u.LastName)
	u.Description = html.EscapeString(u.Description)
}
