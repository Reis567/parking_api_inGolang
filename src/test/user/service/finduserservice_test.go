package service_test

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do UserRepository
type mockUserRepository struct {
	mock.Mock
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

func TestFindUserByIDService(t *testing.T) {
	// Configurar mock do repositório
	mockRepo := new(mockUserRepository)
	service := service.NewUserDomainService(mockRepo)

	// Usuário de teste
	testUser := &model.UserDomain{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Age:       30,
	}

	// Configurar mock para sucesso
	mockRepo.On("FindUserByID", uint(1)).Return(testUser, nil)

	// Chamar o serviço
	user, err := service.FindUserByIDService(1)

	// Verificações
	assert.Nil(t, err, "Erro deve ser nulo")
	assert.NotNil(t, user, "Usuário retornado não deve ser nulo")
	assert.Equal(t, testUser.FirstName, user.GetFirstName(), "Primeiro nome deve ser igual")
	assert.Equal(t, testUser.LastName, user.GetLastName(), "Último nome deve ser igual")
	assert.Equal(t, testUser.Email, user.GetEmail(), "Email deve ser igual")
	mockRepo.AssertExpectations(t)

	// Configurar mock para falha
	mockRepo.On("FindUserByID", uint(999)).Return(nil, rest_err.NewNotFoundError("Usuário não encontrado"))

	// Chamar o serviço com ID inexistente
	user, err = service.FindUserByIDService(999)

	// Verificações
	assert.NotNil(t, err, "Erro deve ser retornado")
	assert.Nil(t, user, "Usuário deve ser nulo")
	assert.Equal(t, "Usuário não encontrado", err.Message, "Mensagem de erro deve ser 'Usuário não encontrado'")
	mockRepo.AssertExpectations(t)
}

func TestFindUserByEmailService(t *testing.T) {
	// Configurar mock do repositório
	mockRepo := new(mockUserRepository)
	service := service.NewUserDomainService(mockRepo)

	// Usuário de teste
	testUser := &model.UserDomain{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Password:  "password123",
		Age:       25,
	}

	// Configurar mock para sucesso
	mockRepo.On("FindUserByEmail", "jane.smith@example.com").Return(testUser, nil)

	// Chamar o serviço
	user, err := service.FindUserByEmailService("jane.smith@example.com")

	// Verificações
	assert.Nil(t, err, "Erro deve ser nulo")
	assert.NotNil(t, user, "Usuário retornado não deve ser nulo")
	assert.Equal(t, testUser.FirstName, user.GetFirstName(), "Primeiro nome deve ser igual")
	assert.Equal(t, testUser.LastName, user.GetLastName(), "Último nome deve ser igual")
	assert.Equal(t, testUser.Email, user.GetEmail(), "Email deve ser igual")
	mockRepo.AssertExpectations(t)

	// Configurar mock para falha
	mockRepo.On("FindUserByEmail", "ghost@example.com").Return(nil, rest_err.NewNotFoundError("Usuário não encontrado"))

	// Chamar o serviço com email inexistente
	user, err = service.FindUserByEmailService("ghost@example.com")

	// Verificações
	assert.NotNil(t, err, "Erro deve ser retornado")
	assert.Nil(t, user, "Usuário deve ser nulo")
	assert.Equal(t, "Usuário não encontrado", err.Message, "Mensagem de erro deve ser 'Usuário não encontrado'")
	mockRepo.AssertExpectations(t)
}
