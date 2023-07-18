package system

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimsUsers struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func Verify(tokenString string) (*ClaimsUsers, error) {

	token, err := jwt.ParseWithClaims(tokenString, &ClaimsUsers{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(*ClaimsUsers); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func Sign(username string, email string) (string, error) {

	claims := ClaimsUsers{
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 1)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return ss, err
}
