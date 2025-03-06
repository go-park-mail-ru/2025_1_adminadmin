package handlers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/satori/uuid"
// )



// func TestSignIn(t *testing.T) {
// 	salt := make([]byte, 8)
// 	type args struct {
// 		w *httptest.ResponseRecorder
// 		r *http.Request
// 	}

// 	testUser := models.User{
// 		Login:        "test123",
// 		PasswordHash: hashPassword(salt, "password123"),
// 		Id:           uuid.NewV4(),
// 		PhoneNumber:  "88005553535",
// 		Description:  "New User",
// 		UserPic:      "default.png",
// 	}
// 	users["test123"] = testUser

// 	tests := []struct {
// 		name         string
// 		args         args
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name: "OK sign in",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"test123","password":"password123"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusOK,
// 			expectedBody: fmt.Sprintf(`{"login":"test123","phone_number":"88005553535","id":"%s","description":"New User","path":"default.png"}`, testUser.Id.String()),
// 		},
// 		{
// 			name: "Invalid login",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"t","password":"password123"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `Неверный формат логина`,
// 		},
// 		{
// 			name: "Invalid credentials",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", bytes.NewBuffer([]byte(`{"login":"test123","password":"password"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 			expectedBody: `Неверные данные`,
// 		},
// 		{
// 			name: "Empty request body",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signin", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: "",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			SignIn(test.args.w, test.args.r)

// 			if test.args.w.Code != test.expectedCode {
// 				t.Errorf("Unexpected response code: expected %d got %d", test.expectedCode, test.args.w.Code)
// 			}

// 			if test.expectedBody != "" {
// 				responseBody := strings.TrimSpace(test.args.w.Body.String())
// 				expectedBody := strings.TrimSpace(test.expectedBody)

// 				if test.expectedCode == http.StatusOK {
// 					if responseBody != expectedBody {
// 						t.Errorf("Unexpected response body: expected %s got %s", test.expectedBody, responseBody)
// 					}
// 				}
// 			}

// 			if test.expectedCode == http.StatusOK {
// 				contentType := test.args.w.Header().Get("Content-Type")
// 				if contentType != "application/json" {
// 					t.Errorf("Unexpected 'Content-Type' header value: expected 'application/json' got %s", contentType)
// 				}

// 				cookies := test.args.w.Result().Cookies()
// 				if len(cookies) == 0 {
// 					t.Error("Cookie was expected but it is absent")
// 					return
// 				}

// 				cookie := cookies[0]
// 				if cookie.Name != "AdminJWT" {
// 					t.Errorf("Unexpected cookie name: expected 'AdminJWT' got %s", cookie.Name)
// 				}
// 			}
// 		})
// 	}
// }

// func TestSignUp(t *testing.T) {
// 	type args struct {
// 		w *httptest.ResponseRecorder
// 		r *http.Request
// 	}

// 	testUser := models.User{
// 		Login:        "existing_user",
// 		PasswordHash: []byte(hashSHA256("Pass@123")),
// 		Id:           uuid.NewV4(),
// 		PhoneNumber:  "88005553535",
// 		Description:  "Existing User",
// 		UserPic:      "default.png",
// 	}
// 	users["existing_user"] = testUser

// 	tests := []struct {
// 		name         string
// 		args         args
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name: "OK sign up",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"new_user","first_name":"Тайлер","second_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusCreated,
// 			expectedBody: `{"login":"new_user","phone_number":"88005553535","description":"New User","path":"default.png"}`,
// 		},
// 		{
// 			name: "Invalid login format",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"u","first_name":"Тайлер","second_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `Неверный формат логина`,
// 		},
// 		{
// 			name: "Invalid password format",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"new_user","first_name":"Тайлер","second_name":"Дерден","phone_number":"88005553535","password":"pass"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `Неверный формат пароля`,
// 		},
// 		{
// 			name: "Duplicate user",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", bytes.NewBuffer([]byte(`{"login":"existing_user","first_name":"Тайлер","second_name":"Дерден","phone_number":"88005553535","password":"Pass@123"}`))),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusConflict,
// 			expectedBody: `Данный логин уже занят`,
// 		},
// 		{
// 			name: "Empty request body",
// 			args: args{
// 				r: httptest.NewRequest("POST", "http://localhost:5458/api/auth/signup", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: "",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			SignUp(test.args.w, test.args.r)

// 			if test.args.w.Code != test.expectedCode {
// 				t.Errorf("Unexpected response code: expected %d got %d", test.expectedCode, test.args.w.Code)
// 			}

// 			if test.expectedBody != "" {
// 				responseBody := strings.TrimSpace(test.args.w.Body.String())

// 				if test.expectedCode == http.StatusCreated {
// 					var responseUser models.User
// 					err := json.Unmarshal([]byte(responseBody), &responseUser)
// 					if err != nil {
// 						t.Fatalf("Failed to unmarshal literal response: %v", err)
// 					}

