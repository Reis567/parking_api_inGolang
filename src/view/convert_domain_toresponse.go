package view

import (
	"meu-novo-projeto/src/controller/model/response"
	"meu-novo-projeto/src/model"
)

// convert_domain_toresponse converte UserDomainInterface para UserResponse
func convert_domain_toresponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:         userDomain.GetID(),
		Email:      userDomain.GetEmail(),
		FirstName:  userDomain.GetFirstName(),
		LastName:   userDomain.GetLastName(),
		Age:        userDomain.GetAge(),
		CreatedAt:  userDomain.GetCreatedAt(),
		UpdatedAt:  userDomain.GetUpdatedAt(),
	}
}
