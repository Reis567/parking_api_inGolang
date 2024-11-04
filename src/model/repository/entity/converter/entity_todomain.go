package entity

import (
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository/entity"
)

// ConvertDomainToEntity converte um UserDomain para UserEntity
func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.FirstName, entity.LastName, entity.Email, entity.Password, entity.Age).(*model.UserDomain)
	domain.SetID(entity.ID)
	domain.CreatedAt = entity.CreatedAt
	domain.UpdatedAt = entity.UpdatedAt
	return domain
}