package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// CreateVeiculo insere um novo veículo no banco de dados
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
