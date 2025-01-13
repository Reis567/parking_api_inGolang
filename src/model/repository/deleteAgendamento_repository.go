package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// DeleteAgendamento exclui um agendamento pelo ID
func (r *agendamentoRepository) DeleteAgendamento(id uint) *rest_err.RestErr {
	var agendamento model.AgendamentoDomain

	if err := r.db.First(&agendamento, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Agendamento não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Agendamento não encontrado para exclusão")
		}
		log.Printf("Erro ao buscar agendamento para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar agendamento para exclusão", err)
	}

	if err := r.db.Delete(&agendamento).Error; err != nil {
		log.Printf("Erro ao excluir agendamento do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir agendamento", err)
	}
	return nil
}
