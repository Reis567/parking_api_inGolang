package converter

import (
	"meu-novo-projeto/src/model"
)

// ConvertEntityToDomain converte um UserDomain (struct do GORM) para a interface UserDomainInterface
func ConvertEntityToDomain(entity *model.UserDomain) model.UserDomainInterface {
	return entity
}
