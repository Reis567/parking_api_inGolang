package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"


)

// CreateUser insere um novo usuário no banco de dados
func (r *userRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	// Tenta salvar o usuário no banco de dados usando o GORM
	if err := r.db.Create(user).Error; err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)

		// Logar o valor do JSON do usuário em caso de erro
		jsonValue, jsonErr := user.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter usuário para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados do usuário: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar usuário", err)
	}

	return user, nil
}
