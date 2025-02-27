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


func (ac *agendamentoControllerInterface) CancelAgendamento(c *gin.Context) {
	// Obter o ID do agendamento a partir da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	// Ler a justificativa opcional do corpo da requisição
	var payload struct {
		Justificativa string `json:"justificativa"`
	}
	// Se não houver corpo, a justificativa ficará vazia
	c.ShouldBindJSON(&payload)

	// Chamar o service para cancelar o agendamento
	updatedAgendamento, serviceErr := ac.service.CancelAgendamentoService(uint(id), payload.Justificativa)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Agendamento cancelado com sucesso",
		"data":    updatedAgendamento,
	})
}