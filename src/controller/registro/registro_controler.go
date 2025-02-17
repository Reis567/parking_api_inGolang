package registro

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewRegistroControllerInterface cria uma nova instância de RegistroControllerInterface
func NewRegistroControllerInterface(serviceInterface service.RegistroEstacionamentoDomainService) RegistroControllerInterface {
	return &registroControllerInterface{
		service: serviceInterface,
	}
}

// RegistroControllerInterface define os métodos do controlador de registro
type RegistroControllerInterface interface {
	CreateRegistro(c *gin.Context)
	FindRegistroByID(c *gin.Context)
	FindAllRegistros(c *gin.Context)
	UpdateRegistro(c *gin.Context)
	DeleteRegistro(c *gin.Context)
	HistoricoRegistros(c *gin.Context)  // Novo método

}

// registroControllerInterface implementa RegistroControllerInterface
type registroControllerInterface struct {
	service service.RegistroEstacionamentoDomainService
}
