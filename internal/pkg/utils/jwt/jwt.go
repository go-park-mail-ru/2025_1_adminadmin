package utils

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/require"
)

func GetLoginFromJWT(JWTStr string, claims jwt.MapClaims, secret string) (string, bool) {
	token, err := jwt.ParseWithClaims(JWTStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", false
	}

	login, ok := claims["login"].(string)
	return login, ok
}

func CheckDoubleSubmitCookie(w http.ResponseWriter, r *http.Request) bool {
	cookieCSRF, err := r.Cookie("CSRF-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	headerCSRF := r.Header.Get("X-CSRF-Token")
	if cookieCSRF.Value == "" || headerCSRF == "" || cookieCSRF.Value != headerCSRF {
		w.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}

func GetIdFromJWT(JWTStr string, claims jwt.MapClaims, secret string) (string, bool) {
	token, err := jwt.ParseWithClaims(JWTStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if secret == "" {
			return nil, fmt.Errorf("JWT_SECRET не установлен")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", false
	}

	id, ok := claims["id"].(string)
	return id, ok
}

func GenerateJWTForTest(t *testing.T, login, secret string, id uuid.UUID) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"id":    id,
	})
	tokenStr, err := token.SignedString([]byte(secret))
	require.NoError(t, err)
	return tokenStr
}
