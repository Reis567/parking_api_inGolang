package repository_test

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_DeleteUser(t *testing.T) {
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

	// Testar exclusão de usuário existente
	deleteErr := repo.DeleteUser(testUser.ID)
	assert.Nil(t, deleteErr, "Erro ao excluir usuário deve ser nulo")

	// Verificar se o usuário foi realmente excluído
	var deletedUser model.UserDomain
	findErr := testDB.First(&deletedUser, testUser.ID).Error
	assert.NotNil(t, findErr, "Usuário excluído não deve ser encontrado")
	assert.Equal(t, "record not found", findErr.Error(), "Mensagem de erro deve indicar que o registro não foi encontrado")

	// Testar exclusão de um usuário inexistente
	deleteErr = repo.DeleteUser(999) // ID inexistente
	assert.NotNil(t, deleteErr, "Erro esperado ao excluir usuário inexistente")
	assert.Equal(t, rest_err.NewNotFoundError("Usuário não encontrado para exclusão").Message, deleteErr.Message, "Mensagem de erro deve ser 'Usuário não encontrado para exclusão'")
}
