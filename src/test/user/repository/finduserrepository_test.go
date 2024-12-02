package repository_test

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_FindUserByID(t *testing.T) {
	// Configuração do banco de dados de teste
	testDB := setupTestDB()
	repo := repository.NewUserRepositoryWithDB(testDB)

	// Inserir um usuário de teste no banco
	testUser := &model.UserDomain{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@example.com",
		Password:  "password123",
		Age:       25,
	}
	testDB.Create(testUser)

	// Testar buscar pelo ID existente
	foundUser, err := repo.FindUserByID(testUser.ID)

	// Verificar se o usuário foi encontrado corretamente
	assert.Nil(t, err, "Erro ao buscar usuário deve ser nulo")
	assert.NotNil(t, foundUser, "Usuário encontrado não deve ser nulo")
	assert.Equal(t, testUser.ID, foundUser.GetID(), "ID do usuário deve ser igual")
	assert.Equal(t, testUser.Email, foundUser.GetEmail(), "Email do usuário deve ser igual")

	// Testar buscar por um ID inexistente
	_, err = repo.FindUserByID(999)
	assert.NotNil(t, err, "Erro esperado ao buscar ID inexistente")
	assert.Equal(t, rest_err.NewNotFoundError("Usuário não encontrado").Message, err.Message, "Mensagem de erro deve ser 'Usuário não encontrado'")
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	// Configuração do banco de dados de teste
	testDB := setupTestDB()
	repo := repository.NewUserRepositoryWithDB(testDB)

	// Inserir um usuário de teste no banco
	testUser := &model.UserDomain{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Age:       30,
	}
	testDB.Create(testUser)

	// Testar buscar pelo email existente
	foundUser, err := repo.FindUserByEmail(testUser.Email)

	// Verificar se o usuário foi encontrado corretamente
	assert.Nil(t, err, "Erro ao buscar usuário deve ser nulo")
	assert.NotNil(t, foundUser, "Usuário encontrado não deve ser nulo")
	assert.Equal(t, testUser.Email, foundUser.GetEmail(), "Email do usuário deve ser igual")

	// Testar buscar por um email inexistente
	_, err = repo.FindUserByEmail("nonexistent@example.com")
	assert.NotNil(t, err, "Erro esperado ao buscar email inexistente")
	assert.Equal(t, rest_err.NewNotFoundError("Usuário não encontrado").Message, err.Message, "Mensagem de erro deve ser 'Usuário não encontrado'")
}
