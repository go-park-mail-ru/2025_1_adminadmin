package models

type SignUpReq struct {
	Login       string `json:"login"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
