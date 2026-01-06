package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("poiuytrewq")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}