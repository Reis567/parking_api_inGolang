package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
		"meu-novo-projeto/src/configuration/logger"
	"go.uber.org/zap"
	"time"
	"os"
	"strconv"
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




func (s *registroEstacionamentoDomainService) FinalizarEstacionamentoService(registroID uint, horaSaida string) (interface{}, *rest_err.RestErr) {
	// Buscar o registro de estacionamento pelo ID
	registro, err := s.repo.FindRegistroByID(registroID)
	if err != nil {
		return nil, err
	}

	// Obter o valor da tarifa por hora da variável de ambiente "TARIFFA_HORA"
	tarifaStr := os.Getenv("TARIFFA_HORA")
	if tarifaStr == "" {
		tarifaStr = "5.0" // valor padrão caso a variável não esteja definida
	}
	valorPorHora, err := strconv.ParseFloat(tarifaStr, 64)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao converter a tarifa por hora", err)
	}

	// Calcular o tempo decorrido e o valor a ser cobrado
	horaEntrada, err := time.Parse(time.RFC3339, registro.GetHoraEntrada())
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao processar hora de entrada", err)
	}
	saida, err := time.Parse(time.RFC3339, horaSaida)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao processar hora de saída", err)
	}
	duracao := saida.Sub(horaEntrada)

	horas := duracao.Hours()
	if horas < 1 {
		horas = 1
	}
	valorCobrado := valorPorHora * horas

	// Atualizar o registro com os dados da saída
	registro.RegistrarSaida(horaSaida, valorCobrado)
	updatedRegistro, updateErr := s.repo.UpdateRegistro(registro.(*model.RegistroEstacionamentoDomain))
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