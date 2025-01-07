package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteRegistroService exclui um registro de estacionamento pelo ID
func (s *registroEstacionamentoDomainService) DeleteRegistroService(id uint) *rest_err.RestErr {
	logger.Info("Init DeleteRegistro service", zap.Uint("registro_id", id))

	deleteErr := s.repo.DeleteRegistro(id)
	if deleteErr != nil {
		logger.Error("Erro ao excluir registro no repositório", zap.Error(deleteErr))
		return deleteErr
	}

	logger.Info("Registro excluído com sucesso", zap.Uint("registro_id", id))
	return nil
}
