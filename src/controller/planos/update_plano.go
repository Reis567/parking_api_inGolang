package billingplan

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (bc *billingPlanControllerInterface) UpdateBillingPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var billingPlanRequest model.BillingPlanDomain

	if err := c.ShouldBindJSON(&billingPlanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	billingPlanRequest.SetID(uint(id))

	updatedPlan, errRes := bc.service.UpdateBillingPlanService(&billingPlanRequest)
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plano de cobrança atualizado com sucesso!", "plano": updatedPlan})
}
