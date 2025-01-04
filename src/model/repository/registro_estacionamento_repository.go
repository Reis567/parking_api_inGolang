package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// registroEstacionamentoRepository é a estrutura que implementa a interface RegistroEstacionamentoRepository
type registroEstacionamentoRepository struct {
	db *gorm.DB
}

// NewRegistroEstacionamentoRepository cria uma nova instância de registroEstacionamentoRepository
func NewRegistroEstacionamentoRepository() RegistroEstacionamentoRepository {
	return &registroEstacionamentoRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewRegistroEstacionamentoRepositoryWithDB(customDB *gorm.DB) RegistroEstacionamentoRepository {
	return &registroEstacionamentoRepository{db: customDB}
}

// RegistroEstacionamentoRepository interface define os métodos para gerenciar Registros de Estacionamento
type RegistroEstacionamentoRepository interface {
	CreateRegistro(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindRegistroByID(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	FindAllRegistros() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	UpdateRegistro(registro *model.RegistroEstacionamentoDomain) (*model.RegistroEstacionamentoDomain, *rest_err.RestErr)
	DeleteRegistro(id uint) *rest_err.RestErr
}
