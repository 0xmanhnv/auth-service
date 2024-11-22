package jwt

import (
	"context"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(ctx context.Context, username string, secretKey string) (string, error) {
	claims := gojwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string, secretKey string) (*gojwt.Token, error) {
	claims := &gojwt.MapClaims{}
	token, err := gojwt.ParseWithClaims(tokenString, claims, func(token *gojwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	return token, err
}
