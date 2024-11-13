package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// CreateUser cria um novo usuário no sistema
func (s *userDomainService) CreateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser model", zap.String("journey", "Create user"))

	// Atribuir datas (ID será gerenciado automaticamente pelo GORM)
	user.(*model.UserDomain).CreatedAt = time.Now().Format(time.RFC3339)
	user.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Encriptar a senha
	user.(*model.UserDomain).EncryptPassword()

	// Salvar o usuário no repositório
	userDomainRepository, err := s.userRepository.CreateUser(user)
	if err != nil {
		logger.Error("Erro ao salvar usuário no banco de dados", zap.Error(err))
		return nil, err
	}

	// Log de sucesso
	logger.Info("Usuário criado com sucesso", zap.Uint("user_id", userDomainRepository.GetID()))
	return userDomainRepository, nil
}
