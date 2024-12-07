package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// FindVagaByID busca uma vaga pelo ID
func (r *vagaRepository) FindVagaByID(id uint) (model.VagaDomainInterface, *rest_err.RestErr) {
	var vaga model.VagaDomain

	// Buscar pelo campo 'id' usando GORM
	if err := r.db.Where("id = ?", id).First(&vaga).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Vaga não encontrada para o ID %d", id)
			return nil, rest_err.NewNotFoundError("Vaga não encontrada")
		}
		log.Printf("Erro ao buscar vaga por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vaga", err)
	}

	return &vaga, nil
}


func (r *vagaRepository) FindAllVagas() ([]model.VagaDomainInterface, *rest_err.RestErr) {
	var vagas []model.VagaDomain

	// Buscar todas as vagas usando GORM
	if err := r.db.Find(&vagas).Error; err != nil {
		log.Printf("Erro ao buscar todas as vagas no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vagas", err)
	}

	// Converter as vagas encontradas para a interface
	vagasInterfaces := make([]model.VagaDomainInterface, len(vagas))
	for i, vaga := range vagas {
		vagasInterfaces[i] = &vaga
	}

	return vagasInterfaces, nil
}