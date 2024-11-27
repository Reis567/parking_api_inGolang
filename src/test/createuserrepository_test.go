package repository_test

import (
	"meu-novo-projeto/src/model/repository"
	"meu-novo-projeto/src/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)


func setupTestDB() *gorm.DB {
	// Cria um banco de dados em memória usando SQLite
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados de teste")
	}

	// Faz o migrate do modelo de UserDomain para o banco
	err = db.AutoMigrate(&model.UserDomain{})
	if err != nil {
		panic("Falha ao migrar o banco de dados de teste")
	}

	return db
}

func TestUserRepository_CreateUser(t *testing.T) {
	// Configuração do banco de dados de teste
	testDB := setupTestDB()
	repo := repository.NewUserRepositoryWithDB(testDB) // Use um construtor que permita passar o DB

	// Dados do usuário de teste
	user := &model.UserDomain{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Age:       30,
	}

	// Testa a criação do usuário
	createdUser, err := repo.CreateUser(user)

	// Verificações
	assert.Nil(t, err, "Erro ao criar usuário deve ser nulo")
	assert.NotNil(t, createdUser, "Usuário criado não deve ser nulo")
	assert.Equal(t, user.Email, createdUser.GetEmail(), "Email do usuário deve ser igual")
	assert.Equal(t, user.FirstName, createdUser.GetFirstName(), "Primeiro nome deve ser igual")
	assert.Equal(t, user.LastName, createdUser.GetLastName(), "Último nome deve ser igual")
}

