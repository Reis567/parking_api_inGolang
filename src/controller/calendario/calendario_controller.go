package calendario

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"

)

// NewCalendarioController cria uma nova instância do controlador de calendário
func NewCalendarioController(service service.RegistroEstacionamentoDomainService) CalendarioControllerInterface {
	return &calendarioController{
		service: service,
	}
}

// CalendarioControllerInterface define os métodos do controlador
type CalendarioControllerInterface interface {
	ListarRegistrosPorData(c *gin.Context)
}

type calendarioController struct {
	service service.RegistroEstacionamentoDomainService
}

// ListarRegistrosPorData retorna os registros de estacionamento para uma data específica
