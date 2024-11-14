package service

import (
	"fmt"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"time"

	"go.uber.org/zap"
)

func (s *userDomainService) UpdateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateUser service", zap.String("journey", "Update user"))

	// Verificar se o usuário existe
	existingUser, err := s.userRepository.FindUserByID(user.GetID())
	if err != nil {
		logger.Error("Erro ao buscar usuário para atualização", zap.Error(err))
		return nil, err
	}

	// Verifica se o retorno é do tipo esperado
	userDomain, ok := existingUser.(*model.UserDomain)
	if !ok {
		logger.Error("Tipo inesperado ao buscar usuário para atualização")
		return nil, rest_err.NewInternalServerError("Erro interno no sistema", nil)
	}

	// Atualizar os campos necessários
	userDomain.FirstName = user.GetFirstName()
	userDomain.LastName = user.GetLastName()
	userDomain.Email = user.GetEmail()
	userDomain.Password = user.GetPassword() // Considere recriptar se necessário
	userDomain.Age = user.GetAge()
	userDomain.UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar as alterações no repositório
	updatedUser, updateErr := s.userRepository.UpdateUser(userDomain)
	if updateErr != nil {
		logger.Error("Erro ao atualizar usuário no banco de dados", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Usuário atualizado com sucesso", zap.String("user_id", fmt.Sprintf("%d", user.GetID())))
	return updatedUser, nil
}
