package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id uint) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr)
}
