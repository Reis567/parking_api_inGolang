package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"time"
	"go.uber.org/zap"
	"meu-novo-projeto/src/configuration/logger"
)

// NewAgendamentoDomainService cria uma instância de agendamentoDomainService
func NewAgendamentoDomainService(agendamentoRepository repository.AgendamentoRepository) AgendamentoDomainService {
	return &agendamentoDomainService{agendamentoRepository}
}

type agendamentoDomainService struct {
	agendamentoRepository repository.AgendamentoRepository
}

// AgendamentoDomainService define os métodos para o serviço de agendamento
type AgendamentoDomainService interface {
	CreateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAgendamentoByIDService(id uint) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAllAgendamentosService() ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
	UpdateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	DeleteAgendamentoService(id uint) *rest_err.RestErr
	//FindAgendamentosByPeriod(inicio, fim time.Time) ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
	//CancelAgendamentoService(id uint) *rest_err.RestErr
}


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
