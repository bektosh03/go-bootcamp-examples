package auth

import (
	"api-gateway/pkg/httperr"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
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

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Admin")
		if tokenString == "" {
			httperr.Unauthorized(w, r, "no token in Admin header")
			return
		}
		claims := &AdminCliams{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			httperr.Unauthorized(w, r, err.Error())
			return
		}
		if !tkn.Valid {
			httperr.Unauthorized(w, r, "token has been expired")
			return
		}

		r.Header.Set("Admin", tokenString)

		next.ServeHTTP(w, r)
	})
}
