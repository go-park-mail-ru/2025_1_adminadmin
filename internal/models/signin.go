package models

import "html"

// easyjson:json
type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *SignInReq) Sanitize() {
	s.Login = html.EscapeString(s.Login)
	s.Password = html.EscapeString(s.Password)
}
