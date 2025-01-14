package agendamento

import (
	"meu-novo-projeto/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateAgendamento atualiza os dados de um agendamento
func (ac *agendamentoControllerInterface) UpdateAgendamento(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var agendamentoRequest model.AgendamentoDomain
	if err := c.ShouldBindJSON(&agendamentoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Definir o ID no objeto recebido
	agendamentoRequest.SetID(uint(id))

	// Chamar o serviço para atualizar o agendamento
	updatedAgendamento, updateErr := ac.service.UpdateAgendamentoService(&agendamentoRequest)
	if updateErr != nil {
		c.JSON(updateErr.Code, gin.H{"error": updateErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agendamento atualizado com sucesso", "data": updatedAgendamento})
}
