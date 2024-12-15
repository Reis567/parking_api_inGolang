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

