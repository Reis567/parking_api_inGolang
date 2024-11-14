package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

func (r *userRepository) UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr) {
	// Verificar se o usuário existe antes de atualizar
	if err := r.db.First(&model.UserDomain{}, user.ID).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Usuário não encontrado para atualização: %v", user.ID)
			return nil, rest_err.NewNotFoundError("Usuário não encontrado para atualização")
		}
		log.Printf("Erro ao buscar usuário para atualização: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário para atualização", err)
	}

	// Atualizar os dados do usuário
	if err := r.db.Save(user).Error; err != nil {
		log.Printf("Erro ao atualizar usuário no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar usuário", err)
	}

	return user, nil
}
