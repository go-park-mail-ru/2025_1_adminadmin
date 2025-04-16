package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

const secret = "test_secret"

func createTestJWT(t *testing.T, claims jwt.MapClaims, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	assert.NoError(t, err)
	return tokenStr
}

func TestGetLoginFromJWT(t *testing.T) {
	claims := jwt.MapClaims{
		"login": "testuser",
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	tokenStr := createTestJWT(t, claims, secret)

	login, ok := GetLoginFromJWT(tokenStr, jwt.MapClaims{}, secret)
	assert.True(t, ok)
	assert.Equal(t, "testuser", login)
}

func TestGetIdFromJWT(t *testing.T) {
	claims := jwt.MapClaims{
		"id":  "12345",
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	tokenStr := createTestJWT(t, claims, secret)

	id, ok := GetIdFromJWT(tokenStr, jwt.MapClaims{}, secret)
	assert.True(t, ok)
	assert.Equal(t, "12345", id)
}

func TestGenerateJWTForTest(t *testing.T) {
	login := "testuser"
	secret := "secret"
	id := uuid.NewV4()

	tokenStr := GenerateJWTForTest(t, login, secret, id)
	assert.NotNil(t, tokenStr)
}

func TestCheckDoubleSubmitCookie(t *testing.T) {
	tests := []struct {
		name            string
		cookieValue     string
		headerValue     string
		expectStatus    int
		expectReturnVal bool
		skipCookie      bool
	}{
		{
			name:            "Valid CSRF Token",
			cookieValue:     "valid-token",
			headerValue:     "valid-token",
			expectStatus:    0,
			expectReturnVal: true,
		},
		{
			name:            "Missing Cookie",
			headerValue:     "some-token",
			expectStatus:    http.StatusUnauthorized,
			expectReturnVal: false,
			skipCookie:      true,
		},
		{
			name:            "Missing Header",
			cookieValue:     "token",
			headerValue:     "",
			expectStatus:    http.StatusForbidden,
			expectReturnVal: false,
		},
		{
			name:            "Mismatched Tokens",
			cookieValue:     "cookie-token",
			headerValue:     "header-token",
			expectStatus:    http.StatusForbidden,
			expectReturnVal: false,
		},
		{
			name:            "Empty Tokens",
			cookieValue:     "",
			headerValue:     "",
			expectStatus:    http.StatusForbidden,
			expectReturnVal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if !tt.skipCookie {
				req.AddCookie(&http.Cookie{
					Name:  "CSRF-Token",
					Value: tt.cookieValue,
				})
			}
			req.Header.Set("X-CSRF-Token", tt.headerValue)

			rr := httptest.NewRecorder()

			result := CheckDoubleSubmitCookie(rr, req)

			if rr.Code != 200 && rr.Code != tt.expectStatus {
				t.Errorf("unexpected status code: got %d, want %d", rr.Code, tt.expectStatus)
			}
			if result != tt.expectReturnVal {
				t.Errorf("unexpected result: got %v, want %v", result, tt.expectReturnVal)
			}
		})
	}
}
