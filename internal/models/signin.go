package models

import "html"

// easyjson:json
type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// easyjson:json
type Check2fa struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (s *SignInReq) Sanitize() {
	s.Login = html.EscapeString(s.Login)
	s.Password = html.EscapeString(s.Password)
}
