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


func (ac *agendamentoControllerInterface) Reservas(c *gin.Context) {
	// Ler o parâmetro de query "status". Se não informado, utiliza "confirmada" como padrão.
	status := c.Query("status")
	if status == "" {
		status = "confirmada"
	}

	// Chama o service para buscar as reservas ativas com o status informado
	reservas, err := ac.service.FindReservasAtivasService(status)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reservas": reservas})
}
