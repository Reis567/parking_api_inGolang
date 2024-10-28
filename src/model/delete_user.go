package model

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteUser exclui um usuário pelo ID
func (ud *UserDomain) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "Delete user"))

	// Simulação de exclusão (substituir pela lógica real de banco de dados)
	if id != ud.ID {
		logger.Error("User not found for deletion", zap.String("user_id", id))
		return rest_err.NewNotFoundError("Usuário não encontrado para exclusão")
	}

	logger.Info("Usuário excluído com sucesso", zap.String("user_id", id))
	return nil
}
