package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
)

func TestSignIn(t *testing.T) {
	salt := make([]byte, 8)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	testUser := models.User{
		Login:        "test123",
		FirstName:    "Иван",
		LastName:     "Иванов",
		PasswordHash: hashPassword(salt, "Pass@123"),
		Id:           uuid.NewV4(),
		PhoneNumber:  "88005553535",
		Description:  "",
		UserPic:      "default.png",
	}
	users["test123"] = testUser

	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedUser models.User
		mockRows     *pgxpoolmock.Rows
	}{
		{
			name: "OK sign in",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"test123","password":"Pass@123"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusOK,
			expectedUser: testUser,
			mockRows: pgxpoolmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "description", "user_pic", "password_hash"}).
				AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.PhoneNumber, testUser.Description, testUser.UserPic, testUser.PasswordHash),
		},
		{
			name: "Invalid login",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"t","password":"Pass@123"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedUser: models.User{},
			mockRows:     nil,
		},
		{
			name: "Invalid credentials",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"test123","password":"wrong_password"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusUnauthorized,
			expectedUser: models.User{},
			mockRows: pgxpoolmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "description", "user_pic", "password_hash"}).
				AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.PhoneNumber, testUser.Description, testUser.UserPic, testUser.PasswordHash),
		},
		{
			name: "Empty request body",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedUser: models.User{},
			mockRows:     nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

			if test.mockRows != nil {
				mockPool.EXPECT().
					Query(gomock.Any(), selectUser, testUser.Login).
					Return(test.mockRows.ToPgxRows(), nil)
			}

			h := &Handler{db: mockPool}
			h.SignIn(test.args.w, test.args.r)

			if test.args.w.Code != test.expectedCode {
				t.Errorf("Unexpected response code: expected %d got %d", test.expectedCode, test.args.w.Code)
			}

			if test.expectedCode == http.StatusOK {
				var responseUser models.User
				err := json.NewDecoder(test.args.w.Body).Decode(&responseUser)
				if err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}

				if responseUser.Id != test.expectedUser.Id ||
					responseUser.Login != test.expectedUser.Login ||
					responseUser.PhoneNumber != test.expectedUser.PhoneNumber ||
					responseUser.Description != test.expectedUser.Description ||
					responseUser.UserPic != test.expectedUser.UserPic {
					t.Errorf("Unexpected response body: expected %+v got %+v", test.expectedUser, responseUser)
				}

				contentType := test.args.w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Unexpected 'Content-Type' header value: expected 'application/json' got %s", contentType)
				}

				cookies := test.args.w.Result().Cookies()
				if len(cookies) == 0 {
					t.Error("Cookie was expected but it is absent")
					return
				}

				cookie := cookies[0]
				if cookie.Name != "AdminJWT" {
					t.Errorf("Unexpected cookie name: expected 'AdminJWT' got %s", cookie.Name)
				}
			}
		})
	}
}

