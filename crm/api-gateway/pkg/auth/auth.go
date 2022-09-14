package auth

import (
	"api-gateway/pkg/httperr"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const jwtSecret = "very secret key"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			httperr.Unauthorized(w, r, "no token in Authorization header")
			return
		}

		_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(jwtSecret), nil
		})
		if err != nil {
			httperr.Unauthorized(w, r, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewJWT(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}

	return tokenString
}
