package http

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
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
			mockGrpc := mocks.NewMockAuthGrpc(ctrl)
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

func TestCheck(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		mockerUsecase  func(mockUsecase *mocks.MockAuthUsecase)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Missing CSRF Cookie",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "CSRF Mismatch",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", "blablabla")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Missing AdminJWT Cookie",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Invalid JWT Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "invalid-token"})
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "User Not Found",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, "unknown-user", secret, userId)
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
			},
			mockerUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().Check(gomock.Any(), "unknown-user").
					Return(models.User{}, auth.ErrUserNotFound)
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  auth.ErrUserNotFound.Error(),
		},
		{
			name: "Successful",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			mockerUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().Check(gomock.Any(), login).
					Return(models.User{Login: login}, nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			r := httptest.NewRequest("GET", "/api/auth/check", nil)
			w := httptest.NewRecorder()
			tt.cookieSetup(r)

			if tt.mockerUsecase != nil {
				tt.mockerUsecase(mockUsecase)
			}

			handler := AuthHandler{uc: mockUsecase, secret: secret}
			handler.Check(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
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

			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		requestBody    string
		mockerUsecase  func(mockUsecase *mocks.MockAuthUsecase)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Missing AdminJWT Cookie",
			cookieSetup: func(r *http.Request) {
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Токен отсутствует",
		},
		{
			name: "Invalid AdminJWT Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "blablabla"})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Недействительный токен: login отсутствует",
		},
		{
			name: "CSRF Token Mismatch",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", "blablabla")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Error Parsing JSON Body",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody:    `invalid json`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Ошибка парсинга JSON",
		},
		{
			name: "Invalid Update User Data (Password)",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: `{"password": "p"}`,
			mockerUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().UpdateUser(gomock.Any(), login, gomock.Any()).
					Return(models.User{}, auth.ErrInvalidPassword)
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  auth.ErrInvalidPassword.Error(),
		},
		{
			name: "Successful Update User",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: `{"password": "newPass@123", "first_name": "Доминик"}`,
			mockerUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().UpdateUser(gomock.Any(), login, gomock.Any()).
					Return(models.User{Login: login}, nil)
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			r := httptest.NewRequest("POST", "/api/auth/update_user", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()

			tt.cookieSetup(r)

			if tt.mockerUsecase != nil {
				tt.mockerUsecase(mockUsecase)
			}

			handler := AuthHandler{mockUsecase, secret}
			handler.UpdateUser(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

func TestGetUserAddresses(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		mockUsecase    func(mockUsecase *mocks.MockAuthUsecase)
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Missing AdminJWT Cookie",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Токен отсутствует",
		},
		{
			name: "Invalid AdminJWT Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "blablabla"})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Недействительный токен: login отсутствует",
		},
		{
			name: "CSRF Token Mismatch",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", "blablabla")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Usecase Error",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().GetUserAddresses(gomock.Any(), login).
					Return(nil, errors.New("usecase error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "usecase error",
		},
		{
			name: "Successful Response",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().GetUserAddresses(gomock.Any(), login).
					Return([]models.Address{
						{Id: uuid.NewV4(), Address: "г. Москва, Кремль", UserId: uuid.NewV4()},
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

			if tt.mockUsecase != nil {
				tt.mockUsecase(mockUsecase)
			}

			r := httptest.NewRequest("GET", "/api/auth/address", nil)
			w := httptest.NewRecorder()
			tt.cookieSetup(r)

			handler := AuthHandler{uc: mockUsecase, secret: secret}
			handler.GetUserAddresses(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

func TestDeleteAddress(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	addressId := uuid.NewV4()
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		mockUsecase    func(mockUsecase *mocks.MockAuthUsecase)
		requestBody    string
		expectedStatus int
		expectedError  string
	}{
		{
			name: "CSRF Token Mismatch",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", "blablabla")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Error Parsing JSON Body",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody:    `invalid json`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Ошибка парсинга JSON",
		},
		{
			name: "Usecase Error",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: fmt.Sprintf(`{"id":"%s"}`, addressId),
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().DeleteAddress(gomock.Any(), addressId).
					Return(errors.New("usecase error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "usecase error",
		},
		{
			name: "Successful Delete",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: fmt.Sprintf(`{"id":"%s"}`, addressId),
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().DeleteAddress(gomock.Any(), addressId).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			if tt.mockUsecase != nil {
				tt.mockUsecase(mockUsecase)
			}

			r := httptest.NewRequest("DELETE", "/api/auth/address", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()
			tt.cookieSetup(r)

			handler := AuthHandler{uc: mockUsecase, secret: secret}
			handler.DeleteAddress(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}

func TestAddAddress(t *testing.T) {
	secret := "secret-value"
	login := "testuser"
	csrf_token := "test-csrf"
	userId := uuid.NewV4()
	address := models.Address{
		Address: "г. Москва, Кремль",
		UserId:  userId,
	}
	tests := []struct {
		name           string
		cookieSetup    func(r *http.Request)
		mockUsecase    func(mockUsecase *mocks.MockAuthUsecase)
		requestBody    string
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Missing AdminJWT Cookie",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Токен отсутствует",
		},
		{
			name: "Invalid AdminJWT Token",
			cookieSetup: func(r *http.Request) {
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "blablabla"})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "Недействительный токен: id отсутствует",
		},
		{
			name: "CSRF Token Mismatch",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", "blablabla")
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "Error Parsing JSON Body",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody:    `invalid json`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Ошибка парсинга JSON",
		},
		{
			name: "Usecase Error",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: fmt.Sprintf(`{"address":"%s"}`, address.Address),
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), address).
					Return(errors.New("usecase error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "usecase error",
		},
		{
			name: "Successful Add",
			cookieSetup: func(r *http.Request) {
				tokenStr := utils.GenerateJWTForTest(t, login, secret, userId)
				r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: tokenStr})
				r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrf_token})
				r.Header.Set("X-CSRF-Token", csrf_token)
			},
			requestBody: fmt.Sprintf(`{"address":"%s"}`, address.Address),
			mockUsecase: func(mockUsecase *mocks.MockAuthUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), address).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUsecase := mocks.NewMockAuthUsecase(ctrl)
			defer ctrl.Finish()

			if tt.mockUsecase != nil {
				tt.mockUsecase(mockUsecase)
			}

			r := httptest.NewRequest("POST", "/api/auth/address", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()
			tt.cookieSetup(r)

			handler := AuthHandler{uc: mockUsecase, secret: secret}
			handler.AddAddress(w, r)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.expectedError != "" {
				assert.Contains(t, w.Body.String(), tt.expectedError)
			}
		})
	}
}
