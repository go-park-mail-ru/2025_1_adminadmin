package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/satori/uuid"
)

const maxRequestBodySize = 10 << 20

var allowedMimeTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

type AuthHandler struct {
	uc     auth.AuthUsecase
	secret string
}

func CreateAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc, secret: os.Getenv("JWT_SECRET")}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignInReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignIn(r.Context(), req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrUserNotFound:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidCredentials:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusUnauthorized)
		default:
			log.LogHandlerError(logger, fmt.Errorf("Неизвестная ошибка: %w", err), http.StatusInternalServerError)
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
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	user, token, csrfToken, err := h.uc.SignUp(r.Context(), req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrInvalidPassword:
			log.LogHandlerError(logger, fmt.Errorf("Неправильный логин или пароль: %w", err), http.StatusBadRequest)
			utils.SendError(w, "Неправильный логин или пароль", http.StatusBadRequest)
		case auth.ErrInvalidName:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidPhone:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrCreatingUser:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		default:
			log.LogHandlerError(logger, fmt.Errorf("Неизвестная ошибка: %w", err), http.StatusInternalServerError)
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
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Check(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

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

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
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
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil || cookie.Value == "" {
		log.LogHandlerError(logger, fmt.Errorf("Пользователь уже разлогинен: %w", err), http.StatusBadRequest)
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
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("Токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ТОшибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("Недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	var updateData models.UpdateUserReq
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.uc.UpdateUser(r.Context(), login, updateData)
	if err != nil {
		switch err {
		case auth.ErrInvalidPassword:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidName:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidPhone:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSamePassword:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSameName:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrSamePhone:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		default:
			log.LogHandlerError(logger, fmt.Errorf("Ошибка обновления данных пользователя: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "Ошибка обновления данных пользователя", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) UpdateUserPic(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("Токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ТОшибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("Недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)

	if err := r.ParseMultipartForm(maxRequestBodySize); err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			log.LogHandlerError(logger, fmt.Errorf("Превышен допустимый размер файла: %w", err), http.StatusRequestEntityTooLarge)
			utils.SendError(w, "Превышен допустимый размер файла", http.StatusRequestEntityTooLarge)
		} else {
			log.LogHandlerError(logger, fmt.Errorf("Невозможно обработать файл: %w", err), http.StatusBadRequest)
			utils.SendError(w, "Невозможно обработать файл", http.StatusBadRequest)
		}
		return
	}

	defer func() {
		if r.MultipartForm != nil {
			r.MultipartForm.RemoveAll()
		}
	}()

	file, _, err := r.FormFile("user_pic")
	if err != nil {
		utils.SendError(w, "Файл не найден в запросе", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка при чтении файла: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка при чтении файла", http.StatusBadRequest)
		return
	}
	file.Seek(0, io.SeekStart)

	mimeType := http.DetectContentType(buffer)
	if _, ok := allowedMimeTypes[mimeType]; !ok {
		log.LogHandlerError(logger, fmt.Errorf("Недопустимый формат файла: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Недопустимый формат файла.", http.StatusBadRequest)
		return
	}

	ext := allowedMimeTypes[mimeType]

	updatedUser, err := h.uc.UpdateUserPic(r.Context(), login, file, ext)
	if err != nil {
		switch err {
		case auth.ErrUserNotFound:
			log.LogHandlerError(logger, err, http.StatusNotFound)
			utils.SendError(w, err.Error(), http.StatusNotFound)
		case auth.ErrBasePath:
			log.LogHandlerError(logger, err, http.StatusInternalServerError)
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		case auth.ErrFileCreation, auth.ErrFileSaving, auth.ErrFileDeletion:
			log.LogHandlerError(logger, fmt.Errorf("Ошибка при работе с файлом: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "Ошибка при работе с файлом", http.StatusInternalServerError)
		default:
			log.LogHandlerError(logger, fmt.Errorf("Ошибка при обновлении аватарки: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "Ошибка при обновлении аватарки", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) GetUserAddresses(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("Токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("Ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("Недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	addresses, err := h.uc.GetUserAddresses(r.Context(), login)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(addresses); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "Ошибка формирования JSON", http.StatusInternalServerError)
	}
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteAddress(r.Context(), address.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Info("Successful")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("Токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "Токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("Ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	idStr, ok := jwtUtils.GetIdFromJWT(JWTStr, claims, h.secret)
	if !ok || idStr == "" {
		log.LogHandlerError(logger, errors.New("Недействительный токен: id отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "Недействительный токен: id отсутствует", http.StatusUnauthorized)
		return
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var address models.Address
	err = json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	address.UserId = id

	err = h.uc.AddAddress(r.Context(), address)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}
