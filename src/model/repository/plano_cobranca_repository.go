package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// billingPlanRepository é a estrutura que implementa a interface BillingPlanRepository
type billingPlanRepository struct {
	db *gorm.DB
}

// NewBillingPlanRepository cria uma nova instância de billingPlanRepository
func NewBillingPlanRepository() BillingPlanRepository {
	return &billingPlanRepository{db: database.DB} // Usa a conexão existente em database.DB
}

func NewBillingPlanRepositoryWithDB(customDB *gorm.DB) BillingPlanRepository {
	return &billingPlanRepository{db: customDB}
}

// BillingPlanRepository interface define os métodos para gerenciar Planos de Cobrança
type BillingPlanRepository interface {
	CreateBillingPlan(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr)
	FindBillingPlanByID(id uint) (model.BillingPlanDomainInterface, *rest_err.RestErr)
	FindAllBillingPlans() ([]model.BillingPlanDomainInterface, *rest_err.RestErr)
	UpdateBillingPlan(plan *model.BillingPlanDomain) (*model.BillingPlanDomain, *rest_err.RestErr)
	DeleteBillingPlan(id uint) *rest_err.RestErr
}

// CreateBillingPlan insere um novo plano de cobrança no banco de dados
func (r *billingPlanRepository) CreateBillingPlan(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(plan).Error; err != nil {
		log.Printf("Erro ao inserir plano de cobrança no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao criar plano de cobrança", err)
	}
	return plan, nil
}

// FindBillingPlanByID busca um plano de cobrança pelo ID
func (r *billingPlanRepository) FindBillingPlanByID(id uint) (model.BillingPlanDomainInterface, *rest_err.RestErr) {
	var plan model.BillingPlanDomain

	if err := r.db.First(&plan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Plano de cobrança não encontrado: ID %d", id)
			return nil, rest_err.NewNotFoundError("Plano de cobrança não encontrado")
		}
		log.Printf("Erro ao buscar plano de cobrança no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar plano de cobrança", err)
	}
	return &plan, nil
}

// FindAllBillingPlans busca todos os planos de cobrança
func (r *billingPlanRepository) FindAllBillingPlans() ([]model.BillingPlanDomainInterface, *rest_err.RestErr) {
	var plans []model.BillingPlanDomain

	if err := r.db.Find(&plans).Error; err != nil {
		log.Printf("Erro ao buscar todos os planos de cobrança: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar planos de cobrança", err)
	}

	plansInterface := make([]model.BillingPlanDomainInterface, len(plans))
	for i, plan := range plans {
		plansInterface[i] = &plan
	}
	return plansInterface, nil
}

// UpdateBillingPlan atualiza um plano de cobrança
func (r *billingPlanRepository) UpdateBillingPlan(plan *model.BillingPlanDomain) (*model.BillingPlanDomain, *rest_err.RestErr) {
	if err := r.db.Save(plan).Error; err != nil {
		log.Printf("Erro ao atualizar plano de cobrança no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar plano de cobrança", err)
	}
	return plan, nil
}

// DeleteBillingPlan remove um plano de cobrança do banco de dados
func (r *billingPlanRepository) DeleteBillingPlan(id uint) *rest_err.RestErr {
	var plan model.BillingPlanDomain

	if err := r.db.First(&plan, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Plano de cobrança não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Plano de cobrança não encontrado")
		}
		log.Printf("Erro ao buscar plano de cobrança para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar plano de cobrança para exclusão", err)
	}

	if err := r.db.Delete(&plan).Error; err != nil {
		log.Printf("Erro ao excluir plano de cobrança: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir plano de cobrança", err)
	}

	log.Printf("Plano de cobrança excluído com sucesso: ID %d", id)
	return nil
}
