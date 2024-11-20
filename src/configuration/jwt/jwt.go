package jwt

import (
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("minha_chave_secreta") // Altere para uma chave segura

func GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"id":  fmt.Sprint(userID),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
