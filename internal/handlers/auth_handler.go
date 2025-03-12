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

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

const (
	minNameLength  = 2
	maxNameLength  = 25
	minPhoneLength = 7
	maxPhoneLength = 15
	maxLoginLength = 20
	minLoginLength = 3
	minPassLength  = 8
	maxPassLength  = 25
)

func isValidName(name string) bool {
	if len(name) < minNameLength || len(name) > maxNameLength {
		return false
	}
	allowedRunes := "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	for _, r := range name {
		if !compareRune(allowedRunes, r) {
			return false
		}
	}
	return true
}

func isValidPhone(phone string) bool {
	if len(phone) < minPhoneLength || len(phone) > maxPhoneLength {
		return false
	}
	for _, r := range phone {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func compareRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

func validLogin(login string) bool {
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
	if len(login) < minLoginLength || len(login) > maxLoginLength {
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

	if len(password) < minPassLength || len(password) > maxPassLength {
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
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	if !validLogin(req.Login) {
		sendError(w, "Неверный формат логина", http.StatusBadRequest)
		return
	}
	var user models.User
	rows, err := h.db.Query(r.Context(), "SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash FROM users WHERE login = $1", req.Login)
	if err != nil {
		sendError(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		sendError(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.PhoneNumber,
		&user.Description, &user.UserPic, &user.PasswordHash)
	if err != nil {
		sendError(w, "Ошибка чтения данных пользователя", http.StatusInternalServerError)
		return
	}

	user.Login = req.Login

	if !checkPassword(user.PasswordHash, req.Password) {
		sendError(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(user.Login)
	if err != nil {
		sendError(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})

	csrfToken := uuid.NewV4().String()
	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", csrfToken)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		sendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validLogin(req.Login) {
		sendError(w, "Неверный формат логина", http.StatusBadRequest)
		return
	}

	if !validPassword(req.Password) {
		sendError(w, "Неверный формат пароля", http.StatusBadRequest)
		return
	}

	if !isValidName(req.FirstName) || !isValidName(req.LastName) {
		sendError(w, "Имя и фамилия должны содержать только русские буквы и быть от 2 до 25 символов", http.StatusBadRequest)
		return
	}

	if !isValidPhone(req.PhoneNumber) {
		sendError(w, "Некорректный номер телефона", http.StatusBadRequest)
		return
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPassword := hashPassword(salt, req.Password)

	userID := uuid.NewV4()
	_, err = h.db.Exec(r.Context(),
		"INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		userID, req.Login, req.FirstName, req.LastName, req.PhoneNumber, "", "default.png", hashedPassword)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			sendError(w, "Данный логин уже занят", http.StatusConflict)
			return
		}
		sendError(w, "Ошибка сохранения пользователя", http.StatusBadRequest)
		return
	}

	user := models.User{
		Id:           userID,
		Login:        req.Login,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Description:  "",
		UserPic:      "default.png",
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
		Path: "/",
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

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	cookieCSRF, err := r.Cookie("CSRF-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headerCSRF := r.Header.Get("X-CSRF-Token")
	if cookieCSRF.Value == "" || headerCSRF == "" || cookieCSRF.Value != headerCSRF {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cookieJWT, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	JWTStr := cookieJWT.Value

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(JWTStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	login, ok := claims["login"].(string)
	if !ok || login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rows, err := h.db.Query(r.Context(), "SELECT id, first_name, last_name, phone_number, description, user_pic FROM users WHERE login = $1", login)
	if err != nil {
		sendError(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Description, &user.UserPic)
		if err != nil {
			sendError(w, "Ошибка чтения данных пользователя", http.StatusInternalServerError)
			return
		}
		user.Login = login
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil || cookie.Value == "" {
		sendError(w, "Пользователь уже разлогинен", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}
