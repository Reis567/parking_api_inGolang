package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// FindAgendamentoByID busca um agendamento pelo ID
func (r *agendamentoRepository) FindAgendamentoByID(id uint) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	var agendamento model.AgendamentoDomain

	if err := r.db.First(&agendamento, id).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Agendamento não encontrado: ID %d", id)
			return nil, rest_err.NewNotFoundError("Agendamento não encontrado")
		}
		log.Printf("Erro ao buscar agendamento por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar agendamento", err)
	}
	return &agendamento, nil
}
