package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteVagaService exclui uma vaga pelo ID
func (s *vagaDomainService) DeleteVagaService(id uint) *rest_err.RestErr {
	logger.Info("Init DeleteVagaService", zap.Uint("vaga_id", id))

	// Excluir vaga
	deleteErr := s.vagaRepository.DeleteVaga(id)
	if deleteErr != nil {
		logger.Error("Erro ao excluir vaga no repositório", zap.Error(deleteErr))
		return deleteErr
	}

	logger.Info("Vaga excluída com sucesso", zap.Uint("vaga_id", id))
	return nil
}
