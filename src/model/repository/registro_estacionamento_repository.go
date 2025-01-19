package repository

import (
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
	"log"
	"time"
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
	FindRegistrosPorPeriodo(inicio, fim time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr)
}



// CreateRegistro insere um novo registro de estacionamento no banco de dados
func (r *registroEstacionamentoRepository) CreateRegistro(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(registro).Error; err != nil {
		log.Printf("Erro ao inserir registro no banco de dados: %v", err)

		// Logar o valor do JSON do registro em caso de erro
		jsonValue, jsonErr := registro.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter registro para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados do registro: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar registro", err)
	}
	return registro, nil
}

func (r *registroEstacionamentoRepository) FindRegistroByID(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	var registro model.RegistroEstacionamentoDomain

	if err := r.db.First(&registro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Registro não encontrado: ID %d", id)
			return nil, rest_err.NewNotFoundError("Registro não encontrado")
		}
		log.Printf("Erro ao buscar registro no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar registro", err)
	}
	return &registro, nil
}

func (r *registroEstacionamentoRepository) FindAllRegistros() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	var registros []model.RegistroEstacionamentoDomain

	if err := r.db.Find(&registros).Error; err != nil {
		log.Printf("Erro ao buscar todos os registros: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar registros", err)
	}

	registrosInterface := make([]model.RegistroEstacionamentoDomainInterface, len(registros))
	for i, registro := range registros {
		registrosInterface[i] = &registro
	}
	return registrosInterface, nil
}

func (r *registroEstacionamentoRepository) UpdateRegistro(registro *model.RegistroEstacionamentoDomain) (*model.RegistroEstacionamentoDomain, *rest_err.RestErr) {
	if err := r.db.Save(registro).Error; err != nil {
		log.Printf("Erro ao atualizar registro no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar registro", err)
	}
	return registro, nil
}


func (r *registroEstacionamentoRepository) DeleteRegistro(id uint) *rest_err.RestErr {
	var registro model.RegistroEstacionamentoDomain

	if err := r.db.First(&registro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Registro não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Registro não encontrado")
		}
		log.Printf("Erro ao buscar registro para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar registro para exclusão", err)
	}

	if err := r.db.Delete(&registro).Error; err != nil {
		log.Printf("Erro ao excluir registro: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir registro", err)
	}
	log.Printf("Registro excluído com sucesso: ID %d", id)
	return nil
}


func (r *registroEstacionamentoRepository) FindRegistrosPorPeriodo(inicio, fim time.Time) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
    var registros []model.RegistroEstacionamentoDomain

    if err := r.db.Where("hora_entrada BETWEEN ? AND ?", inicio, fim).Find(&registros).Error; err != nil {
        log.Printf("Erro ao buscar registros no período: %v", err)
        return nil, rest_err.NewInternalServerError("Erro ao buscar registros por período", err)
    }

    registrosInterface := make([]model.RegistroEstacionamentoDomainInterface, len(registros))
    for i, registro := range registros {
        registrosInterface[i] = &registro
    }
    return registrosInterface, nil
}
