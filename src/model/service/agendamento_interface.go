package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"time"
	"go.uber.org/zap"
	"meu-novo-projeto/src/configuration/logger"
	"os"
	"strconv"
)

// NewAgendamentoDomainService cria uma instância de agendamentoDomainService
func NewAgendamentoDomainService(agendamentoRepository repository.AgendamentoRepository,vagaRepo repository.VagaRepository,RgRepo repository.RegistroEstacionamentoRepository) AgendamentoDomainService {
	return &agendamentoDomainService{agendamentoRepository,vagaRepo,RgRepo}
}

type agendamentoDomainService struct {
	agendamentoRepository repository.AgendamentoRepository
	vagaRepo repository.VagaRepository
	RgRepo repository.RegistroEstacionamentoRepository
}

// AgendamentoDomainService define os métodos para o serviço de agendamento
type AgendamentoDomainService interface {
	CreateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAgendamentoByIDService(id uint) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAllAgendamentosService() ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
	UpdateAgendamentoService(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	DeleteAgendamentoService(id uint) *rest_err.RestErr
	VerificarReservaPorPlacaService(placa string) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FinalizarEstacionamentoService(registroID uint, horaSaida string) (interface{}, *rest_err.RestErr)
	FindReservasAtivasService(status string) ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
	CancelAgendamentoService(id uint, justificativa string) (model.AgendamentoDomainInterface, *rest_err.RestErr)

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


func (s *agendamentoDomainService) VerificarReservaPorPlacaService(placa string) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Verificando reserva para a placa", zap.String("placa", placa))

	// Chamar o repositório para verificar se há uma reserva confirmada para a placa
	reserva, err := s.agendamentoRepository.VerificarReservaPorPlaca(placa)
	if err != nil {
		logger.Error("Erro ao verificar reserva no repositório", zap.Error(err))
		return nil, err
	}

	if reserva == nil {
		logger.Info("Nenhuma reserva encontrada para a placa", zap.String("placa", placa))
		return nil, rest_err.NewNotFoundError("Nenhuma reserva encontrada para esta placa")
	}

	logger.Info("Reserva encontrada com sucesso", zap.Uint("reserva_id", reserva.GetID()))
	return reserva, nil
}




func (s *agendamentoDomainService) FinalizarEstacionamentoService(registroID uint, horaSaida string) (interface{}, *rest_err.RestErr) {
	// Buscar o registro de estacionamento pelo ID
	registro, err := s.RgRepo.FindRegistroByID(registroID)
	if err != nil {
		return nil, err
	}

	// Obter o valor da tarifa por hora da variável de ambiente "TARIFFA_HORA"
	tarifaStr := os.Getenv("TARIFFA_HORA")

	valorPorHora, parseErr := strconv.ParseFloat(tarifaStr, 64)
	if parseErr != nil {
		return nil, rest_err.NewInternalServerError("Erro ao converter a tarifa por hora", parseErr)
	}

	// Calcular o tempo decorrido e o valor a ser cobrado
	horaEntrada, timeErr := time.Parse(time.RFC3339, registro.GetHoraEntrada())
	if timeErr != nil {
		return nil, rest_err.NewInternalServerError("Erro ao processar hora de entrada", timeErr)
	}
	saida, timeErr := time.Parse(time.RFC3339, horaSaida)
	if timeErr != nil {
		return nil, rest_err.NewInternalServerError("Erro ao processar hora de saída", timeErr)
	}
	duracao := saida.Sub(horaEntrada)

	horas := duracao.Hours()
	if horas < 1 {
		horas = 1
	}
	valorCobrado := valorPorHora * horas

	// Atualizar o registro com os dados da saída
	registro.RegistrarSaida(horaSaida, valorCobrado)
	updatedRegistro, updateErr := s.RgRepo.UpdateRegistro(registro.(*model.RegistroEstacionamentoDomain))
	if updateErr != nil {
		return nil, updateErr
	}

	// Atualizar a vaga para o status "disponivel"
	vaga, vagaErr := s.vagaRepo.FindVagaByID(registro.GetVagaID())
	if vagaErr != nil {
		return nil, vagaErr
	}
	vaga.(*model.VagaDomain).Status = "disponivel"
	vaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)
	_, vagaUpdateErr := s.vagaRepo.UpdateVaga(vaga.(*model.VagaDomain))
	if vagaUpdateErr != nil {
		return nil, vagaUpdateErr
	}

	// Retornar o registro atualizado ou os dados do cálculo
	return updatedRegistro, nil
}



func (s *agendamentoDomainService) FindReservasAtivasService(status string) ([]model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Buscando reservas ativas", zap.String("status", status))

	reservas, err := s.agendamentoRepository.FindAgendamentosByStatus(status)
	if err != nil {
		logger.Error("Erro ao buscar reservas ativas", zap.Error(err))
		return nil, err
	}

	logger.Info("Reservas encontradas", zap.Int("total", len(reservas)))
	return reservas, nil
}


func (s *agendamentoDomainService) CancelAgendamentoService(id uint, justificativa string) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando cancelamento do agendamento",
		zap.Uint("agendamento_id", id),
		zap.String("justificativa", justificativa),
	)

	// Buscar o agendamento existente
	agendamento, err := s.agendamentoRepository.FindAgendamentoByID(id)
	if err != nil {
		logger.Error("Erro ao buscar agendamento para cancelamento", zap.Error(err))
		return nil, err
	}

	// Atualizar o status para "cancelado"
	agendamento.AtualizarStatus("cancelado")
	// Se desejar registrar a justificativa, você pode adicioná-la a um log ou, se houver campo apropriado, atribuí-la.

	// Persistir a atualização
	updatedAgendamento, updateErr := s.agendamentoRepository.UpdateAgendamento(agendamento.(*model.AgendamentoDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar agendamento para cancelamento", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Agendamento cancelado com sucesso", zap.Uint("agendamento_id", updatedAgendamento.GetID()))
	return updatedAgendamento, nil
}
