package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/validation"
	"meu-novo-projeto/src/model"

	"meu-novo-projeto/src/controller/model/request"
	"go.uber.org/zap"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

var (
	validate = validator.New()

)

// CreateUser é responsável por criar um novo usuário
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	// Criar o usuário usando o serviço de domínio
	user := model.NewUserDomain(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password, userRequest.Age)
	createdUser, err := uc.service.CreateUser(user)

	if err != nil {
		logger.Error("Erro ao criar o usuário", zap.Error(err))
		c.JSON(err.Code, err)
		return
	}

	// Log de sucesso ao criar o usuário
	logger.Info("Usuário criado com sucesso", zap.String("user_email", createdUser.GetEmail()))

	// Retornar resposta de sucesso
	c.JSON(201, gin.H{
		"message": "Usuário criado com sucesso!",
		"user": gin.H{
			"id":         createdUser.GetID(),
			"first_name": createdUser.GetFirstName(),
			"last_name":  createdUser.GetLastName(),
			"email":      createdUser.GetEmail(),
			"age":        createdUser.GetAge(),
			"created_at": createdUser.GetCreatedAt(),
			"updated_at": createdUser.GetUpdatedAt(),
		},
	})
}
