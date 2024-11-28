package repository_test

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_UpdateUser(t *testing.T) {
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

	// Atualizar os dados do usuário
	updatedData := &model.UserDomain{
		ID:        testUser.ID,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Password:  "newpassword123",
		Age:       35,
	}
	updatedUser, err := repo.UpdateUser(updatedData)

	// Verificações para usuário existente
	assert.Nil(t, err, "Erro ao atualizar usuário deve ser nulo")
	assert.NotNil(t, updatedUser, "Usuário atualizado não deve ser nulo")
	assert.Equal(t, updatedData.FirstName, updatedUser.FirstName, "Primeiro nome deve ser atualizado")
	assert.Equal(t, updatedData.LastName, updatedUser.LastName, "Último nome deve ser atualizado")
	assert.Equal(t, updatedData.Email, updatedUser.Email, "Email deve ser atualizado")
	assert.Equal(t, updatedData.Age, updatedUser.Age, "Idade deve ser atualizada")

	// Testar atualização de um usuário inexistente
	nonExistentUser := &model.UserDomain{
		ID:        999, // ID inexistente
		FirstName: "Ghost",
		LastName:  "User",
		Email:     "ghost.user@example.com",
		Password:  "ghostpassword",
		Age:       40,
	}
	updatedUser, err = repo.UpdateUser(nonExistentUser)

	// Verificações para usuário inexistente
	assert.Nil(t, updatedUser, "Usuário atualizado deve ser nulo para ID inexistente")
	assert.NotNil(t, err, "Erro esperado ao atualizar usuário inexistente")
	assert.Equal(t, rest_err.NewNotFoundError("Usuário não encontrado para atualização").Message, err.Message, "Mensagem de erro deve ser 'Usuário não encontrado para atualização'")
}
