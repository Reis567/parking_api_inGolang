package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// UpdateAgendamento atualiza os dados de um agendamento
func (r *agendamentoRepository) UpdateAgendamento(agendamento *model.AgendamentoDomain) (*model.AgendamentoDomain, *rest_err.RestErr) {
	if err := r.db.Save(agendamento).Error; err != nil {
		log.Printf("Erro ao atualizar agendamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar agendamento", err)
	}
	return agendamento, nil
}
