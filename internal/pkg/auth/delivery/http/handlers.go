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
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/mailru/easyjson"
	"github.com/satori/uuid"
)

const maxRequestBodySize = 10 << 20

var allowedMimeTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
}

type AuthHandler struct {
	client gen.AuthServiceClient
	secret string
}

func CreateAuthHandler(client gen.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: client, secret: os.Getenv("JWT_SECRET")}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignInReq
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	req.Sanitize()

	user, err := h.client.SignIn(r.Context(), &gen.SignInRequest{
		Login:    req.Login,
		Password: req.Password})

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrUserNotFound:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		case auth.ErrInvalidCredentials:
			log.LogHandlerError(logger, err, http.StatusUnauthorized)
			utils.SendError(w, err.Error(), http.StatusUnauthorized)
		default:
			log.LogHandlerError(logger, fmt.Errorf("неизвестная ошибка: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "неизвестная ошибка", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    user.Token,
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    user.CsrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", user.CsrfToken)

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	var req models.SignUpReq
	err := easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	req.Sanitize()

	user, err := h.client.SignUp(r.Context(), &gen.SignUpRequest{
		Login:       req.Login,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrInvalidPassword:
			log.LogHandlerError(logger, fmt.Errorf("неправильный логин или пароль: %w", err), http.StatusBadRequest)
			utils.SendError(w, "неправильный логин или пароль", http.StatusBadRequest)
		case auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrCreatingUser:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		default:
			log.LogHandlerError(logger, fmt.Errorf("неизвестная ошибка: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "пользователь с таким логином уже существует", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AdminJWT",
		Value:    user.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "CSRF-Token",
		Value:    user.CsrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("X-CSRF-Token", user.CsrfToken)

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
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

	user, err := h.client.Check(r.Context(), &gen.CheckRequest{
		Login: login,
	})

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		}
	}

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil || cookie.Value == "" {
		log.LogHandlerError(logger, fmt.Errorf("пользователь уже разлогинен: %w", err), http.StatusBadRequest)
		utils.SendError(w, "пользователь уже разлогинен", http.StatusBadRequest)
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

	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}

func (h *AuthHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		return
	}

	var updateData models.UpdateUserReq
	if err := easyjson.UnmarshalFromReader(r.Body, &updateData); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	updateData.Sanitize()

	user, err := h.client.UpdateUser(r.Context(), &gen.UpdateUserRequest{
		Login:       login,
		Description: updateData.Description,
		FirstName:   updateData.FirstName,
		LastName:    updateData.LastName,
		PhoneNumber: updateData.PhoneNumber,
		Password:    updateData.Password,
	})
	if err != nil {
		switch err {
		case auth.ErrInvalidPassword, auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrSamePassword:
			log.LogHandlerError(logger, err, http.StatusBadRequest)
			utils.SendError(w, err.Error(), http.StatusBadRequest)
		default:
			log.LogHandlerError(logger, fmt.Errorf("ошибка обновления данных пользователя: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "ошибка обновления данных пользователя", http.StatusInternalServerError)
		}
		return
	}

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) UpdateUserPic(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)

	if err := r.ParseMultipartForm(maxRequestBodySize); err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			log.LogHandlerError(logger, fmt.Errorf("превышен допустимый размер файла: %w", err), http.StatusRequestEntityTooLarge)
			utils.SendError(w, "превышен допустимый размер файла", http.StatusRequestEntityTooLarge)
		} else {
			log.LogHandlerError(logger, fmt.Errorf("невозможно обработать файл: %w", err), http.StatusBadRequest)
			utils.SendError(w, "невозможно обработать файл", http.StatusBadRequest)
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
		utils.SendError(w, "файл не найден в запросе", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении файла: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении файла", http.StatusBadRequest)
		return
	}
	file.Seek(0, io.SeekStart)

	mimeType := http.DetectContentType(buffer)
	if _, ok := allowedMimeTypes[mimeType]; !ok {
		log.LogHandlerError(logger, fmt.Errorf("недопустимый формат файла: %w", err), http.StatusBadRequest)
		utils.SendError(w, "недопустимый формат файла.", http.StatusBadRequest)
		return
	}

	ext := allowedMimeTypes[mimeType]

	picBytes, _ := io.ReadAll(file)

	user, err := h.client.UpdateUserPic(r.Context(), &gen.UpdateUserPicRequest{
		Login: login,
		UserPic: picBytes,
		FileExtension: ext,
	})
	if err != nil {
		switch err {
		case auth.ErrUserNotFound:
			log.LogHandlerError(logger, err, http.StatusNotFound)
			utils.SendError(w, err.Error(), http.StatusNotFound)
		case auth.ErrBasePath:
			log.LogHandlerError(logger, err, http.StatusInternalServerError)
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
		case auth.ErrFileCreation, auth.ErrFileSaving, auth.ErrFileDeletion:
			log.LogHandlerError(logger, fmt.Errorf("ошибка при работе с файлом: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "ошибка при работе с файлом", http.StatusInternalServerError)
		default:
			log.LogHandlerError(logger, fmt.Errorf("ошибка при обновлении аватарки: %w", err), http.StatusInternalServerError)
			utils.SendError(w, "ошибка при обновлении аватарки", http.StatusInternalServerError)
		}
		return
	}

	parsedUUID, err := uuid.FromString(user.Id)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("некорректный id: %w", err), http.StatusUnauthorized)
		utils.SendError(w, "некорректный id", http.StatusUnauthorized)
	}

	newModel := models.User{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          parsedUUID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}

	data, err := json.Marshal(newModel)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) GetUserAddresses(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		return
	}

	addresses, err := h.client.GetUserAddresses(r.Context(), &gen.AddressRequest{
		Login: login,
	})
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var modelAddresses []models.Address
	for _, addr := range addresses.Addresses {
		parsedUUIDa, err := uuid.FromString(addr.Id)
		if err != nil {
			log.LogHandlerError(logger, fmt.Errorf("некорректный id адреса: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "некорректный id адреса", http.StatusUnauthorized)
		}
		parsedUUIDu, err := uuid.FromString(addr.UserId)
		if err != nil {
			log.LogHandlerError(logger, fmt.Errorf("некорректный id пользователя: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "некорректный id пользователя", http.StatusUnauthorized)
		}
		modelAddresses = append(modelAddresses, models.Address{
			Id:      parsedUUIDa,
			Address: addr.Address,
			UserId:  parsedUUIDu,
		})
	}

	data, err := json.Marshal(modelAddresses)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка маршалинга: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "не удалось сериализовать данные", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func (h *AuthHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		return
	}

	var address models.Address
	err := easyjson.UnmarshalFromReader(r.Body, &address)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	address.Sanitize()

	_, err = h.client.DeleteAddress(r.Context(), &gen.DeleteAddressRequest{
		Id: address.Id.String(),
	})
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Info("Successful")
	w.Header().Set("Content-Type", "application/json")
}

func (h *AuthHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	idStr, ok := jwtUtils.GetIdFromJWT(JWTStr, claims, h.secret)
	if !ok || idStr == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: id отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "недействительный токен: id отсутствует", http.StatusUnauthorized)
		return
	}

	if !jwtUtils.CheckDoubleSubmitCookie(w, r) {
		utils.SendError(w, "некорректный CSRF-токен", http.StatusForbidden)
		log.LogHandlerError(logger, errors.New("некорректный CSRF-токен"), http.StatusForbidden)
		return
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var address models.Address
	err = easyjson.UnmarshalFromReader(r.Body, &address)
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}
	address.UserId = id
	address.Sanitize()

	_, err = h.client.AddAddress(r.Context(), &gen.Address{
		Id:      address.Id.String(),
		Address: address.Address,
		UserId:  address.UserId.String(),
	})
	if err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		utils.SendError(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	log.LogHandlerInfo(logger, "Successful", http.StatusOK)
}
