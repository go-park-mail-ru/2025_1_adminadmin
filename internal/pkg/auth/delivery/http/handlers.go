package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/utils/send_error"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct {
	uc auth.AuthUsecase
}

func CreateAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignIn(r.Context(), req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrUserNotFound:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidCredentials:
			utils.SendError(w, err.Error(), http.StatusUnauthorized)
		case auth.ErrGeneratingToken:
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		default:
			utils.SendError(w, "Неизвестная ошибка", http.StatusInternalServerError)
		}
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignUp(r.Context(), req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidPassword:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidName:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidPhone:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrCreatingUser:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrGeneratingToken:
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		default:
			utils.SendError(w, "Неизвестная ошибка", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Check(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.uc.Check(r.Context(), login)

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil || cookie.Value == "" {
		utils.SendError(w, "Пользователь уже разлогинен", http.StatusBadRequest)
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

func (h *AuthHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(JWTStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		utils.SendError(w, "Недействительный токен", http.StatusUnauthorized)
		return
	}

	login, ok := claims["login"].(string)
	if !ok || login == "" {
		utils.SendError(w, "Недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	var updateData models.UpdateUserReq
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.uc.UpdateUser(r.Context(), login, updateData)
	if err != nil {
		switch err {
		case auth.ErrInvalidPassword:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidName:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidPhone:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSamePassword:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSameName:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSamePhone:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSameDescription:
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrUserNotFound:
			utils.SendError(w, err.Error(), http.StatusNotFound)
		default:
			utils.SendError(w, "Ошибка обновления данных пользователя", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) UpdateUserPic(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	tokenString := cookie.Value

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		utils.SendError(w, "Недействительный токен", http.StatusUnauthorized)
		return
	}

	login, ok := claims["login"].(string)
	if !ok || login == "" {
		utils.SendError(w, "Недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	const maxRequestBodySize = 10 << 20 
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)

	if err := r.ParseMultipartForm(maxRequestBodySize); err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			utils.SendError(w, "Превышен допустимый размер файла", http.StatusRequestEntityTooLarge)
		} else {
			utils.SendError(w, "Невозможно обработать файл", http.StatusBadRequest)
		}
		return
	}

	file, handler, err := r.FormFile("user_pic")
	if err != nil {
		utils.SendError(w, "Файл не найден в запросе", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512) 
	if _, err := file.Read(buffer); err != nil {
		utils.SendError(w, "Ошибка при чтении файла", http.StatusBadRequest)
		return
	}
	file.Seek(0, io.SeekStart) 

	mimeType := http.DetectContentType(buffer)
	if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/webp" {
		utils.SendError(w, "Недопустимый формат файла. Разрешены только JPEG, PNG и WEBP", http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(handler.Filename)

	updatedUser, err := h.uc.UpdateUserPic(r.Context(), login, file, ext)
	if err != nil {
		switch err {
		case auth.ErrUserNotFound:
			utils.SendError(w, err.Error(), http.StatusNotFound)
		case auth.ErrBasePath:
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		case auth.ErrFileCreation, auth.ErrFileSaving, auth.ErrFileDeletion:
			utils.SendError(w, "Ошибка при работе с файлом", http.StatusInternalServerError)
		default:
			utils.SendError(w, "Ошибка при обновлении аватарки", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}