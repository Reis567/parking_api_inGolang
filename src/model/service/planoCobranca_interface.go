package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"time"

	"go.uber.org/zap"
)

// BillingPlanDomainService define os métodos para o serviço de planos de cobrança
type BillingPlanDomainService interface {
	CreateBillingPlanService(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr)
	FindBillingPlanByIDService(id uint) (model.BillingPlanDomainInterface, *rest_err.RestErr)
	FindAllBillingPlansService() ([]model.BillingPlanDomainInterface, *rest_err.RestErr)
	UpdateBillingPlanService(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr)
	DeleteBillingPlanService(id uint) *rest_err.RestErr
}

// billingPlanDomainService implementa a interface BillingPlanDomainService
type billingPlanDomainService struct {
	repo repository.BillingPlanRepository
}

// NewBillingPlanDomainService cria uma nova instância de billingPlanDomainService
func NewBillingPlanDomainService(repo repository.BillingPlanRepository) BillingPlanDomainService {
	return &billingPlanDomainService{repo: repo}
}

// CreateBillingPlanService cria um novo plano de cobrança
func (s *billingPlanDomainService) CreateBillingPlanService(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateBillingPlan service")

	createdPlan, err := s.repo.CreateBillingPlan(plan)
	if err != nil {
		logger.Error("Erro ao criar plano de cobrança no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Plano de cobrança criado com sucesso", zap.Uint("plan_id", createdPlan.GetID()))
	return createdPlan, nil
}

// FindBillingPlanByIDService busca um plano de cobrança pelo ID
func (s *billingPlanDomainService) FindBillingPlanByIDService(id uint) (model.BillingPlanDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindBillingPlanByID service", zap.Uint("plan_id", id))

	plan, err := s.repo.FindBillingPlanByID(id)
	if err != nil {
		logger.Error("Erro ao buscar plano de cobrança no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Plano de cobrança encontrado com sucesso", zap.Uint("plan_id", plan.GetID()))
	return plan, nil
}

// FindAllBillingPlansService busca todos os planos de cobrança
func (s *billingPlanDomainService) FindAllBillingPlansService() ([]model.BillingPlanDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllBillingPlans service")

	plans, err := s.repo.FindAllBillingPlans()
	if err != nil {
		logger.Error("Erro ao buscar todos os planos de cobrança no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Planos de cobrança retornados com sucesso", zap.Int("count", len(plans)))
	return plans, nil
}

// UpdateBillingPlanService atualiza os dados de um plano de cobrança
func (s *billingPlanDomainService) UpdateBillingPlanService(plan model.BillingPlanDomainInterface) (model.BillingPlanDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateBillingPlan service", zap.Uint("plan_id", plan.GetID()))

	existingPlan, err := s.repo.FindBillingPlanByID(plan.GetID())
	if err != nil {
		logger.Error("Erro ao buscar plano de cobrança para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos do plano de cobrança
	existingPlan.(*model.BillingPlanDomain).NomePlano = plan.GetNomePlano()
	existingPlan.(*model.BillingPlanDomain).Descricao = plan.GetDescricao()
	existingPlan.(*model.BillingPlanDomain).Valor = plan.GetValor()
	existingPlan.(*model.BillingPlanDomain).Duracao = plan.GetDuracao()
	existingPlan.(*model.BillingPlanDomain).Tipo = plan.GetTipo()
	existingPlan.(*model.BillingPlanDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	updatedPlan, updateErr := s.repo.UpdateBillingPlan(existingPlan.(*model.BillingPlanDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar plano de cobrança no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Plano de cobrança atualizado com sucesso", zap.Uint("plan_id", updatedPlan.GetID()))
	return updatedPlan, nil
}

// DeleteBillingPlanService remove um plano de cobrança
func (s *billingPlanDomainService) DeleteBillingPlanService(id uint) *rest_err.RestErr {
	logger.Info("Init DeleteBillingPlan service", zap.Uint("plan_id", id))

	deleteErr := s.repo.DeleteBillingPlan(id)
	if deleteErr != nil {
		logger.Error("Erro ao excluir plano de cobrança no repositório", zap.Error(deleteErr))
		return deleteErr
	}

	logger.Info("Plano de cobrança excluído com sucesso", zap.Uint("plan_id", id))
	return nil
}
