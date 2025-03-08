package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/satori/uuid"
	"golang.org/x/crypto/argon2"
)

var users = make(map[string]models.User)

func hashPassword(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func checkPassword(passHash []byte, plainPassword string) bool {
	salt := make([]byte, 8)
	copy(salt, passHash[:8])
	userPassHash := hashPassword(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}

var allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

func validLogin(login string) bool {
	if len(login) < 3 || len(login) > 20 {
		return false
	}
	for _, char := range login {
		if !strings.Contains(allowedChars, string(char)) {
			return false
		}
	}
	return true
}

func validPassword(password string) bool {
	var up, low, digit, special bool

	if len(password) < 8 || len(password) > 25 {
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

func generateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
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

	var user models.User
	err = h.db.QueryRow(r.Context(), 
		"SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash FROM users WHERE login = $1", 
		req.Login).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Description, &user.UserPic, &user.PasswordHash)
 
	if err != nil || !checkPassword(user.PasswordHash, req.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := generateToken(user.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})

	csrfToken := uuid.NewV4().String()

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", csrfToken)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
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

	if req.FirstName == "" || req.LastName == "" {
		http.Error(w, "Имя и фамилия обязательны", http.StatusBadRequest)
		return
	}

	if req.PhoneNumber == "" {
		http.Error(w, "Телефон обязателен", http.StatusBadRequest)
		return
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPassword := hashPassword(salt, req.Password)

	userID := uuid.NewV4()
	_, err = h.db.Exec(r.Context(), 
		"INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		userID, req.Login, req.FirstName, req.LastName, req.PhoneNumber, "", "default.jpg", hashedPassword)

	if err != nil {
		http.Error(w, "Ошибка сохранения пользователя", http.StatusBadRequest)
		return
	}

	user := models.User{
		Id:           userID,
		Login:        req.Login,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Description:  "",
		UserPic:      "default.jpg",
		PasswordHash: hashedPassword,
	}

	token, err := generateToken(user.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	})

	csrfToken := uuid.NewV4().String()

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("X-CSRF-Token", csrfToken)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Check(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	csrfToken := r.Header.Get("X-CSRF-Token")
	if csrfToken == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	csrfCookie, err := r.Cookie("CSRF-Token")
	if err != nil || csrfCookie.Value != csrfToken {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if token == nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
