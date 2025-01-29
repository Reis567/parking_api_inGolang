package billingplan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (bc *billingPlanControllerInterface) DeleteBillingPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	errRes := bc.service.DeleteBillingPlanService(uint(id))
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plano de cobrança excluído com sucesso!"})
}
