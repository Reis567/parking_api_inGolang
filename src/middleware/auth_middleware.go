package middleware

import (
	"meu-novo-projeto/src/configuration/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica a validade do token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o token do cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token não fornecido"})
			c.Abort()
			return
		}

		// Remove o prefixo "Bearer " do token, se existir
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Valida o token
		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token inválido ou expirado", "error": err.Error()})
			c.Abort()
			return
		}

		// Define os dados do usuário no contexto para uso posterior
		c.Set("userID", claims["id"])
		c.Set("email", claims["sub"])

		// Permite a continuação da requisição
		c.Next()
	}
}
