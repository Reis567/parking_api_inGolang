package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// DeleteVaga exclui uma vaga pelo ID
func (r *vagaRepository) DeleteVaga(id uint) *rest_err.RestErr {
	var vaga model.VagaDomain

	// Buscar a vaga para garantir que ela existe
	if err := r.db.First(&vaga, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Vaga não encontrada para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Vaga não encontrada para exclusão")
		}
		log.Printf("Erro ao buscar vaga para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar vaga para exclusão", err)
	}

	// Excluir a vaga
	if err := r.db.Delete(&vaga).Error; err != nil {
		log.Printf("Erro ao excluir vaga do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir vaga", err)
	}

	log.Printf("Vaga excluída com sucesso: ID %d", id)
	return nil
}
