package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

var users = make(map[string]models.User)

func hashSHA256(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func validLogin(login string) bool {
	if len(login) < 3 || len(login) > 20 {
		return false
	}
	for _, char := range login {
		isLatin := !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z')
		if isLatin && !unicode.IsDigit(char) && char != '_' && char != '-' {
			return false
		}
	}
	return true
}

func validPassword(password string) bool {
	var up, low, digit, special bool

	if len(password) < 8 {
		return false
	}

	for _, char := range password {

		switch {
		case unicode.IsUpper(char):
			up = true
		case unicode.IsLower(char):
			low = true
		case unicode.IsDigit(char):
			digit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		default:
			return false
		}
	}

	return up && low && digit && special
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validLogin(req.Login) {
		http.Error(w, "Неверный формат логина", http.StatusBadRequest)
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

	if !validLogin(req.Login) {
		http.Error(w, "Неверный формат логина", http.StatusBadRequest)
		return
	}

	if !validPassword(req.Password) {
		http.Error(w, "Неверный формат пароля", http.StatusBadRequest)
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
