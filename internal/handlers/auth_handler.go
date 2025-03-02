package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"unicode"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang-jwt/jwt"
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

func generateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		PhoneNumber:  req.PhoneNumber,
		Description:  "New User",
		UserPic:      "default.png",
		PasswordHash: hashedPassword,
	}
	users[req.Login] = user

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
