package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v", req)
	user := models.User{
		Login:       "test",
		Id:          uuid.FromStringOrNil("8b6e19d0-833e-45c6-a5c4-cff9cf4d295b"),
		Description: "dsf",
		UserPic:     "default.png",
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v", req)
	user := models.User{
		Login:       "test",
		Id:          uuid.FromStringOrNil("8b6e19d0-833e-45c6-a5c4-cff9cf4d295b"),
		Description: "dsf",
		UserPic:     "default.png",
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
