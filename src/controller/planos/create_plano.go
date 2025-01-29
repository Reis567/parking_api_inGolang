package billingplan

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (bc *billingPlanControllerInterface) CreateBillingPlan(c *gin.Context) {
	var billingPlanRequest model.BillingPlanDomain

	// Fazer o binding do corpo da requisição
	if err := c.ShouldBindJSON(&billingPlanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// Chamar o serviço para criar o plano de cobrança
	createdPlan, err := bc.service.CreateBillingPlanService(&billingPlanRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Plano de cobrança criado com sucesso!", "plano": createdPlan})
}
