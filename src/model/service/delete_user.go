package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteUser exclui um usuário pelo ID
func (s *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "Delete user"))

	// Simulação de exclusão (substituir pela lógica real de banco de dados)
	// Aqui você implementaria a lógica de exclusão no banco de dados
	if id == "" {
		logger.Error("User not found for deletion", zap.String("user_id", id))
		return rest_err.NewNotFoundError("Usuário não encontrado para exclusão")
	}

	logger.Info("Usuário excluído com sucesso", zap.String("user_id", id))
	return nil
}
