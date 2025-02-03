package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"go.uber.org/zap"
	"time"
	"meu-novo-projeto/src/configuration/logger"
)

// NewVagaDomainService cria uma instância de vagaDomainService
func NewVagaDomainService(vagaRepository repository.VagaRepository) VagaDomainService {
	return &vagaDomainService{vagaRepository}
}

type vagaDomainService struct {
	vagaRepository repository.VagaRepository
}

// Interface do serviço de domínio das vagas
type VagaDomainService interface {
	CreateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr)
	FindVagaByIDService(id uint) (model.VagaDomainInterface, *rest_err.RestErr)
	FindAllVagasService() ([]model.VagaDomainInterface, *rest_err.RestErr)
	UpdateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr)
	DeleteVagaService(id uint) *rest_err.RestErr

	// Novos métodos:
	BuscarVagaDisponivelService(tipo string) (model.VagaDomainInterface, *rest_err.RestErr)
	CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	AtualizarStatusVagaService(id uint, status string) *rest_err.RestErr
}



func (s *vagaDomainService) CreateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateVagaService", zap.String("journey", "Create vaga"))

	// Atribuir timestamps de criação e atualização
	vaga.(*model.VagaDomain).CreatedAt = time.Now().Format(time.RFC3339)
	vaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Criar a vaga no repositório
	createdVaga, err := s.vagaRepository.CreateVaga(vaga)
	if err != nil {
		logger.Error("Erro ao criar vaga no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Vaga criada com sucesso", zap.String("vaga_id", createdVaga.GetSerial()))
	return createdVaga, nil
}


func (s *vagaDomainService) FindVagaByIDService(id uint) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindVagaByIDService", zap.Uint("vaga_id", id))

	vaga, err := s.vagaRepository.FindVagaByID(id)
	if err != nil {
		logger.Error("Erro ao buscar vaga no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Vaga encontrada com sucesso", zap.Uint("vaga_id", id))
	return vaga, nil
}

// FindAllVagasService busca todas as vagas
func (s *vagaDomainService) FindAllVagasService() ([]model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllVagasService")

	vagas, err := s.vagaRepository.FindAllVagas()
	if err != nil {
		logger.Error("Erro ao buscar todas as vagas no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Vagas encontradas com sucesso", zap.Int("total", len(vagas)))
	return vagas, nil
}



func (s *vagaDomainService) UpdateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateVagaService", zap.Uint("vaga_id", vaga.GetID()))

	// Buscar vaga existente
	existingVaga, err := s.vagaRepository.FindVagaByID(vaga.GetID())
	if err != nil {
		logger.Error("Erro ao buscar vaga para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos necessários
	existingVaga.(*model.VagaDomain).Tipo = vaga.GetTipo()
	existingVaga.(*model.VagaDomain).Status = vaga.GetStatus()
	existingVaga.(*model.VagaDomain).Localizacao = vaga.GetLocalizacao()
	existingVaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar as alterações
	updatedVaga, updateErr := s.vagaRepository.UpdateVaga(existingVaga.(*model.VagaDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar vaga no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Vaga atualizada com sucesso", zap.Uint("vaga_id", updatedVaga.GetID()))
	return updatedVaga, nil
}


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



// BuscarVagaDisponivelService busca a primeira vaga disponível para o tipo informado.
func (s *vagaDomainService) BuscarVagaDisponivelService(tipo string) (model.VagaDomainInterface, *rest_err.RestErr) {
	// Exemplo: Chame um método do repositório que retorne uma vaga disponível
	vaga, err := s.vagaRepository.FindVagaDisponivel(tipo)
	if err != nil {
		logger.Error("Erro ao buscar vaga disponível", zap.Error(err))
		return nil, err
	}
	return vaga, nil
}

// CreateRegistroService salva o registro de entrada no estacionamento.
func (s *vagaDomainService) CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	// Exemplo: Se você tiver um repositório para registros, delegue a criação para ele.
	// Caso contrário, implemente a lógica necessária para salvar o registro.
	createdRegistro, err := s.vagaRepository.CreateRegistro(registro)
	if err != nil {
		logger.Error("Erro ao criar registro", zap.Error(err))
		return nil, err
	}
	return createdRegistro, nil
}

// AtualizarStatusVagaService atualiza o status da vaga.
func (s *vagaDomainService) AtualizarStatusVagaService(id uint, status string) *rest_err.RestErr {
	// Buscar a vaga existente
	vaga, err := s.vagaRepository.FindVagaByID(id)
	if err != nil {
		logger.Error("Erro ao buscar vaga para atualizar status", zap.Error(err))
		return err
	}

	// Atualizar o status
	vaga.(*model.VagaDomain).Status = status
	vaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	_, updateErr := s.vagaRepository.UpdateVaga(vaga.(*model.VagaDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar status da vaga", zap.Error(updateErr))
		return updateErr
	}

	return nil
}
