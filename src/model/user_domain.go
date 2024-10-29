package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

)

// UserDomainInterface define os métodos de acesso para UserDomain
type UserDomainInterface interface {
	GetID() string
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetCreatedAt() string
	GetUpdatedAt() string
}

// NewUserDomain é o construtor que cria uma nova instância de UserDomain e retorna UserDomainInterface
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

// Função para gerar um ID único
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Métodos Get para cada campo da estrutura UserDomain
func (ud *UserDomain) GetID() string {
	return ud.ID
}

func (ud *UserDomain) GetFirstName() string {
	return ud.FirstName
}

func (ud *UserDomain) GetLastName() string {
	return ud.LastName
}

func (ud *UserDomain) GetEmail() string {
	return ud.Email
}

func (ud *UserDomain) GetPassword() string {
	return ud.Password
}

func (ud *UserDomain) GetAge() int8 {
	return ud.Age
}

func (ud *UserDomain) GetCreatedAt() string {
	return ud.CreatedAt
}

func (ud *UserDomain) GetUpdatedAt() string {
	return ud.UpdatedAt
}

// EncryptPassword encripta a senha do usuário usando MD5
func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
