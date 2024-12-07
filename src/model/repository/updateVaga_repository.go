package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// UpdateVaga atualiza os dados de uma vaga existente
func (r *vagaRepository) UpdateVaga(vaga *model.VagaDomain) (*model.VagaDomain, *rest_err.RestErr) {
	// Verificar se a vaga existe antes de atualizar
	if err := r.db.First(&model.VagaDomain{}, vaga.ID).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Vaga não encontrada para atualização: ID %d", vaga.ID)
			return nil, rest_err.NewNotFoundError("Vaga não encontrada para atualização")
		}
		log.Printf("Erro ao buscar vaga para atualização: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vaga para atualização", err)
	}

	// Atualizar os dados da vaga
	if err := r.db.Save(vaga).Error; err != nil {
		log.Printf("Erro ao atualizar vaga no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar vaga", err)
	}

	log.Printf("Vaga atualizada com sucesso: ID %d", vaga.ID)
	return vaga, nil
}
