package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Chave secreta usada para assinar os tokens
var jwtKey = []byte("minha_chave_secreta")

// GenerateToken gera um JWT com base no ID e no email do usuário
func GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"id":  fmt.Sprint(userID),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken verifica e valida o JWT recebido
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Analisa o token e valida a assinatura
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é o esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao analisar token: %w", err)
	}

	// Verifica se o token é válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido ou expirado")
}
