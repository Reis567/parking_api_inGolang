package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"strconv" 
	"go.uber.org/zap"

)


func (s *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "Delete user"))

	// Converte o ID para uint
	idUint, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("ID inválido para exclusão", zap.String("user_id", id), zap.Error(err))
		return rest_err.NewBadRequestError("ID inválido. Deve ser um número inteiro.")
	}

	// Chama o repositório para excluir o usuário
	deleteErr := s.userRepository.DeleteUser(uint(idUint))
	if deleteErr != nil {
		logger.Error("Erro ao excluir usuário no repositório", zap.Error(deleteErr))
		return deleteErr
	}

	logger.Info("Usuário excluído com sucesso", zap.String("user_id", id))
	return nil
}
