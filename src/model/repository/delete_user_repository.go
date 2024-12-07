package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

func (r *userRepository) DeleteUser(id uint) *rest_err.RestErr {
	// Buscar o usuário para garantir que ele existe antes de tentar excluir
	var user model.UserDomain
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Usuário não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Usuário não encontrado para exclusão")
		}
		log.Printf("Erro ao buscar usuário para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar usuário para exclusão", err)
	}

	// Excluir o usuário
	if err := r.db.Delete(&user).Error; err != nil {
		log.Printf("Erro ao excluir usuário do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir usuário", err)
	}

	log.Printf("Usuário excluído com sucesso: ID %d", id)
	return nil
}
