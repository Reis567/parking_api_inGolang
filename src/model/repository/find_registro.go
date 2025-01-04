package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// FindRegistroByID busca um registro de estacionamento pelo ID
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
