package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// FindUserByEmail busca um usuário pelo email
func (r *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	var user model.UserDomain

	// Buscar pelo campo 'email' usando GORM
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por email no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	return &user, nil
}

// FindUserByID busca um usuário pelo ID
func (r *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	var user model.UserDomain

	// Buscar pelo campo 'id' usando GORM
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	return &user, nil
}
