package view
import (
	"meu-novo-projeto/src/controller/model/response"
	"meu-novo-projeto/src/model"
)

// Converte UserDomainInterface para UserResponse
func ConvertDomainToResponse(user model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:        user.GetID(), // Agora retorna uint
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
		Age:       user.GetAge(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}