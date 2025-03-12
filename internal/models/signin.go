package models

type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
