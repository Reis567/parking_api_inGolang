package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

func (s *userDomainService) LoginUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "Login user"))

	// Buscar usuário pelo email
	storedUser, err := s.userRepository.FindUserByEmail(user.GetEmail())
	if err != nil {
		logger.Error("Erro ao buscar usuário no banco de dados", zap.Error(err))
		return nil, rest_err.NewUnauthorizedError("Credenciais inválidas")
	}

	// Verificar a senha
	if !storedUser.CheckPassword(user.GetPassword()) {
		logger.Error("Senha inválida para o usuário", zap.String("email", user.GetEmail()))
		return nil, rest_err.NewUnauthorizedError("Credenciais inválidas")
	}

	logger.Info("Usuário autenticado com sucesso", zap.String("email", user.GetEmail()))
	return storedUser, nil
}
