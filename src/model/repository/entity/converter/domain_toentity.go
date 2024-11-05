package converter

import (
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository/entity"
)

// ConvertDomainToEntity converte um UserDomain para UserEntity
func ConvertDomainToEntity(domain model.UserDomainInterface) entity.UserEntity {
	return entity.UserEntity{
		ID:        domain.GetID(),
		FirstName: domain.GetFirstName(),
		LastName:  domain.GetLastName(),
		Email:     domain.GetEmail(),
		Password:  domain.GetPassword(),
		CreatedAt: domain.GetCreatedAt(),
		UpdatedAt: domain.GetUpdatedAt(),
		Age:       domain.GetAge(),
	}
}
