package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
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
