package repository

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/configuration/database"
	"gorm.io/gorm"
	"log"
)

// vagaRepository é a estrutura que implementa a interface VagaRepository
type vagaRepository struct {
	db *gorm.DB
}

// NewVagaRepository cria uma nova instância de vagaRepository
func NewVagaRepository() VagaRepository {
	return &vagaRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewVagaRepositoryWithDB(customDB *gorm.DB) VagaRepository {
	return &vagaRepository{db: customDB}
}

// VagaRepository interface define os métodos para gerenciar Vagas
type VagaRepository interface {
	CreateVaga(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr)
	FindVagaByID(id uint) (model.VagaDomainInterface, *rest_err.RestErr)
	FindAllVagas() ([]model.VagaDomainInterface, *rest_err.RestErr)
	UpdateVaga(vaga *model.VagaDomain) (*model.VagaDomain, *rest_err.RestErr)
	DeleteVaga(id uint) *rest_err.RestErr
}


func (r *vagaRepository) CreateVaga(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	// Tenta salvar a vaga no banco de dados usando o GORM
	if err := r.db.Create(vaga).Error; err != nil {
		log.Printf("Erro ao inserir vaga no banco de dados: %v", err)

		// Logar o valor do JSON da vaga em caso de erro
		jsonValue, jsonErr := vaga.GetJSONValue()
		if jsonErr != nil {
			log.Printf("Erro ao converter vaga para JSON: %v", jsonErr)
		} else {
			log.Printf("Dados da vaga: %s", jsonValue)
		}

		return nil, rest_err.NewInternalServerError("Erro ao criar vaga", err)
	}

	return vaga, nil
}


func (r *vagaRepository) FindVagaByID(id uint) (model.VagaDomainInterface, *rest_err.RestErr) {
	var vaga model.VagaDomain

	// Buscar pelo campo 'id' usando GORM
	if err := r.db.Where("id = ?", id).First(&vaga).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Vaga não encontrada para o ID %d", id)
			return nil, rest_err.NewNotFoundError("Vaga não encontrada")
		}
		log.Printf("Erro ao buscar vaga por ID no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vaga", err)
	}

	return &vaga, nil
}


func (r *vagaRepository) FindAllVagas() ([]model.VagaDomainInterface, *rest_err.RestErr) {
	var vagas []model.VagaDomain

	// Buscar todas as vagas usando GORM
	if err := r.db.Find(&vagas).Error; err != nil {
		log.Printf("Erro ao buscar todas as vagas no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vagas", err)
	}

	// Converter as vagas encontradas para a interface
	vagasInterfaces := make([]model.VagaDomainInterface, len(vagas))
	for i, vaga := range vagas {
		vagasInterfaces[i] = &vaga
	}

	return vagasInterfaces, nil
}


func (r *vagaRepository) UpdateVaga(vaga *model.VagaDomain) (*model.VagaDomain, *rest_err.RestErr) {
	// Verificar se a vaga existe antes de atualizar
	if err := r.db.First(&model.VagaDomain{}, vaga.ID).Error; err != nil {
		if err.Error() == "record not found" {
			log.Printf("Vaga não encontrada para atualização: ID %d", vaga.ID)
			return nil, rest_err.NewNotFoundError("Vaga não encontrada para atualização")
		}
		log.Printf("Erro ao buscar vaga para atualização: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar vaga para atualização", err)
	}

	// Atualizar os dados da vaga
	if err := r.db.Save(vaga).Error; err != nil {
		log.Printf("Erro ao atualizar vaga no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar vaga", err)
	}

	log.Printf("Vaga atualizada com sucesso: ID %d", vaga.ID)
	return vaga, nil
}


func (r *vagaRepository) DeleteVaga(id uint) *rest_err.RestErr {
	var vaga model.VagaDomain

	// Buscar a vaga para garantir que ela existe
	if err := r.db.First(&vaga, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Vaga não encontrada para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Vaga não encontrada para exclusão")
		}
		log.Printf("Erro ao buscar vaga para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar vaga para exclusão", err)
	}

	// Excluir a vaga
	if err := r.db.Delete(&vaga).Error; err != nil {
		log.Printf("Erro ao excluir vaga do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir vaga", err)
	}

	log.Printf("Vaga excluída com sucesso: ID %d", id)
	return nil
}
