package relatorios

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewRelatoriosController cria uma nova instância de RelatoriosController
func NewRelatoriosController(serviceInterface service.RelatoriosService) RelatoriosController {
	return &relatoriosController{
		service: serviceInterface,
	}
}

// RelatoriosController define os métodos do controlador de relatórios
type RelatoriosController interface {
	CalcularReceita(c *gin.Context)
	CalcularOcupacaoAtual(c *gin.Context)
	VeiculosMaisFrequentes(c *gin.Context)
	CalcularLotacao(c *gin.Context)
}

// relatoriosController implementa RelatoriosController
type relatoriosController struct {
	service service.RelatoriosService
}