func TestSignUp(t *testing.T) {
	salt := make([]byte, 8)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	testUser := models.User{
		Login:        "existing_user",
		PasswordHash: hashPassword(salt, "password123"),
		Id:           uuid.NewV4(),
		FirstName:    "Тайлер",
		LastName:     "Дерден",
		PhoneNumber:  "88005553535",
		Description:  "Existing User",
		UserPic:      "default.png",
	}
	users["existing_user"] = testUser

	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedBody string
		mockExec     error
	}{
		{
			name: "OK sign up",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"new_user","first_name":"Тайлер","last_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"login":"new_user","first_name":"Тайлер","last_name":"Дерден","phone_number":"88005553535","description":"","user_pic":"default.png"}`,
			mockExec:     nil,
		},
		{
			name: "Invalid login format",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"u","first_name":"Тайлер","last_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `Неверный формат логина`,
			mockExec:     nil,
		},
		{
			name: "Invalid password format",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"new_user","first_name":"Тайлер","last_name":"Дерден","phone_number":"88005553535","password":"pass"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `Неверный формат пароля`,
			mockExec:     nil,
		},
		{
			name: "Duplicate user",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"existing_user","first_name":"Тайлер","last_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusConflict,
			expectedBody: `Данный логин уже занят`,
			mockExec:     fmt.Errorf("duplicate key"),
		},
		{
			name: "Empty request body",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
			mockExec:     nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)

			if test.mockExec == nil && test.expectedCode == http.StatusCreated {
				mockPool.EXPECT().
					Exec(gomock.Any(), insertUser,
						gomock.Any(),
						"new_user",
						"Тайлер",
						"Дерден",
						"88005553535",
						"",
						"default.png",
						gomock.Any(),
					).
					Return(nil, nil)
			} else if test.mockExec != nil {
				mockPool.EXPECT().
					Exec(gomock.Any(), insertUser,
						gomock.Any(),
						"existing_user",
						"Тайлер",
						"Дерден",
						"88005553535",
						"",
						"default.png",
						gomock.Any(),
					).
					Return(nil, test.mockExec)
			}

			h := &Handler{db: mockPool}
			h.SignUp(test.args.w, test.args.r)

			if test.args.w.Code != test.expectedCode {
				t.Errorf("Unexpected response code: expected %d got %d", test.expectedCode, test.args.w.Code)
			}

			if test.expectedBody != "" {
				responseBody := strings.TrimSpace(test.args.w.Body.String())

				if test.expectedCode == http.StatusCreated {
					var responseUser models.User
					err := json.Unmarshal([]byte(responseBody), &responseUser)
					if err != nil {
						t.Fatalf("Failed to unmarshal literal response: %v", err)
					}

					var expectedUser models.User
					err = json.Unmarshal([]byte(test.expectedBody), &expectedUser)
					if err != nil {
						t.Fatalf("failed to unmarshal expected response: %v", err)
					}

					if responseUser.Login != expectedUser.Login ||
						responseUser.FirstName != expectedUser.FirstName ||
						responseUser.LastName != expectedUser.LastName ||
						responseUser.PhoneNumber != expectedUser.PhoneNumber ||
						responseUser.Description != expectedUser.Description {
						t.Errorf("Unexpected request body: expected %+v got %+v", expectedUser, responseUser)
					}
				} else {
					if responseBody != test.expectedBody {
						t.Errorf("Unexpected response body: expected %s got %s", test.expectedBody, responseBody)
					}
				}
			}

			if test.expectedCode == http.StatusCreated {
				contentType := test.args.w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Unexpected 'Content-Type' header value: expected 'application/json' got %s", contentType)
				}

				cookies := test.args.w.Result().Cookies()
				if len(cookies) == 0 {
					t.Error("Cookie was expected but it is absent")
					return
				}

				cookie := cookies[0]
				if cookie.Name != "AdminJWT" {
					t.Errorf("Unexpected cookie name: expected 'AdminJWT' got %s", cookie.Name)
				}
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": "test123",
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	validTokenStr, _ := validToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	csrfToken := "some_csrf_token"

	tests := []struct {
		name         string
		args         args
		token        string
		csrfToken    string
		expectedCode int
	}{
		{
			name: "Valid token and CSRF",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
				w: httptest.NewRecorder(),
			},
			token:        validTokenStr,
			csrfToken:    csrfToken,
			expectedCode: http.StatusOK,
		},
		{
			name: "No token",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
				w: httptest.NewRecorder(),
			},
			token:        "",
			csrfToken:    csrfToken,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Missing CSRF token",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
				w: httptest.NewRecorder(),
			},
			token:        validTokenStr,
			csrfToken:    "",
			expectedCode: http.StatusForbidden,
		},
		{
			name: "Invalid CSRF token",
			args: args{
				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
				w: httptest.NewRecorder(),
			},
			token:        validTokenStr,
			csrfToken:    "wrong_csrf_token",
			expectedCode: http.StatusForbidden,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	h := &Handler{db: mockPool}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.token != "" {
				test.args.r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: test.token})
			}

			if test.csrfToken != "" {
				test.args.r.Header.Set("X-CSRF-Token", test.csrfToken)
				test.args.r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})
			}

			h.Check(test.args.w, test.args.r)

			if test.args.w.Code != test.expectedCode {
				t.Errorf("Unexpected response code: expected %d got: %d", test.expectedCode, test.args.w.Code)
			}
		})
	}
}

func TestLogOut(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	tests := []struct {
		name         string
		args         args
		expectedCode int
	}{
		{
			name: "Successful logout",
			args: args{
				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/logout", nil),
				w: httptest.NewRecorder(),
			},
			expectedCode: http.StatusOK,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.args.r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: "some_token"})
			test.args.r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: "some_csrf_token"})

			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			h := &Handler{db: mockPool}
			h.LogOut(test.args.w, test.args.r)

			if test.args.w.Code != test.expectedCode {
				t.Errorf("Unexpected response code: expected %d got %d", test.expectedCode, test.args.w.Code)
			}

			cookies := test.args.w.Result().Cookies()
			if len(cookies) != 2 {
				t.Errorf("Expected 2 cookies to be set, got %d", len(cookies))
			}

			for _, cookie := range cookies {
				if cookie.Name != "AdminJWT" && cookie.Name != "CSRF-Token" {
					t.Errorf("Unexpected cookie name: %s", cookie.Name)
				}
				if cookie.Value != "" {
					t.Errorf("Expected cookie value to be empty, got %s", cookie.Value)
				}
				if cookie.Expires.After(time.Now()) {
					t.Errorf("Expected cookie to be expired, got expiration: %v", cookie.Expires)
				}
			}
		})
	}
}

func TestValidLogin(t *testing.T) {
	tests := []struct {
		name  string
		login string
		want  bool
	}{
		{"Valid login", "user_123", true},
		{"Too short", "ab", false},
		{"Too long", "abcdefghijklmnopqrstuvwxyzABCDEF", false},
		{"Invalid characters", "user@name", false},
		{"Empty login", "", false},
		{"Valid with hyphen", "user-name", true},
		{"Valid with underscore", "user_name", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validLogin(tt.login); got != tt.want {
				t.Errorf("validLogin(%q) = %v, want %v", tt.login, got, tt.want)
			}
		})
	}
}

func TestValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Valid password", "Password1!", true},
		{"Too short", "Pass1!", false},
		{"Too long", "ThisPasswordIsWayTooLong123!", false},
		{"No uppercase", "password1!", false},
		{"No lowercase", "PASSWORD1!", false},
		{"No digit", "Password!", false},
		{"No special character", "Password1", false},
		{"Empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPassword(tt.password); got != tt.want {
				t.Errorf("validPassword(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}
