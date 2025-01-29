package billingplan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (bc *billingPlanControllerInterface) FindBillingPlanByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	plan, errRes := bc.service.FindBillingPlanByIDService(uint(id))
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plano": plan})
}


func (bc *billingPlanControllerInterface) FindAllBillingPlans(c *gin.Context) {
	plans, errRes := bc.service.FindAllBillingPlansService()
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"planos": plans})
}