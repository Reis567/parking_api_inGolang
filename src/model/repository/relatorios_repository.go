package repository

import (

	"time"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// relatoriosRepository é a estrutura que implementa a interface RelatoriosRepository
type relatoriosRepository struct {
	db *gorm.DB
}

// NewRelatoriosRepository cria uma nova instância de relatoriosRepository
func NewRelatoriosRepository(db *gorm.DB) RelatoriosRepository {
	return &relatoriosRepository{db: db}
}

// RelatoriosRepository define os métodos para consultas relacionadas a relatórios
type RelatoriosRepository interface {
	FindRegistrosPorPeriodo(inicio, fim time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	CountVagasPorStatus(status string) (int, *rest_err.RestErr)
	CountTotalVagas() (int, *rest_err.RestErr)
}
