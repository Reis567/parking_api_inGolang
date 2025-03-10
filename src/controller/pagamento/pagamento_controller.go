package pagamento

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewPagamentoControllerInterface cria uma nova instância de PagamentoControllerInterface
func NewPagamentoControllerInterface(serviceInterface service.PagamentoDomainService) PagamentoControllerInterface {
	return &pagamentoControllerInterface{
		service: serviceInterface,
	}
}

// PagamentoControllerInterface define os métodos do controlador de pagamento
type PagamentoControllerInterface interface {
	CreatePagamento(c *gin.Context)
	UpdatePagamento(c *gin.Context)
	FindPagamentosHistorico(c *gin.Context)
}

// pagamentoControllerInterface implementa PagamentoControllerInterface
type pagamentoControllerInterface struct {
	service service.PagamentoDomainService
}
