package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// UpdateAgendamentoService atualiza os dados de um agendamento
func (s *agendamentoDomainService) UpdateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateAgendamento service", zap.Uint("agendamento_id", agendamento.GetID()))

	// Buscar o agendamento existente
	existingAgendamento, err := s.agendamentoRepository.FindAgendamentoByID(agendamento.GetID())
	if err != nil {
		logger.Error("Erro ao buscar agendamento para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos
	existingAgendamento.(*model.AgendamentoDomain).DataHoraReserva = agendamento.GetDataHoraReserva()
	existingAgendamento.(*model.AgendamentoDomain).TipoVaga = agendamento.GetTipoVaga()
	existingAgendamento.(*model.AgendamentoDomain).Status = agendamento.GetStatus()
	existingAgendamento.(*model.AgendamentoDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar no repositório
	updatedAgendamento, updateErr := s.agendamentoRepository.UpdateAgendamento(existingAgendamento.(*model.AgendamentoDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar agendamento no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Agendamento atualizado com sucesso", zap.Uint("agendamento_id", updatedAgendamento.GetID()))
	return updatedAgendamento, nil
}
