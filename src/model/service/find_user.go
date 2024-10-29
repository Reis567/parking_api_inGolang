package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindUserByID busca um usuário pelo ID
func (s *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID service", zap.String("journey", "Find user by ID"))

	// Simulação de busca (substituir pela lógica real de banco de dados)
	// Aqui você implementaria a lógica para buscar o usuário pelo ID no banco de dados
	if id == "" {
		logger.Error("User not found", zap.String("user_id", id))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}

	user := model.NewUserDomain("John", "Doe", "johndoe@example.com", "password", 30) // Simulação de retorno
	logger.Info("Usuário encontrado com sucesso", zap.String("user_id", id))
	return user, nil
}

// FindUserByEmail busca um usuário pelo email
func (s *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "Find user by email"))

	// Simulação de busca (substituir pela lógica real de banco de dados)
	// Aqui você implementaria a lógica para buscar o usuário pelo email no banco de dados
	if email == "" {
		logger.Error("User not found", zap.String("user_email", email))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}

	user := model.NewUserDomain("John", "Doe", email, "password", 30) // Simulação de retorno
	logger.Info("Usuário encontrado com sucesso", zap.String("user_email", email))
	return user, nil
}
