package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteAgendamentoService exclui um agendamento pelo ID
func (s *agendamentoDomainService) DeleteAgendamentoService(id uint) *rest_err.RestErr {
	logger.Info("Init DeleteAgendamento service", zap.Uint("agendamento_id", id))

	err := s.agendamentoRepository.DeleteAgendamento(id)
	if err != nil {
		logger.Error("Erro ao excluir agendamento no repositório", zap.Error(err))
		return err
	}

	logger.Info("Agendamento excluído com sucesso", zap.Uint("agendamento_id", id))
	return nil
}
