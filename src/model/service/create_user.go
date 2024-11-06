package service

import (
	"fmt"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"time"

	"go.uber.org/zap"
)

// CreateUser cria um novo usuário no sistema
func (s *userDomainService) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser model", zap.String("journey", "Create user"))

	// Atribuir ID e datas
	user.(*model.UserDomain).ID = generateID()
	user.(*model.UserDomain).CreatedAt = time.Now().Format(time.RFC3339)
	user.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Encriptar a senha
	user.(*model.UserDomain).EncryptPassword()
	userDomainRepository , err := s.userRepository.CreateUser(user)
	if err != nil {
		logger.Error("Erro ao salvar usuário no banco de dados", zap.Error(err))
		return nil, err
	}

	// Log e retorno do usuário criado (simulação)
	logger.Info("Usuário criado com sucesso", zap.String("user_id", user.GetID()))
	return userDomainRepository, nil
}

// Exemplo de função para geração de ID
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
