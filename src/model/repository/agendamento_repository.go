package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// agendamentoRepository é a estrutura que implementa a interface AgendamentoRepository
type agendamentoRepository struct {
	db *gorm.DB
}

// NewAgendamentoRepository cria uma nova instância de agendamentoRepository
func NewAgendamentoRepository() AgendamentoRepository {
	return &agendamentoRepository{db: database.DB}
}

func NewAgendamentoRepositoryWithDB(customDB *gorm.DB) AgendamentoRepository {
	return &agendamentoRepository{db: customDB}
}

// AgendamentoRepository interface define os métodos para gerenciar Agendamentos
type AgendamentoRepository interface {
	CreateAgendamento(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAgendamentoByID(id uint) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAllAgendamentos() ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
	UpdateAgendamento(agendamento *model.AgendamentoDomain) (*model.AgendamentoDomain, *rest_err.RestErr)
	DeleteAgendamento(id uint) *rest_err.RestErr
}
