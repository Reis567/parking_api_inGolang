package repository

import (
	"database/sql"
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository/entity"
	"meu-novo-projeto/src/model/repository/entity/converter"
)

// FindUserByEmail busca um usuário pelo email
func (r *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, first_name, last_name, email, password, age, created_at, updated_at FROM users WHERE email = ?`

	row := r.db.QueryRow(query, email)

	var userEntity entity.UserEntity
	err := row.Scan(
		&userEntity.ID,
		&userEntity.FirstName,
		&userEntity.LastName,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Age,
		&userEntity.CreatedAt,
		&userEntity.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por email no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	userDomain := converter.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}

// FindUserByID busca um usuário pelo ID
func (r *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, first_name, last_name, email, password, age, created_at, updated_at FROM users WHERE id = ?`

	row := r.db.QueryRow(query, id)

	var userEntity entity.UserEntity
	err := row.Scan(
		&userEntity.ID,
		&userEntity.FirstName,
		&userEntity.LastName,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Age,
		&userEntity.CreatedAt,
		&userEntity.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	userDomain := converter.ConvertEntityToDomain(userEntity)
	return userDomain, nil
}
