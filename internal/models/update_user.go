package models

import "html"

// easyjson:json
type UpdateUserReq struct {
	Description string `json:"description,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
}

func (u *UpdateUserReq) Sanitize() {
	u.Description = html.EscapeString(u.Description)
	u.FirstName = html.EscapeString(u.FirstName)
	u.LastName = html.EscapeString(u.LastName)
	u.PhoneNumber = html.EscapeString(u.PhoneNumber)
	u.Password = html.EscapeString(u.Password)
}
