package repository

import (
	"log"
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


func (r *relatoriosRepository) FindRegistrosPorPeriodo(inicio, fim time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	var registros []model.RegistroEstacionamentoDomain
	if err := r.db.Where("hora_entrada BETWEEN ? AND ?", inicio, fim).Find(&registros).Error; err != nil {
		log.Printf("Erro ao buscar registros por período: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar registros", err)
	}

	// Converter para interface
	interfaces := make([]model.RegistroEstacionamentoDomainInterface, len(registros))
	for i := range registros {
		interfaces[i] = &registros[i]
	}

	return interfaces, nil
}


func (r *relatoriosRepository) CountVagasPorStatus(status string) (int, *rest_err.RestErr) {
	var count int64
	if err := r.db.Model(&model.VagaDomain{}).Where("status = ?", status).Count(&count).Error; err != nil {
		log.Printf("Erro ao contar vagas com status %s: %v", status, err)
		return 0, rest_err.NewInternalServerError("Erro ao contar vagas", err)
	}

	return int(count), nil
}


func (r *relatoriosRepository) CountTotalVagas() (int, *rest_err.RestErr) {
	var count int64
	if err := r.db.Model(&model.VagaDomain{}).Count(&count).Error; err != nil {
		log.Printf("Erro ao contar total de vagas: %v", err)
		return 0, rest_err.NewInternalServerError("Erro ao contar total de vagas", err)
	}

	return int(count), nil
}
