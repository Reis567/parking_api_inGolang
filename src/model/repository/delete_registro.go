package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// DeleteRegistro exclui um registro de estacionamento pelo ID
func (r *registroEstacionamentoRepository) DeleteRegistro(id uint) *rest_err.RestErr {
	var registro model.RegistroEstacionamentoDomain

	if err := r.db.First(&registro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Registro não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Registro não encontrado")
		}
		log.Printf("Erro ao buscar registro para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar registro para exclusão", err)
	}

	if err := r.db.Delete(&registro).Error; err != nil {
		log.Printf("Erro ao excluir registro: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir registro", err)
	}
	log.Printf("Registro excluído com sucesso: ID %d", id)
	return nil
}
