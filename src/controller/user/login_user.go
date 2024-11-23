package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/controller/model/request"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/configuration/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Iniciando LoginUserController")

	// Fazer o binding do corpo da requisição
	var loginRequest request.UserLogin
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		logger.Error("Erro ao fazer o binding do JSON para UserLogin", zap.Error(err))
		c.JSON(400, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	logger.Info("Dados decodificados para UserLogin", zap.Any("loginRequest", loginRequest))

	// Criar instância do domínio para o login
	userDomain := model.NewUserLoginDomain(loginRequest.Email, loginRequest.Password)

	// Validar email e senha
	user, err := uc.service.LoginUserService(userDomain)
	if err != nil {
		logger.Error("Erro ao autenticar usuário", zap.Error(err))
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	// Gerar JWT
	token, jwtErr := jwt.GenerateToken(user.GetID(), user.GetEmail())
	if jwtErr != nil {
		logger.Error("Erro ao gerar token JWT", zap.Error(jwtErr))
		c.JSON(500, gin.H{"message": "Erro interno ao gerar token de autenticação"})
		return
	}

	// Retornar o token de autenticação
	logger.Info("Usuário autenticado com sucesso", zap.String("user_email", user.GetEmail()))
	c.JSON(200, gin.H{
		"message": "Usuário logado com sucesso!",
		"token":   token,
	})
}
