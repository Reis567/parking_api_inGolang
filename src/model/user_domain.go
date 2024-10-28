package model

import (
	"crypto/md5"
	"encoding/hex"

	"meu-novo-projeto/src/configuration/rest_err"
	"time"
)

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

func NewUserDomain(firstName, lastName, email, password string, age int8) UserDomainInterface {
	user := &UserDomain{
		ID:        generateID(),                            // Gerar um ID único
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Age:       age,
		CreatedAt: time.Now().Format(time.RFC3339),         // Definir data de criação
		UpdatedAt: time.Now().Format(time.RFC3339),         // Definir data de atualização
	}
	user.EncryptPassword() // Encripta a senha automaticamente
	return user
}

func (ud *UserDomain) EncryptPassword() {
	hash:= md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password=hex.EncodeToString(hash.Sum(nil))
}


type UserDomainInterface interface {
	CreateUser(user UserDomain) (*UserDomain, *rest_err.RestErr)
	FindUserByID(id string) (*UserDomain, *rest_err.RestErr)
	FindUserByEmail(email string) (*UserDomain, *rest_err.RestErr)
	UpdateUser(user UserDomain) (*UserDomain, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
