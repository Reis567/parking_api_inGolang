package repository

import (
	"database/sql" // Importação necessária para usar *sql.DB
	"log"
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"

)

// userRepository é uma estrutura que implementa a interface UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository() UserRepository {
	return &userRepository{db: database.DB} // Usa a conexão existente em database.DB
}

// CreateUser insere um novo usuário no banco de dados
func (r *userRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	// Atribuir um novo ID ao usuário
	user.SetID(model.GenerateID())

	query := `INSERT INTO users (id, first_name, last_name, email, password, age, created_at, updated_at) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, user.GetID(), user.GetFirstName(), user.GetLastName(), user.GetEmail(),
		user.GetPassword(), user.GetAge(), user.GetCreatedAt(), user.GetUpdatedAt())
	if err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)

		// Logar o valor do JSON do usuário em caso de erro
		jsonValue, jsonErr := user.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter usuário para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados do usuário: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar usuário", err)
	}

	return user, nil
}

