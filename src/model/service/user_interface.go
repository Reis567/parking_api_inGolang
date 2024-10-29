package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

// NewUserDomainService cria uma instância de userDomainService
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

// Interface do serviço de domínio do usuário
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
