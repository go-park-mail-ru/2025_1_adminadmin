package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/satori/uuid"
)

var users = make(map[string]models.User)

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
		PasswordHash: usecase.HashPassword(salt, "Pass@123"),
		Id:           uuid.NewV4(),
		PhoneNumber:  "88005553535",
		Description:  "",
		UserPic:      "default.png",
	}
	users["test123"] = testUser

	tests := []struct {
		name           string
		requestBody    string
		login          string
		password       string
		usecaseErr     error
		expectedStatus int
	}{
		{
			name:           "Successful sign in",
			requestBody:    `{"login":"test123","password":"Pass@123"}`,
			login:          "test123",
			password:       "Pass@123",
			usecaseErr:     nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Fail",
			requestBody:    `{"login":"testuser2","password":"12345678a"`,
			login:       "testuser2",
			password:       "12345678a",
			usecaseErr:     nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

}
