package agendamento

import (
	"meu-novo-projeto/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAgendamento cria um novo agendamento
func (ac *agendamentoControllerInterface) CreateAgendamento(c *gin.Context) {
	var agendamentoRequest model.AgendamentoDomain

	// Fazer o binding do corpo da requisição
	if err := c.ShouldBindJSON(&agendamentoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Chamar o serviço para criar o agendamento
	createdAgendamento, err := ac.service.CreateAgendamentoService(&agendamentoRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Agendamento criado com sucesso", "data": createdAgendamento})
}
