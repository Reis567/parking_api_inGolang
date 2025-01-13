package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// CreateAgendamento insere um novo agendamento no banco de dados
func (r *agendamentoRepository) CreateAgendamento(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(agendamento).Error; err != nil {
		log.Printf("Erro ao inserir agendamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao criar agendamento", err)
	}
	return agendamento, nil
}
