package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
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
			mockGrpc := mocks.NewMockAuthServiceClient(ctrl)
			defer ctrl.Finish()

			if tt.name != "Invalid JSON" {
				mockGrpc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
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

			handler := CreateAuthHandler(mockGrpc)
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
			mockUsecase := mocks.NewMockAuthServiceClient(ctrl)
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
