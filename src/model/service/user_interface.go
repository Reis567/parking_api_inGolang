package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
)

// NewUserDomainService cria uma instância de userDomainService
func NewUserDomainService( userRepository repository.UserRepository ) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct{
	userRepository repository.UserRepository
}

// Interface do serviço de domínio do usuário
type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(id uint) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(id string) *rest_err.RestErr
	LoginUserService(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
}
