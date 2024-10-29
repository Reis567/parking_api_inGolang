package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"time"

	"go.uber.org/zap"
)

// UpdateUser atualiza um usuário existente
func (s *userDomainService) UpdateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateUser service", zap.String("journey", "Update user"))

	// Simulação de atualização (substituir pela lógica real de banco de dados)
	// Aqui você implementaria a lógica de atualização no banco de dados
	if user.GetID() == "" {
		logger.Error("User not found for update", zap.String("user_id", user.GetID()))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado para atualização")
	}

	// Simular a atualização dos campos e a data de atualização
	userUpdated := model.NewUserDomain(user.GetFirstName(), user.GetLastName(), user.GetEmail(), user.GetPassword(), user.GetAge())
	userUpdated.(*model.UserDomain).ID = user.GetID()
	userUpdated.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	logger.Info("Usuário atualizado com sucesso", zap.String("user_id", user.GetID()))
	return userUpdated, nil
}
