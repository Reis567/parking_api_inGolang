package view

import (
	"meu-novo-projeto/src/controller/model/response"
	"meu-novo-projeto/src/model"
)

func convert_domain_toresponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{ID: "",}
}