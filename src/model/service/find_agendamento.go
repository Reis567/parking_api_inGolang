package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindAgendamentoByIDService busca um agendamento pelo ID
func (s *agendamentoDomainService) FindAgendamentoByIDService(id uint) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAgendamentoByID service", zap.Uint("agendamento_id", id))

	agendamento, err := s.agendamentoRepository.FindAgendamentoByID(id)
	if err != nil {
		logger.Error("Erro ao buscar agendamento no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Agendamento encontrado com sucesso", zap.Uint("agendamento_id", id))
	return agendamento, nil
}

// FindAllAgendamentosService busca todos os agendamentos
func (s *agendamentoDomainService) FindAllAgendamentosService() ([]model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllAgendamentos service")

	agendamentos, err := s.agendamentoRepository.FindAllAgendamentos()
	if err != nil {
		logger.Error("Erro ao buscar todos os agendamentos no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Todos os agendamentos encontrados com sucesso", zap.Int("total", len(agendamentos)))
	return agendamentos, nil
}
