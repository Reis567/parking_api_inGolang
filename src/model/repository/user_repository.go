package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/configuration/database"
	"gorm.io/gorm"
	"log"
)
// userRepository é uma estrutura que implementa a interface UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository() UserRepository {
	return &userRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewUserRepositoryWithDB(customDB *gorm.DB) UserRepository {
	return &userRepository{db: customDB}
}


type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id uint) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr)
	DeleteUser(id uint) *rest_err.RestErr
}


func (r *userRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	// Tenta salvar o usuário no banco de dados usando o GORM
	if err := r.db.Create(user).Error; err != nil {
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


func (r *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	var user model.UserDomain

	// Buscar pelo campo 'email' usando GORM
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por email no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	return &user, nil
}

// FindUserByID busca um usuário pelo ID
func (r *userRepository) FindUserByID(id uint) (model.UserDomainInterface, *rest_err.RestErr) {
	var user model.UserDomain

	// Buscar pelo campo 'id' usando GORM
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário", err)
	}

	return &user, nil
}


func (r *userRepository) UpdateUser(user *model.UserDomain) (*model.UserDomain, *rest_err.RestErr) {
	// Verificar se o usuário existe antes de atualizar
	if err := r.db.First(&model.UserDomain{}, user.ID).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Usuário não encontrado para atualização: %v", user.ID)
			return nil, rest_err.NewNotFoundError("Usuário não encontrado para atualização")
		}
		log.Printf("Erro ao buscar usuário para atualização: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário para atualização", err)
	}

	// Atualizar os dados do usuário
	if err := r.db.Save(user).Error; err != nil {
		log.Printf("Erro ao atualizar usuário no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar usuário", err)
	}

	return user, nil
}
func (r *userRepository) DeleteUser(id uint) *rest_err.RestErr {
	// Buscar o usuário para garantir que ele existe antes de tentar excluir
	var user model.UserDomain
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Usuário não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Usuário não encontrado para exclusão")
		}
		log.Printf("Erro ao buscar usuário para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar usuário para exclusão", err)
	}

	// Excluir o usuário
	if err := r.db.Delete(&user).Error; err != nil {
		log.Printf("Erro ao excluir usuário do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir usuário", err)
	}

	log.Printf("Usuário excluído com sucesso: ID %d", id)
	return nil
}


func (r *userRepository) GetCurrentUser(userID uint) (model.UserDomainInterface, *rest_err.RestErr) {
	var user model.UserDomain

	if err := r.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Printf("Erro ao buscar usuário atual: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário atual", err)
	}

	return &user, nil
}

func (r *userRepository) GetUserParkingHistory(userID uint) ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	var registros []model.RegistroEstacionamentoDomain

	if err := r.db.Where("user_id = ?", userID).Find(&registros).Error; err != nil {
		log.Printf("Erro ao buscar histórico de estacionamento do usuário: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar histórico de estacionamento", err)
	}

	registrosInterface := make([]model.RegistroEstacionamentoDomainInterface, len(registros))
	for i, registro := range registros {
		registrosInterface[i] = &registro
	}
	return registrosInterface, nil
}
