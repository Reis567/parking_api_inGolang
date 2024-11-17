package user

import (
	"bytes"                              // Para criar um buffer para reposição do corpo da requisição
	"io"                                 // Para trabalhar com a interface io.NopCloser
	"meu-novo-projeto/src/configuration/logger" // Para o logger

	"meu-novo-projeto/src/model"         // Para usar o domínio do usuário
	"meu-novo-projeto/src/controller/model/request" // Para o modelo do request
	"github.com/gin-gonic/gin"           // Para o framework Gin
	"github.com/go-playground/validator/v10" // Para validação dos dados de entrada
	"go.uber.org/zap"                    // Para logging estruturado
)


var (
	validate = validator.New()

)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
    logger.Info("Iniciando CreateUserController")

    // Ler o corpo da requisição para debug
    rawData, err := c.GetRawData()
    if err != nil {
        logger.Error("Erro ao obter dados brutos da requisição", zap.Error(err))
        c.JSON(400, gin.H{"message": "Erro ao processar a requisição"})
        return
    }
    logger.Info("Raw JSON recebido", zap.ByteString("raw_data", rawData))

    // Reposicionar o corpo para o binding funcionar
    c.Request.Body = io.NopCloser(bytes.NewBuffer(rawData))

    // Mapear JSON para estrutura
    var userRequest request.UserRequest
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        logger.Error("Erro ao fazer o binding do JSON para UserRequest", zap.Error(err))
        c.JSON(400, gin.H{"message": "Dados inválidos", "error": err.Error()})
        return
    }
    logger.Info("Dados decodificados para UserRequest", zap.Any("userRequest", userRequest))

    // Validação adicional
    if err := validate.Struct(userRequest); err != nil {
        logger.Error("Erro de validação dos dados do usuário", zap.Error(err))
        c.JSON(400, gin.H{"message": "Erro de validação", "error": err.Error()})
        return
    }

    // Criar usuário no domínio
    user := model.NewUserDomain(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password, userRequest.Age)
    createdUser, restErr := uc.service.CreateUserService(user)
    if restErr != nil {
        logger.Error("Erro ao criar o usuário", zap.Error(restErr))
        c.JSON(400, gin.H{"message": "Erro ao criar o usuário", "error": restErr.Message})
        return
    }

    // Log de sucesso
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
