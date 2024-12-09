package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
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
}
