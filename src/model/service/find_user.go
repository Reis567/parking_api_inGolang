package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindUserByID busca um usuário pelo ID
func (s *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID service", zap.String("journey", "Find user by ID"))

	// Consulta o repositório para buscar o usuário pelo ID
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		logger.Error("Erro ao buscar usuário no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Usuário encontrado com sucesso", zap.String("user_id", id))
	return user, nil
}

// FindUserByEmail busca um usuário pelo email
func (s *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "Find user by email"))

	// Consulta o repositório para buscar o usuário pelo email
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Erro ao buscar usuário no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Usuário encontrado com sucesso", zap.String("user_email", email))
	return user, nil
}