package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AdminCliams struct {
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

const (
	AdminUserName = "user"
	AdminPassword = "123"
)

func NewJWTForAdmin() string {

	claims := &AdminCliams{
		IsAdmin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}
	return tokenString

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
