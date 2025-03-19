package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv" 
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

	// Converte o ID para uint
	idUint, parseErr := strconv.ParseUint(userID, 10, 32)
	if parseErr != nil {
		logger.Error("ID do usuário inválido", zap.String("user_id", userID), zap.Error(parseErr))
		c.JSON(400, gin.H{
			"message": "ID do usuário inválido. Deve ser um número inteiro válido.",
		})
		return
	}

	// Chama o serviço para buscar o usuário
	user, err := uc.service.FindUserByIDService(uint(idUint))
	if err != nil {
		// Como err é do tipo *rest_err.RestErr, basta acessá-lo diretamente
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

func (uc *userControllerInterface) GetCurrentUser(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := uc.service.FindUserByIDService(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(200, user)
}

// Retorna o histórico de estacionamento do usuário
func (uc *userControllerInterface) GetUserParkingHistory(c *gin.Context) {
	userID := c.Param("id")
	// Chamar serviço que retorna o histórico (exemplo fictício)
	parkingHistory, err := uc.service.GetUserParkingHistoryService(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(200, parkingHistory)
}

// Retorna os veículos cadastrados pelo usuário
func (uc *userControllerInterface) GetUserVehicles(c *gin.Context) {
	userID := c.Param("id")
	vehicles, err := uc.service.GetUserVehiclesService(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(200, vehicles)
}