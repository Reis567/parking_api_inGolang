package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
)

// NewRegistroEstacionamentoDomainService cria uma instância de registroEstacionamentoDomainService
func NewRegistroEstacionamentoDomainService(repo repository.RegistroEstacionamentoRepository) RegistroEstacionamentoDomainService {
	return &registroEstacionamentoDomainService{repo}
}

type registroEstacionamentoDomainService struct {
	repo repository.RegistroEstacionamentoRepository
}

// RegistroEstacionamentoDomainService define os métodos para o serviço de registro de estacionamento
type RegistroEstacionamentoDomainService interface {
	CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindRegistroByIDService(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindAllRegistrosService() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	UpdateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	DeleteRegistroService(id uint) *rest_err.RestErr
}
