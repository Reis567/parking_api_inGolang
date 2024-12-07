package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/configuration/database"
	"gorm.io/gorm"
)

// vagaRepository é a estrutura que implementa a interface VagaRepository
type vagaRepository struct {
	db *gorm.DB
}

// NewVagaRepository cria uma nova instância de vagaRepository
func NewVagaRepository() VagaRepository {
	return &vagaRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewVagaRepositoryWithDB(customDB *gorm.DB) VagaRepository {
	return &vagaRepository{db: customDB}
}

// VagaRepository interface define os métodos para gerenciar Vagas
type VagaRepository interface {
	CreateVaga(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr)
	FindVagaByID(id uint) (model.VagaDomainInterface, *rest_err.RestErr)
	UpdateVaga(vaga *model.VagaDomain) (*model.VagaDomain, *rest_err.RestErr)
	DeleteVaga(id uint) *rest_err.RestErr
}
