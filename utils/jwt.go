package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(userId uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte("Secret"))
	return token, err
}

type Claims struct {
	jwt.StandardClaims
}

func ParseJwt(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Secret"), nil
	})
	if err != nil || !token.Valid {
		//c.Status(fiber.StatusUnauthorized)
		return "", err

	}
	claims := token.Claims.(*Claims)

	return claims.Issuer, nil
}
