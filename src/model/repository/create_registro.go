package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

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
