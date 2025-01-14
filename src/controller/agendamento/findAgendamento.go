package agendamento

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FindAgendamentoByID busca um agendamento pelo ID
func (ac *agendamentoControllerInterface) FindAgendamentoByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	agendamento, fetchErr := ac.service.FindAgendamentoByIDService(uint(id))
	if fetchErr != nil {
		c.JSON(fetchErr.Code, gin.H{"error": fetchErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": agendamento})
}

// FindAllAgendamentos busca todos os agendamentos
func (ac *agendamentoControllerInterface) FindAllAgendamentos(c *gin.Context) {
	agendamentos, err := ac.service.FindAllAgendamentosService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": agendamentos})
}
