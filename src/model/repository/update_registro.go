package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// UpdateRegistro atualiza os dados de um registro de estacionamento
func (r *registroEstacionamentoRepository) UpdateRegistro(registro *model.RegistroEstacionamentoDomain) (*model.RegistroEstacionamentoDomain, *rest_err.RestErr) {
	if err := r.db.Save(registro).Error; err != nil {
		log.Printf("Erro ao atualizar registro no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar registro", err)
	}
	return registro, nil
}
