package agendamento

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewAgendamentoControllerInterface cria uma nova instância de AgendamentoControllerInterface
func NewAgendamentoControllerInterface(serviceInterface service.AgendamentoDomainService) AgendamentoControllerInterface {
	return &agendamentoControllerInterface{
		service: serviceInterface,
	}
}

// AgendamentoControllerInterface define os métodos do controlador de agendamento
type AgendamentoControllerInterface interface {
	CreateAgendamento(c *gin.Context)
	FindAgendamentoByID(c *gin.Context)
	FindAllAgendamentos(c *gin.Context)
	UpdateAgendamento(c *gin.Context)
	DeleteAgendamento(c *gin.Context)
}

// agendamentoControllerInterface implementa AgendamentoControllerInterface
type agendamentoControllerInterface struct {
	service service.AgendamentoDomainService
}
