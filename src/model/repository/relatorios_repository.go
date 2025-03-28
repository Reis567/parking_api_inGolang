package repository

import (
	"log"
	"time"

	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// relatoriosRepository é a estrutura que implementa a interface RelatoriosRepository
type relatoriosRepository struct {
	db *gorm.DB
}

// NewRelatoriosRepository cria uma nova instância de relatoriosRepository usando o banco de dados padrão
func NewRelatoriosRepository() RelatoriosRepository {
	return &relatoriosRepository{db: database.DB}
}

// NewRelatoriosRepositoryWithDB cria uma nova instância de relatoriosRepository com um banco de dados personalizado (para testes)
func NewRelatoriosRepositoryWithDB(customDB *gorm.DB) RelatoriosRepository {
	return &relatoriosRepository{db: customDB}
}

// RelatoriosRepository define os métodos para consultas relacionadas a relatórios
type RelatoriosRepository interface {
	FindRegistrosPorPeriodo(inicio, fim time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
	CountVagasPorStatus(status string) (int, *rest_err.RestErr)
	CountTotalVagas() (int, *rest_err.RestErr)
	CalcularTempoMedioPermanencia(inicio, fim time.Time) (float64, *rest_err.RestErr)
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


func (r *relatoriosRepository) CalcularTempoMedioPermanencia(inicio, fim time.Time) (float64, *rest_err.RestErr) {
	var registros []model.RegistroEstacionamentoDomain
	if err := r.db.Where("hora_entrada BETWEEN ? AND ? AND hora_saida IS NOT NULL", inicio, fim).Find(&registros).Error; err != nil {
		log.Printf("Erro ao buscar registros de permanência: %v", err)
		return 0, rest_err.NewInternalServerError("Erro ao calcular tempo médio de permanência", err)
	}

	if len(registros) == 0 {
		return 0, rest_err.NewNotFoundError("Nenhum registro encontrado no período")
	}

	var totalTempo float64
	for _, registro := range registros {
		entrada, _ := time.Parse(time.RFC3339, registro.GetHoraEntrada())
		saida, _ := time.Parse(time.RFC3339, registro.GetHoraSaida())
		totalTempo += saida.Sub(entrada).Minutes() // Converte para minutos
	}

	tempoMedio := totalTempo / float64(len(registros))
	return tempoMedio, nil
}