// 					var expectedUser models.User
// 					err = json.Unmarshal([]byte(test.expectedBody), &expectedUser)
// 					if err != nil {
// 						t.Fatalf("failed to unmarshal expected response: %v", err)
// 					}

// 					if responseUser.Login != expectedUser.Login ||
// 						responseUser.PhoneNumber != expectedUser.PhoneNumber ||
// 						responseUser.Description != expectedUser.Description ||
// 						responseUser.UserPic != expectedUser.UserPic {
// 						t.Errorf("Unexpected request body: expected %+v got %+v", expectedUser, responseUser)
// 					}

// 					if responseUser.Id == uuid.Nil {
// 						t.Error("Expected not empty Id, got empty Id")
// 					}
// 				} else {
// 					if responseBody != test.expectedBody {
// 						t.Errorf("Unexpected response body: expected %s got %s", test.expectedBody, responseBody)
// 					}
// 				}
// 			}

// 			if test.expectedCode == http.StatusCreated {
// 				contentType := test.args.w.Header().Get("Content-Type")
// 				if contentType != "application/json" {
// 					t.Errorf("Unexpected 'Content-Type' header value: expected 'application/json' got %s", contentType)
// 				}

// 				cookies := test.args.w.Result().Cookies()
// 				if len(cookies) == 0 {
// 					t.Error("Cookie was expected but it is absent")
// 					return
// 				}

// 				cookie := cookies[0]
// 				if cookie.Name != "AdminJWT" {
// 					t.Errorf("Unexpected cookie name: expected 'AdminJWT' got %s", cookie.Name)
// 				}
// 			}
// 		})
// 	}
// }

// func TestCheck(t *testing.T) {
// 	type args struct {
// 		w *httptest.ResponseRecorder
// 		r *http.Request
// 	}

// 	os.Setenv("JWT_SECRET", "test_secret")
// 	defer os.Unsetenv("JWT_SECRET")

// 	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"login": "test123",
// 		"exp":   time.Now().Add(time.Hour).Unix(),
// 	})
// 	validTokenStr, _ := validToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

// 	invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"login": "test123",
// 		"exp":   time.Now().Add(time.Hour).Unix(),
// 	})
// 	invalidTokenStr, _ := invalidToken.SignedString([]byte("wrong_secret"))

// 	wrongMethodToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
// 		"login": "test123",
// 		"exp":   time.Now().Add(time.Hour).Unix(),
// 	})
// 	wrongMethodTokenStr, _ := wrongMethodToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

// 	csrfToken := "some_csrf_token"

// 	tests := []struct {
// 		name         string
// 		args         args
// 		token        string
// 		expectedCode int
// 	}{
// 		{
// 			name: "Valid token",
// 			args: args{
// 				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			token:        validTokenStr,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name: "No token",
// 			args: args{
// 				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			token:        "",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name: "Invalid token (wrong signature)",
// 			args: args{
// 				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			token:        invalidTokenStr,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name: "Invalid token (wrong signing method)",
// 			args: args{
// 				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			token:        wrongMethodTokenStr,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name: "Malformed token",
// 			args: args{
// 				r: httptest.NewRequest("GET", "http://localhost:5458/api/auth/check", nil),
// 				w: httptest.NewRecorder(),
// 			},
// 			token:        "just a random string",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			if test.token != "" {
// 				test.args.r.AddCookie(&http.Cookie{Name: "AdminJWT", Value: test.token})
// 			}

// 			test.args.r.Header.Set("X-CSRF-Token", csrfToken)
// 			test.args.r.AddCookie(&http.Cookie{Name: "CSRF-Token", Value: csrfToken})

// 			Check(test.args.w, test.args.r)

// 			if test.args.w.Code != test.expectedCode {
// 				t.Errorf("Unexpected response code: expected %d got: %d", test.expectedCode, test.args.w.Code)
// 			}
// 		})
// 	}
// }

// func TestValidLogin(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		login string
// 		want  bool
// 	}{
// 		{"Valid login", "user_123", true},
// 		{"Too short", "ab", false},
// 		{"Too long", "abcdefghijklmnopqrstuvwxyzABCDEF", false},
// 		{"Invalid characters", "user@name", false},
// 		{"Empty login", "", false},
// 		{"Valid with hyphen", "user-name", true},
// 		{"Valid with underscore", "user_name", true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := validLogin(tt.login); got != tt.want {
// 				t.Errorf("validLogin(%q) = %v, want %v", tt.login, got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestValidPassword(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		password string
// 		want     bool
// 	}{
// 		{"Valid password", "Password1!", true},
// 		{"Too short", "Pass1!", false},
// 		{"Too long", "ThisPasswordIsWayTooLong123!", false},
// 		{"No uppercase", "password1!", false},
// 		{"No lowercase", "PASSWORD1!", false},
// 		{"No digit", "Password!", false},
// 		{"No special character", "Password1", false},
// 		{"Empty password", "", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := validPassword(tt.password); got != tt.want {
// 				t.Errorf("validPassword(%q) = %v, want %v", tt.password, got, tt.want)
// 			}
// 		})
// 	}
// }