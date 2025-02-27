package agendamento

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewAgendamentoControllerInterface cria uma nova instância de AgendamentoControllerInterface
func NewAgendamentoControllerInterface(
	agendamentoService service.AgendamentoDomainService, 
	vagaService service.VagaDomainService, 
	registroService service.RegistroEstacionamentoDomainService,
	pagamentoService service.PagamentoDomainService,
) AgendamentoControllerInterface {
	return &agendamentoControllerInterface{
		service:          agendamentoService,
		vagaService:      vagaService,
		registroService:  registroService,
		pagamentoService: pagamentoService,
	}
}

// AgendamentoControllerInterface define os métodos do controlador de agendamento
type AgendamentoControllerInterface interface {
	CreateAgendamento(c *gin.Context)
	FindAgendamentoByID(c *gin.Context)
	FindAllAgendamentos(c *gin.Context)
	UpdateAgendamento(c *gin.Context)
	DeleteAgendamento(c *gin.Context)
	RegistrarEntrada(c *gin.Context)
	FinalizarEstacionamento(c *gin.Context)
	Reservas(c *gin.Context)
	CancelAgendamento(c *gin.Context)
}

// agendamentoControllerInterface implementa AgendamentoControllerInterface
type agendamentoControllerInterface struct {
	service          service.AgendamentoDomainService
	vagaService      service.VagaDomainService
	registroService  service.RegistroEstacionamentoDomainService
	pagamentoService service.PagamentoDomainService
}
