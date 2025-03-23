package models

import "html"

// easyjson:json
type SignUpReq struct {
	Login       string `json:"login"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (s *SignUpReq) Sanitize() {
	s.Login = html.EscapeString(s.Login)
	s.FirstName = html.EscapeString(s.FirstName)
	s.LastName = html.EscapeString(s.LastName)
	s.PhoneNumber = html.EscapeString(s.PhoneNumber)
	s.Password = html.EscapeString(s.Password)
}
