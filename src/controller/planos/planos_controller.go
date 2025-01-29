package billingplan

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewBillingPlanControllerInterface cria uma nova instância de BillingPlanControllerInterface
func NewBillingPlanControllerInterface(serviceInterface service.BillingPlanDomainService) BillingPlanControllerInterface {
	return &billingPlanControllerInterface{
		service: serviceInterface,
	}
}

// BillingPlanControllerInterface define os métodos do controlador de plano de cobrança
type BillingPlanControllerInterface interface {
	CreateBillingPlan(c *gin.Context)
	FindBillingPlanByID(c *gin.Context)
	FindAllBillingPlans(c *gin.Context)
	UpdateBillingPlan(c *gin.Context)
	DeleteBillingPlan(c *gin.Context)
}

// billingPlanControllerInterface implementa BillingPlanControllerInterface
type billingPlanControllerInterface struct {
	service service.BillingPlanDomainService
}
