package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// CreateVaga insere uma nova vaga no banco de dados
func (r *vagaRepository) CreateVaga(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	// Tenta salvar a vaga no banco de dados usando o GORM
	if err := r.db.Create(vaga).Error; err != nil {
		log.Printf("Erro ao inserir vaga no banco de dados: %v", err)

		// Logar o valor do JSON da vaga em caso de erro
		jsonValue, jsonErr := vaga.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter vaga para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados da vaga: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar vaga", err)
	}

	return vaga, nil
}
