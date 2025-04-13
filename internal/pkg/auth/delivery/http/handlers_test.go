package http

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	salt := make([]byte, 8)
	type args struct {
		login    string
		password string
	}

	var tests = []struct {
		name           string
		requestBody    string
		args           args
		ucErr          error
		expectedStatus int
	}{
		{
			name:        "Success",
			requestBody: `{"login":"test123","password":"Pass@123"}`,
			args: args{
				login:    "test123",
				password: "Pass@123",
			},
			ucErr:          nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"login":"testuser","password":"abc123"`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "User Not Found",
			requestBody: `{"login":"unknown","password":"somepass"}`,
			args: args{
				login:    "unknown",
				password: "somepass",
			},
			ucErr:          auth.ErrUserNotFound,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Invalid Credentials",
			requestBody: `{"login":"testuser","password":"wrongpass"}`,
			args: args{
				login:    "testuser",
				password: "wrongpass",
			},
			ucErr:          auth.ErrInvalidCredentials,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:        "Unknown Error",
			requestBody: `{"login":"testuser","password":"errorpass"}`,
			args: args{
				login:    "testuser",
				password: "errorpass",
			},
			ucErr:          errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			if tt.name != "Invalid JSON" {
				mockUsecase.EXPECT().SignIn(gomock.Any(), models.SignInReq{
					Login:    tt.args.login,
					Password: tt.args.password,
				}).Return(models.User{
					Login:        tt.args.login,
					FirstName:    "Иван",
					LastName:     "Иванов",
					PasswordHash: usecase.HashPassword(salt, tt.args.password),
					Id:           uuid.NewV4(),
					PhoneNumber:  "88005553535",
					Description:  "",
					UserPic:      "default.png",
				}, "jwt_token", "csrf_token", tt.ucErr)
			}

			r := httptest.NewRequest("POST", "/api/auth/signin", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()

			handler := CreateAuthHandler(mockUsecase)
			handler.SignIn(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestSignUp(t *testing.T) {
	salt := make([]byte, 8)
	type args struct {
		login        string
		password     string
		first_name   string
		last_name    string
		phone_number string
	}

	var tests = []struct {
		name           string
		requestBody    string
		args           args
		ucErr          error
		expectedStatus int
	}{
		{
			name:        "Success",
			requestBody: `{"login":"test123","password":"Pass@123","first_name":"Иван","last_name":"Иванов","phone_number":"88005553535"}`,
			args: args{
				login:        "test123",
				password:     "Pass@123",
				first_name:   "Иван",
				last_name:    "Иванов",
				phone_number: "88005553535",
			},
			ucErr:          nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid JSON",
			requestBody:    `{"login":"testuser","password":"abc123"`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Invalid Login or Password",
			requestBody: `{"login":"testuser","password":"wrongpass","first_name":"Иван","last_name":"Иванов","phone_number":"88005553535"}`,
			args: args{
				login:        "testuser",
				password:     "wrongpass",
				first_name:   "Иван",
				last_name:    "Иванов",
				phone_number: "88005553535",
			},
			ucErr:          auth.ErrInvalidLogin,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Invalid Name or Phone",
			requestBody: `{"login":"testuser","password":"validPass123","first_name":"!@#","last_name":"Иванов","phone_number":"88005553535"}`,
			args: args{
				login:        "testuser",
				password:     "validPass123",
				first_name:   "!@#",
				last_name:    "Иванов",
				phone_number: "88005553535",
			},
			ucErr:          auth.ErrInvalidName,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "User Creation Failed",
			requestBody: `{"login":"newuser","password":"validPass123","first_name":"Иван","last_name":"Иванов","phone_number":"88005553535"}`,
			args: args{
				login:        "newuser",
				password:     "validPass123",
				first_name:   "Иван",
				last_name:    "Иванов",
				phone_number: "88005553535",
			},
			ucErr:          auth.ErrCreatingUser,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Unknown Error",
			requestBody: `{"login":"testuser","password":"errorpass","first_name":"Иван","last_name":"Иванов","phone_number":"88005553535"}`,
			args: args{
				login:        "testuser",
				password:     "errorpass",
				first_name:   "Иван",
				last_name:    "Иванов",
				phone_number: "88005553535",
			},
			ucErr:          errors.New("unknown error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			if tt.name != "Invalid JSON" {
				mockUsecase.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
					Login:       tt.args.login,
					Password:    tt.args.password,
					FirstName:   tt.args.first_name,
					LastName:    tt.args.last_name,
					PhoneNumber: tt.args.phone_number,
				}).Return(models.User{
					Login:        tt.args.login,
					FirstName:    tt.args.first_name,
					LastName:     tt.args.last_name,
					PasswordHash: usecase.HashPassword(salt, tt.args.password),
					Id:           uuid.NewV4(),
					PhoneNumber:  tt.args.phone_number,
					Description:  "",
					UserPic:      "default.png",
				}, "jwt_token", "csrf_token", tt.ucErr)
			}

			r := httptest.NewRequest("POST", "/api/auth/signup", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()

			handler := CreateAuthHandler(mockUsecase)
			handler.SignUp(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestLogOut(t *testing.T) {
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "User already logged out (no AdminJWT cookie)",
			cookieSetup: func(r *http.Request) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Пользователь уже разлогинен",
		},
		{
			name: "Successful_logout",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{
					Name:  "AdminJWT",
					Value: "some-valid-jwt-token",
				})
			},
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			r := httptest.NewRequest("POST", "/api/auth/logout", nil)
			w := httptest.NewRecorder()
			tt.cookieSetup(r)

			handler := CreateAuthHandler(mockUsecase)
			handler.LogOut(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)

			body := w.Body.String()
			if tt.expectedError != "" && !strings.Contains(body, tt.expectedError) {
				t.Errorf("expected error message to contain %s, got %s", tt.expectedError, body)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	secret := "some secret"
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		requestBody    string
		mockUsecase    func(mockUsecase *mocks.MockAuthUsecase)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Missing AdminJWT Cookie",
			cookieSetup: func(r *http.Request) {
				// Не добавляем AdminJWT cookie, чтобы симулировать ошибку
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Токен отсутствует",
		},
		{
			name: "Invalid AdminJWT Token",
			cookieSetup: func(r *http.Request) {
				// Добавляем невалидный токен
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "invalid-token"})
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Недействительный токен: login отсутствует",
		},
		{
			name: "Error Parsing JSON Body",
			cookieSetup: func(r *http.Request) {
				// Добавляем валидный токен
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"login": "test123",
					"id":    uuid.NewV4(),
					"exp":   time.Now().Add(24 * time.Hour).Unix(),
				})
				tokenStr, _ := token.SignedString([]byte(secret))
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
			},
			requestBody:    `{"password": "newpassword"}`, // Невалидный JSON
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Ошибка парсинга JSON",
		},
		{
			name: "Invalid Update User Data (Password)",
			cookieSetup: func(r *http.Request) {
				// Добавляем валидный токен
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "valid-token"})
			},
			requestBody: `{"password": "newpassword"}`,
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				// Симулируем ошибку неверного пароля
				mockUsecase.EXPECT().UpdateUser(gomock.Any(), "valid-login", gomock.Any()).Return(models.User{}, auth.ErrInvalidPassword)
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Неверный пароль",
		},
		{
			name: "Successful Update User",
			cookieSetup: func(r *http.Request) {
				// Добавляем валидный токен
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "valid-token"})
			},
			requestBody: `{"password": "newpassword", "name": "newname"}`,
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				// Симулируем успешное обновление пользователя
				mockUsecase.EXPECT().UpdateUser(gomock.Any(), "valid-login", gomock.Any()).Return(models.User{
					Login: "valid-login",
				}, nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			r := httptest.NewRequest("POST", "/api/auth/update-user", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()

			// Устанавливаем куки, если необходимо
			tt.cookieSetup(r)

			// Логирование куки перед обработчиком
			log.Println("Cookies in request:", r.Cookies())

			// Настроим моки для каждого теста
			if tt.mockUsecase != nil {
				tt.mockUsecase(mockUsecase)
			}

			handler := CreateAuthHandler(mockUsecase)

			handler.UpdateUser(w, r)

			// Проверка статуса
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Проверка сообщения об ошибке
			if tt.expectedError != "" {
				body := w.Body.String()
				if !strings.Contains(body, tt.expectedError) {
					t.Errorf("expected error message to contain %s, got %s", tt.expectedError, body)
				}
			}
		})
	}
}
