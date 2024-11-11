package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FindUserByEmail busca um usuário pelo email
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	email := c.Param("email") // Obtenção do email a partir dos parâmetros da URL
	logger.Info("Iniciando FindUserByEmailController", zap.String("email", email))

	// Chama o serviço para buscar o usuário
	user, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Erro ao buscar usuário", zap.Error(err))
		c.JSON(err.Code, err)
		return
	}

	// Retorna a resposta com os dados do usuário
	c.JSON(200, gin.H{
		"id":         user.GetID(),
		"first_name": user.GetFirstName(),
		"last_name":  user.GetLastName(),
		"email":      user.GetEmail(),
		"age":        user.GetAge(),
		"created_at": user.GetCreatedAt(),
		"updated_at": user.GetUpdatedAt(),
	})
}


// FindUserByID busca um usuário pelo ID
func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	userID := c.Param("id") // Obtenção do ID a partir dos parâmetros da URL
	logger.Info("Iniciando FindUserByIDController", zap.String("user_id", userID))

	// Chama o serviço para buscar o usuário
	user, err := uc.service.FindUserByID(userID)
	if err != nil {
		logger.Error("Erro ao buscar usuário", zap.Error(err))
		c.JSON(err.Code, err)
		return
	}

	// Retorna a resposta com os dados do usuário
	c.JSON(200, gin.H{
		"id":         user.GetID(),
		"first_name": user.GetFirstName(),
		"last_name":  user.GetLastName(),
		"email":      user.GetEmail(),
		"age":        user.GetAge(),
		"created_at": user.GetCreatedAt(),
		"updated_at": user.GetUpdatedAt(),
	})
}