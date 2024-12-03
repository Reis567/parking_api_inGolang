package service_test

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"time"
)

// Mock do UserRepository
type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return nil, args.Get(1).(*rest_err.RestErr)
}

func (m *mockUserRepository) FindUserByID(id uint) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return nil, args.Get(1).(*rest_err.RestErr)
}

func (m *mockUserRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return nil, args.Get(1).(*rest_err.RestErr)
}

func (m *mockUserRepository) UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*model.UserDomain), nil
	}
	return nil, args.Get(1).(*rest_err.RestErr)
}

func (m *mockUserRepository) DeleteUser(id uint) *rest_err.RestErr {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*rest_err.RestErr)
	}
	return nil
}

// Teste para FindUserByIDService
func TestFindUserByIDService(t *testing.T) {
	// Configurar mock do repositório
	mockRepo := new(mockUserRepository)
	userService := service.NewUserDomainService(mockRepo)

	// Usuário de teste
	testUser := &model.UserDomain{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Age:       30,
	}

	// Cenário de sucesso
	mockRepo.On("FindUserByID", uint(1)).Return(testUser, nil)
	user, err := userService.FindUserByIDService(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John", user.GetFirstName())
	assert.Equal(t, "Doe", user.GetLastName())
	assert.Equal(t, "john.doe@example.com", user.GetEmail())
	mockRepo.AssertExpectations(t)

	// Cenário de falha
	mockRepo.On("FindUserByID", uint(999)).Return(nil, rest_err.NewNotFoundError("Usuário não encontrado"))
	user, err = userService.FindUserByIDService(999)

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "Usuário não encontrado", err.Message)
	mockRepo.AssertExpectations(t)
}

// Teste para FindUserByEmailService
func TestFindUserByEmailService(t *testing.T) {
	// Configurar mock do repositório
	mockRepo := new(mockUserRepository)
	userService := service.NewUserDomainService(mockRepo)

	// Usuário de teste
	testUser := &model.UserDomain{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Password:  "password123",
		Age:       25,
	}

	// Cenário de sucesso
	mockRepo.On("FindUserByEmail", "jane.smith@example.com").Return(testUser, nil)
	user, err := userService.FindUserByEmailService("jane.smith@example.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Jane", user.GetFirstName())
	assert.Equal(t, "Smith", user.GetLastName())
	assert.Equal(t, "jane.smith@example.com", user.GetEmail())
	mockRepo.AssertExpectations(t)

	// Cenário de falha
	mockRepo.On("FindUserByEmail", "ghost@example.com").Return(nil, rest_err.NewNotFoundError("Usuário não encontrado"))
	user, err = userService.FindUserByEmailService("ghost@example.com")

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "Usuário não encontrado", err.Message)
	mockRepo.AssertExpectations(t)
}


func TestCreateUserService(t *testing.T) {
	// Configurar mock do repositório
	mockRepo := new(mockUserRepository)
	userService := service.NewUserDomainService(mockRepo)

	// Usuário de teste
	testUser := &model.UserDomain{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Age:       30,
	}

	// Cenário de sucesso
	mockRepo.On("CreateUser", mock.Anything).Return(func(user model.UserDomainInterface) model.UserDomainInterface {
		user.(*model.UserDomain).ID = 1
		user.(*model.UserDomain).CreatedAt = time.Now().Format(time.RFC3339)
		user.(*model.UserDomain).UpdatedAt = time.Now().Format(time.RFC3339)
		return user
	}, nil)
	mockRepo.On("FindUserByEmail", "john.doe@example.com").Return(nil, nil)

	// Chamar o serviço
	createdUser, err := userService.CreateUserService(testUser)

	// Verificações
	assert.Nil(t, err, "Erro deve ser nulo")
	assert.NotNil(t, createdUser, "Usuário criado não deve ser nulo")
	assert.Equal(t, "John", createdUser.GetFirstName(), "Primeiro nome deve ser igual")
	assert.Equal(t, "Doe", createdUser.GetLastName(), "Último nome deve ser igual")
	assert.Equal(t, "john.doe@example.com", createdUser.GetEmail(), "Email deve ser igual")
	mockRepo.AssertExpectations(t)

	// Cenário de falha - Email já registrado
	mockRepo.On("FindUserByEmail", "john.doe@example.com").Return(testUser, nil)

	createdUser, err = userService.CreateUserService(testUser)

	assert.NotNil(t, err, "Erro deve ser retornado")
	assert.Nil(t, createdUser, "Usuário não deve ser criado")
	assert.Equal(t, "Email ja registrado no sistema", err.Message, "Mensagem de erro deve ser 'Email ja registrado no sistema'")
	mockRepo.AssertExpectations(t)

	// Cenário de falha - Erro ao salvar no repositório
	mockRepo.On("FindUserByEmail", "unique@example.com").Return(nil, nil)
	mockRepo.On("CreateUser", mock.Anything).Return(nil, rest_err.NewInternalServerError("Erro ao criar usuário", nil))

	testUser.Email = "unique@example.com"
	createdUser, err = userService.CreateUserService(testUser)

	assert.NotNil(t, err, "Erro deve ser retornado")
	assert.Nil(t, createdUser, "Usuário não deve ser criado")
	assert.Equal(t, "Erro ao criar usuário", err.Message, "Mensagem de erro deve ser 'Erro ao criar usuário'")
	mockRepo.AssertExpectations(t)
}