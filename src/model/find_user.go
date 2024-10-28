package model

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

//ID
func (ud *UserDomain) FindUserByID(id string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByID model", zap.String("journey", "Find user by ID"))

	// Simulação de busca (substituir pela lógica real de banco de dados)
	if id != ud.ID {
		logger.Error("User not found", zap.String("user_id", id))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}

	logger.Info("Usuário encontrado com sucesso", zap.String("user_id", id))
	return ud, nil
}
// EMAIL
func (ud *UserDomain) FindUserByEmail(email string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail model", zap.String("journey", "Find user by email"))

	// Simulação de busca (substituir pela lógica real de banco de dados)
	if email != ud.Email {
		logger.Error("User not found", zap.String("user_email", email))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado")
	}

	logger.Info("Usuário encontrado com sucesso", zap.String("user_email", email))
	return ud, nil
}
