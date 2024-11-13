package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// UpdateUser atualiza um usuário existente no banco de dados
func (r *userRepository) UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr) {
	if err := r.db.Save(user).Error; err != nil {
		log.Printf("Erro ao atualizar usuário no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar usuário", err)
	}
	return user, nil
}
