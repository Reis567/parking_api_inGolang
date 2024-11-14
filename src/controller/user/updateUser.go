package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/validation"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Iniciando UpdateUserController")

	// Obter o ID do usuário dos parâmetros da URL
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Error("ID do usuário inválido", zap.String("user_id", userIDStr), zap.Error(err))
		c.JSON(400, gin.H{
			"message": "ID do usuário inválido. Deve ser um número inteiro.",
		})
		return
	}

	// Fazer o binding do corpo da requisição para o struct UserRequest
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao fazer o binding do JSON para UserRequest", zap.Error(err))
		validationErr := validation.ValidateUserError(err)
		c.JSON(validationErr.Code, validationErr)
		return
	}

	// Criar o domínio do usuário com os dados recebidos
	user := model.NewUserDomain(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	).(*model.UserDomain)

	// Atribuir o ID do usuário ao domínio
	user.ID = uint(userID)

	// Chamar o serviço para atualizar o usuário
	updatedUser, updateErr := uc.service.UpdateUserService(user)
	if updateErr != nil {
		logger.Error("Erro ao atualizar usuário", zap.Error(updateErr))
		c.JSON(updateErr.Code, updateErr)
		return
	}

	// Retornar resposta de sucesso
	c.JSON(200, gin.H{
		"message": "Usuário atualizado com sucesso!",
		"user": gin.H{
			"id":         updatedUser.GetID(),
			"first_name": updatedUser.GetFirstName(),
			"last_name":  updatedUser.GetLastName(),
			"email":      updatedUser.GetEmail(),
			"age":        updatedUser.GetAge(),
			"created_at": updatedUser.GetCreatedAt(),
			"updated_at": updatedUser.GetUpdatedAt(),
		},
	})
}
