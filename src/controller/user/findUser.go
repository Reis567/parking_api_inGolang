package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/google/uuid"
	"net/mail"
)


// FindUserByEmail busca um usuário pelo email
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	email := c.Param("email") // Obtenção do email a partir dos parâmetros da URL
	logger.Info("Iniciando FindUserByEmailController", zap.String("email", email))

	// Verifica se o email é válido usando mail.ParseAddress
	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Email do usuário inválido", zap.String("email", email), zap.Error(err))
		c.JSON(400, gin.H{
			"message": "Email do usuário inválido. Por favor, forneça um email válido.",
		})
		return
	}

	// Chama o serviço para buscar o usuário
	user, err := uc.service.FindUserByEmailService(email)
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


func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	userID := c.Param("id") // Obtenção do ID a partir dos parâmetros da URL
	logger.Info("Iniciando FindUserByIDController", zap.String("user_id", userID))

	// Verifica se o ID é um UUID válido usando a biblioteca uuid
	if _, err := uuid.Parse(userID); err != nil {
		logger.Error("ID do usuário inválido", zap.String("user_id", userID), zap.Error(err))
		c.JSON(400, gin.H{
			"message": "ID do usuário inválido. Deve ser um UUID válido.",
		})
		return
	}

	// Chama o serviço para buscar o usuário
	user, err := uc.service.FindUserByIDService(userID)
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