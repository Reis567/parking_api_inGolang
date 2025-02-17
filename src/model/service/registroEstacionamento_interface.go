package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
		"meu-novo-projeto/src/configuration/logger"
	"go.uber.org/zap"
	"time"

)

// NewRegistroEstacionamentoDomainService cria uma instância de registroEstacionamentoDomainService
func NewRegistroEstacionamentoDomainService(repo repository.RegistroEstacionamentoRepository,vagaRepo repository.VagaRepository) RegistroEstacionamentoDomainService {
	return &registroEstacionamentoDomainService{repo,vagaRepo,}
}

type registroEstacionamentoDomainService struct {
	repo repository.RegistroEstacionamentoRepository
	vagaRepo repository.VagaRepository
}

// RegistroEstacionamentoDomainService define os métodos para o serviço de registro de estacionamento
type RegistroEstacionamentoDomainService interface {
	CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindRegistroByIDService(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindAllRegistrosService() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	UpdateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	DeleteRegistroService(id uint) *rest_err.RestErr
	FindRegistrosPorDataService(data time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) 
	FindHistoricoRegistrosService(dataInicio, dataFim time.Time, placa, status string) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)

}
func (s *registroEstacionamentoDomainService) CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateRegistro service", zap.String("journey", "Create registro"))

	createdRegistro, err := s.repo.CreateRegistro(registro)
	if err != nil {
		logger.Error("Erro ao criar registro no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registro criado com sucesso", zap.Uint("registro_id", createdRegistro.GetID()))
	return createdRegistro, nil
}


func (s *registroEstacionamentoDomainService) FindAllRegistrosService() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllRegistros service")

	registros, err := s.repo.FindAllRegistros()
	if err != nil {
		logger.Error("Erro ao buscar todos os registros no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registros retornados com sucesso", zap.Int("count", len(registros)))
	return registros, nil
}
func (s *registroEstacionamentoDomainService) FindRegistroByIDService(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindRegistroByID service", zap.Uint("registro_id", id))

	registro, err := s.repo.FindRegistroByID(id)
	if err != nil {
		logger.Error("Erro ao buscar registro no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registro encontrado com sucesso", zap.Uint("registro_id", registro.GetID()))
	return registro, nil
}


func (s *registroEstacionamentoDomainService) UpdateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateRegistro service", zap.Uint("registro_id", registro.GetID()))

	existingRegistro, err := s.repo.FindRegistroByID(registro.GetID())
	if err != nil {
		logger.Error("Erro ao buscar registro para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar campos
	existingRegistro.(*model.RegistroEstacionamentoDomain).HoraSaida = registro.GetHoraSaida()
	existingRegistro.(*model.RegistroEstacionamentoDomain).Status = registro.GetStatus()
	existingRegistro.(*model.RegistroEstacionamentoDomain).ValorCobrado = registro.GetValorCobrado()
	existingRegistro.(*model.RegistroEstacionamentoDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	updatedRegistro, updateErr := s.repo.UpdateRegistro(existingRegistro.(*model.RegistroEstacionamentoDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar registro no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Registro atualizado com sucesso", zap.Uint("registro_id", updatedRegistro.GetID()))
	return updatedRegistro, nil
}

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


func (s *registroEstacionamentoDomainService) FindRegistrosPorDataService(data time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindRegistrosPorData service", zap.Time("data", data))

	registros, err := s.repo.FindRegistrosPorData(data)
	if err != nil {
		logger.Error("Erro ao buscar registros no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registros encontrados com sucesso", zap.Int("count", len(registros)))
	return registros, nil
}


func (s *registroEstacionamentoDomainService) FindHistoricoRegistrosService(dataInicio, dataFim time.Time, placa, status string) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
    logger.Info("Init FindHistoricoRegistrosService", zap.Time("dataInicio", dataInicio), zap.Time("dataFim", dataFim), zap.String("placa", placa), zap.String("status", status))

    // Buscar registros pelo período
    registros, err := s.repo.FindRegistrosPorPeriodo(dataInicio, dataFim)
    if err != nil {
        logger.Error("Erro ao buscar registros por período", zap.Error(err))
        return nil, err
    }

    // Filtrar adicionalmente por placa e status, se fornecidos
    var filtrados []model.RegistroEstacionamentoDomainInterface
    for _, registro := range registros {
        // Se a placa for informada e não coincidir, pular
        if placa != "" && registro.GetPlaca() != placa {
            continue
        }
        // Se o status for informado e não coincidir, pular
        if status != "" && registro.GetStatus() != status {
            continue
        }
        filtrados = append(filtrados, registro)
    }

    logger.Info("Registros filtrados encontrados", zap.Int("count", len(filtrados)))
    return filtrados, nil
}
