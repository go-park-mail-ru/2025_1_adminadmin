package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

var users = make(map[string]models.User)

func hashSHA256(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, exists := users[req.Login]
	if !exists || user.PasswordHash != hashSHA256(req.Password) {
		http.Error(w, "Неверные данные", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, exists := users[req.Login]
	if exists {
		http.Error(w, "Данный логин уже занят", http.StatusConflict)
		return
	}

	hashedPassword := hashSHA256(req.Password)

	user := models.User{
		Login:        req.Login,
		Id:           uuid.NewV4(),
		Description:  "New User",
		UserPic:      "default.png",
		PasswordHash: hashedPassword,
	}
	users[req.Login] = user
	w.WriteHeader(http.StatusCreated)
}
