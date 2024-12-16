package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// UpdateVeiculo atualiza os dados de um veículo
func (r *veiculoRepository) UpdateVeiculo(veiculo *model.VehicleDomain) (*model.VehicleDomain, *rest_err.RestErr) {
	if err := r.db.Save(veiculo).Error; err != nil {
		log.Printf("Erro ao atualizar veículo no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar veículo", err)
	}

	return veiculo, nil
}
