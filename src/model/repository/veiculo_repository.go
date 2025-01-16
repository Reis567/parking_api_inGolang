package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
	"log"
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
func (r *veiculoRepository) CreateVeiculo(veiculo model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(veiculo).Error; err != nil {
		log.Printf("Erro ao inserir veículo no banco de dados: %v", err)

		// Logar o valor do JSON do veículo em caso de erro
		jsonValue, jsonErr := veiculo.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter veículo para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados do veículo: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar veículo", err)
	}

	return veiculo, nil
}


func (r *veiculoRepository) FindVeiculoByID(id uint) (model.VehicleDomainInterface, *rest_err.RestErr) {
	var veiculo model.VehicleDomain

	if err := r.db.First(&veiculo, id).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Veículo não encontrado: ID %d", id)
			return nil, rest_err.NewNotFoundError("Veículo não encontrado")
		}
		log.Printf("Erro ao buscar veículo por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar veículo", err)
	}

	return &veiculo, nil
}

// FindAllVeiculos busca todos os veículos
func (r *veiculoRepository) FindAllVeiculos() ([]model.VehicleDomainInterface, *rest_err.RestErr) {
	var veiculos []model.VehicleDomain
	if err := r.db.Find(&veiculos).Error; err != nil {
		log.Printf("Erro ao buscar veículos no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar veículos", err)
	}

	// Converte para interfaces
	veiculoInterfaces := make([]model.VehicleDomainInterface, len(veiculos))
	for i, v := range veiculos {
		veiculoInterfaces[i] = &v
	}

	return veiculoInterfaces, nil
}


func (r *veiculoRepository) UpdateVeiculo(veiculo *model.VehicleDomain) (*model.VehicleDomain, *rest_err.RestErr) {
	if err := r.db.Save(veiculo).Error; err != nil {
		log.Printf("Erro ao atualizar veículo no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar veículo", err)
	}

	return veiculo, nil
}


func (r *veiculoRepository) DeleteVeiculo(id uint) *rest_err.RestErr {
	var veiculo model.VehicleDomain

	if err := r.db.First(&veiculo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Veículo não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Veículo não encontrado para exclusão")
		}
		log.Printf("Erro ao buscar veículo para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar veículo para exclusão", err)
	}

	if err := r.db.Delete(&veiculo).Error; err != nil {
		log.Printf("Erro ao excluir veículo do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir veículo", err)
	}

	log.Printf("Veículo excluído com sucesso: ID %d", id)
	return nil
}
