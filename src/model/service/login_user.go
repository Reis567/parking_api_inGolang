package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

func (s *userDomainService) LoginUserService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "Login user"))

	// Buscar usuário pelo email
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Erro ao buscar usuário no banco de dados", zap.Error(err))
		return nil, rest_err.NewUnauthorizedError("Credenciais inválidas")
	}

	// Verificar a senha
	if !user.CheckPassword(password) {
		logger.Error("Senha inválida para o usuário", zap.String("email", email))
		return nil, rest_err.NewUnauthorizedError("Credenciais inválidas")
	}

	logger.Info("Usuário autenticado com sucesso", zap.String("email", email))
	return user, nil
}
