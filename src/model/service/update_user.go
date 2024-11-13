package service

import (
	"fmt"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"time"

	"go.uber.org/zap"
)

// UpdateUser atualiza um usuário existente
func (s *userDomainService) UpdateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateUser service", zap.String("journey", "Update user"))

	// Verificar se o usuário existe
	existingUser, err := s.userRepository.FindUserByID(user.GetID())
	if err != nil {
		logger.Error("Erro ao buscar usuário para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos necessários
	existingUser.(*model.UserDomain).FirstName = user.GetFirstName()
	existingUser.(*model.UserDomain).LastName = user.GetLastName()
	existingUser.(*model.UserDomain).Email = user.GetEmail()
	existingUser.(*model.UserDomain).Password = user.GetPassword() // Considere recriptar se necessário
	existingUser.(*model.UserDomain).Age = user.GetAge()
	existingUser.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar as alterações no repositório
	updatedUser, updateErr := s.userRepository.UpdateUser(existingUser.(*model.UserDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar usuário no banco de dados", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Usuário atualizado com sucesso", zap.String("user_id", fmt.Sprintf("%d", user.GetID())))
	return updatedUser, nil
}
