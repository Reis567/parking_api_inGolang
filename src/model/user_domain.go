package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"


)

type UserDomainInterface interface {	
	GetID() uint
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	CheckPassword(password string) bool
}

func NewUserDomain(firstName, lastName, email, password string, age int8) UserDomainInterface {
	user := &UserDomain{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Age:       age,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	user.EncryptPassword() // Encripta a senha automaticamente
	return user
}

// UserDomain representa a estrutura de um usuário no sistema
type UserDomain struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"-"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Age       int8           `json:"age"`
}

// Função para gerar um ID único
func GenerateID() uint {
	return uint(time.Now().UnixNano()) // ajustado para uint
}

// Métodos Get para cada campo da estrutura UserDomain
func (ud *UserDomain) GetID() uint {
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

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
func (ud *UserDomain) CheckPassword(password string) bool {
	hash := md5.New() // Use o mesmo método de criptografia usado ao criar a senha
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)) == ud.Password
}


func (ud *UserDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(ud)
	if err != nil {
		return "", fmt.Errorf("erro ao converter UserDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

func (ud *UserDomain) SetID(id uint) {
	ud.ID = id
}
