package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// veiculoRepository é a estrutura que implementa a interface VeiculoRepository
type veiculoRepository struct {
	db *gorm.DB
}

// NewVeiculoRepository cria uma nova instância de veiculoRepository
func NewVeiculoRepository() VeiculoRepository {
	return &veiculoRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewVeiculoRepositoryWithDB(customDB *gorm.DB) VeiculoRepository {
	return &veiculoRepository{db: customDB}
}

// VeiculoRepository interface define os métodos para gerenciar Veículos
type VeiculoRepository interface {
	CreateVeiculo(veiculo model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr)
	FindVeiculoByID(id uint) (model.VehicleDomainInterface, *rest_err.RestErr)
	FindAllVeiculos() ([]model.VehicleDomainInterface, *rest_err.RestErr)
	UpdateVeiculo(veiculo *model.VehicleDomain) (*model.VehicleDomain, *rest_err.RestErr)
	DeleteVeiculo(id uint) *rest_err.RestErr
}
