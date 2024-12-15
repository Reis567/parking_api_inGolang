package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// FindVeiculoByID busca um veículo pelo ID
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
