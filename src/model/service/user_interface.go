package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"go.uber.org/zap"
	"time"
	"meu-novo-projeto/src/configuration/logger"
	"strconv"
	"fmt"
)

// NewUserDomainService cria uma instância de userDomainService
func NewUserDomainService( userRepository repository.UserRepository ) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct{
	userRepository repository.UserRepository
}

// Interface do serviço de domínio do usuário
type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(id uint) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(id string) *rest_err.RestErr
	LoginUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
func (s *userDomainService) CreateUserService(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser model", zap.String("journey", "Create user"))

	// Atribuir datas (ID será gerenciado automaticamente pelo GORM)
	user.(*model.UserDomain).CreatedAt = time.Now().Format(time.RFC3339)
	user.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Encriptar a senha
	user.(*model.UserDomain).EncryptPassword()

	userComEmail, _ := s.FindUserByEmailService(user.GetEmail())
	if userComEmail!=nil{
		return nil ,rest_err.NewBadRequestError("Email ja registrado no sistema")
	}

	// Salvar o usuário no repositório
	userDomainRepository, err := s.userRepository.CreateUser(user)
	if err != nil {
		logger.Error("Erro ao salvar usuário no banco de dados", zap.Error(err))
		return nil, err
	}

	// Log de sucesso
	logger.Info("Usuário criado com sucesso", zap.Uint("user_id", userDomainRepository.GetID()))
	return userDomainRepository, nil
}



// FindUserByID busca um usuário pelo ID
func (s *userDomainService) FindUserByIDService(id uint) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID service", zap.String("journey", "Find user by ID"))

	// Consulta o repositório para buscar o usuário pelo ID
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		logger.Error("Erro ao buscar usuário no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Usuário encontrado com sucesso", zap.Uint("user_id", id))
	return user, nil
}

// FindUserByEmail busca um usuário pelo email
func (s *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "Find user by email"))

	// Consulta o repositório para buscar o usuário pelo email
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Erro ao buscar usuário no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Usuário encontrado com sucesso", zap.String("user_email", email))
	return user, nil
}



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


func (s *userDomainService) DeleteUserService(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "Delete user"))

	// Converte o ID para uint
	idUint, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("ID inválido para exclusão", zap.String("user_id", id), zap.Error(err))
		return rest_err.NewBadRequestError("ID inválido. Deve ser um número inteiro.")
	}

	// Chama o repositório para excluir o usuário
	deleteErr := s.userRepository.DeleteUser(uint(idUint))
	if deleteErr != nil {
		logger.Error("Erro ao excluir usuário no repositório", zap.Error(deleteErr))
		return deleteErr
	}

	logger.Info("Usuário excluído com sucesso", zap.String("user_id", id))
	return nil
}

func (s *userDomainService) GetUserParkingHistoryService(userID uint) ([]model.ParkingHistory, *rest_err.RestErr) {
    logger.Info("Init GetUserParkingHistoryService", zap.Uint("userID", userID))

    // Recupera o histórico de estacionamento do usuário
    parkingHistory, dbErr := s.userRepository.GetUserParkingHistory(userID)
    if dbErr != nil {
        logger.Error("Erro ao recuperar histórico de estacionamento", zap.Error(dbErr))
        return nil, dbErr
    }

    logger.Info("Histórico de estacionamento recuperado com sucesso", zap.Uint("userID", userID))
    return parkingHistory, nil
}

func (s *userDomainService) GetUserVehiclesService(userID string) ([]model.VehicleDomainInterface, *rest_err.RestErr) {
    logger.Info("Init GetUserVehiclesService", zap.String("userID", userID))

    // Convertendo o ID para uint
    idUint, err := strconv.Atoi(userID)
    if err != nil {
        logger.Error("ID inválido para recuperação de veículos", zap.String("userID", userID), zap.Error(err))
        return nil, rest_err.NewBadRequestError("ID inválido. Deve ser um número inteiro.")
    }

    // Recupera os veículos do usuário
    vehicles, dbErr := s.userRepository.GetUserVehicles(uint(idUint))
    if dbErr != nil {
        logger.Error("Erro ao recuperar veículos do usuário", zap.Error(dbErr))
        return nil, dbErr
    }

    // Converte []model.VehicleDomain para []model.VehicleDomainInterface
    var vehicleInterfaces []model.VehicleDomainInterface
    for _, v := range vehicles {
        vehicleInterfaces = append(vehicleInterfaces, v)
    }

    logger.Info("Veículos do usuário recuperados com sucesso", zap.String("userID", userID))
    return vehicleInterfaces, nil
}
