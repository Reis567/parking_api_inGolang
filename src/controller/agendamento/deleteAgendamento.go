package agendamento

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteAgendamento exclui um agendamento pelo ID
func (ac *agendamentoControllerInterface) DeleteAgendamento(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chamar o serviço para deletar o agendamento
	deleteErr := ac.service.DeleteAgendamentoService(uint(id))
	if deleteErr != nil {
		c.JSON(deleteErr.Code, gin.H{"error": deleteErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agendamento excluído com sucesso"})
}
