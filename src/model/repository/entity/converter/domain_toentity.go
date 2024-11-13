package converter

import (
	"meu-novo-projeto/src/model"
)

// ConvertDomainToEntity converte um UserDomain para ser usado como entidade no GORM
func ConvertDomainToEntity(domain model.UserDomainInterface) *model.UserDomain {
	return &model.UserDomain{
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
