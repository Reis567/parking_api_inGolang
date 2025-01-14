package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// CreateAgendamentoService cria um novo agendamento no sistema
func (s *agendamentoDomainService) CreateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateAgendamento service")

	createdAgendamento, err := s.agendamentoRepository.CreateAgendamento(agendamento)
	if err != nil {
		logger.Error("Erro ao criar agendamento no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Agendamento criado com sucesso", zap.Uint("agendamento_id", createdAgendamento.GetID()))
	return createdAgendamento, nil
}
