package user

import (
	"fmt"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/configuration/validation"
	"meu-novo-projeto/src/controller/model/request"
	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

// Instancia do validador
var validate = validator.New()

// CreateUser é responsável por criar um novo usuário
func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	logger.Info("Iniciando CreateUserController")

	// Tentar fazer o binding do JSON para o struct UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao fazer o binding do JSON para UserRequest", zap.Error(err))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Validar os dados usando o validator
	if err := validate.Struct(userRequest); err != nil {
		logger.Error("Erro de validação dos dados do usuário", zap.Error(err))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Simulação de lógica para criar o usuário
	fmt.Println(userRequest)

	// Log de sucesso ao criar o usuário
	logger.Info("Usuário criado com sucesso", zap.String("user_email", userRequest.Email))

	// Retornar resposta de sucesso
	c.JSON(201, gin.H{
		"message": "Usuário criado com sucesso!",
		"user":    userRequest,
	})
}
