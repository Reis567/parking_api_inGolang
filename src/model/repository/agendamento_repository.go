package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
	"log"
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
	VerificarReservaPorPlaca(placa string) (model.AgendamentoDomainInterface, *rest_err.RestErr)
	FindAgendamentosByStatus(status string) ([]model.AgendamentoDomainInterface, *rest_err.RestErr)
}


func (r *agendamentoRepository) CreateAgendamento(agendamento model.AgendamentoDomainInterface) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(agendamento).Error; err != nil {
		log.Printf("Erro ao inserir agendamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao criar agendamento", err)
	}
	return agendamento, nil
}


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


// FindAllAgendamentos busca todos os agendamentos
func (r *agendamentoRepository) FindAllAgendamentos() ([]model.AgendamentoDomainInterface, *rest_err.RestErr) {
	var agendamentos []model.AgendamentoDomain
	if err := r.db.Find(&agendamentos).Error; err != nil {
		log.Printf("Erro ao buscar agendamentos no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar agendamentos", err)
	}

	agendamentoInterfaces := make([]model.AgendamentoDomainInterface, len(agendamentos))
	for i, a := range agendamentos {
		agendamentoInterfaces[i] = &a
	}

	return agendamentoInterfaces, nil
}

func (r *agendamentoRepository) UpdateAgendamento(agendamento *model.AgendamentoDomain) (*model.AgendamentoDomain, *rest_err.RestErr) {
	if err := r.db.Save(agendamento).Error; err != nil {
		log.Printf("Erro ao atualizar agendamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar agendamento", err)
	}
	return agendamento, nil
}


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


func (r *agendamentoRepository) VerificarReservaPorPlaca(placa string) (model.AgendamentoDomainInterface, *rest_err.RestErr) {
	var reserva model.AgendamentoDomain

	// Verificar se há uma reserva confirmada para a placa
	if err := r.db.Where("placa = ? AND status = ?", placa, "confirmada").First(&reserva).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Nenhuma reserva confirmada encontrada para a placa: %s", placa)
			return nil, nil // Retorna nil indicando que não há reserva
		}
		log.Printf("Erro ao verificar reserva no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao verificar reserva", err)
	}

	return &reserva, nil
}


// Na interface AgendamentoRepository, adicione:


// Implementação:
func (r *agendamentoRepository) FindAgendamentosByStatus(status string) ([]model.AgendamentoDomainInterface, *rest_err.RestErr) {
	var agendamentos []model.AgendamentoDomain

	// Buscar agendamentos cujo status seja o informado
	if err := r.db.Where("status = ?", status).Find(&agendamentos).Error; err != nil {
		log.Printf("Erro ao buscar agendamentos com status %s: %v", status, err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar reservas", err)
	}

	// Converter para a interface
	agendamentoInterfaces := make([]model.AgendamentoDomainInterface, len(agendamentos))
	for i, a := range agendamentos {
		agendamentoInterfaces[i] = &a
	}

	return agendamentoInterfaces, nil
}
