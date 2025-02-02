package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWT_SECRET_KEY = []byte("your-secret-key")

func GenerateJWT(adminID string) (string, error) {
	claims := jwt.MapClaims{
		"adminID": adminID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET_KEY)
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWT_SECRET_KEY, nil
	})

	return token, err
}
