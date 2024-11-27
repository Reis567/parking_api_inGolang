package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/configuration/database"
	"gorm.io/gorm"
)
// userRepository é uma estrutura que implementa a interface UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository() UserRepository {
	return &userRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewUserRepositoryWithDB(customDB *gorm.DB) UserRepository {
	return &userRepository{db: customDB}
}


type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id uint) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr)
	DeleteUser(id uint) *rest_err.RestErr
}
