package model

import "meu-novo-projeto/src/configuration/rest_err"

// UserDomain representa a estrutura de um usuário no sistema
type UserDomain struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Age       int8
	CreatedAt string
	UpdatedAt string
}

// UserDomainInterface define os métodos para operações de usuário
type UserDomainInterface interface {
	CreateUser(user UserDomain) (*UserDomain, *rest_err.RestErr)
	FindUserByID(id string) (*UserDomain, *rest_err.RestErr)
	FindUserByEmail(email string) (*UserDomain, *rest_err.RestErr)
	UpdateUser(user UserDomain) (*UserDomain, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
