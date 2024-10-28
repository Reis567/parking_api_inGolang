package model

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"time"

	"go.uber.org/zap"
)

// UpdateUser atualiza um usuário existente
func (ud *UserDomain) UpdateUser(user UserDomain) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init UpdateUser model", zap.String("journey", "Update user"))

	// Simulação de atualização (substituir pela lógica real de banco de dados)
	if user.ID != ud.ID {
		logger.Error("User not found for update", zap.String("user_id", user.ID))
		return nil, rest_err.NewNotFoundError("Usuário não encontrado para atualização")
	}

	// Atualizar campos
	ud.FirstName = user.FirstName
	ud.LastName = user.LastName
	ud.Email = user.Email
	ud.Password = user.Password
	ud.Age = user.Age
	ud.UpdatedAt = time.Now().Format(time.RFC3339)

	logger.Info("Usuário atualizado com sucesso", zap.String("user_id", user.ID))
	return ud, nil
}
