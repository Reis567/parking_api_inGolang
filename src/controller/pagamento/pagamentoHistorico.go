package pagamento

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FindPagamentosHistorico lista os pagamentos realizados, permitindo filtrar por período, status e método de pagamento.
// Exemplo de URL: /pagamento/historico?dataInicio=2025-01-01&dataFim=2025-01-31&status=Concluido&metodo=cartao
func (pc *pagamentoControllerInterface) FindPagamentosHistorico(c *gin.Context) {
	// Ler os parâmetros de query
	dataInicioStr := c.Query("dataInicio")
	dataFimStr := c.Query("dataFim")
	status := c.Query("status")
	metodo := c.Query("metodo")

	// Verificar se os parâmetros de data foram informados
	if dataInicioStr == "" || dataFimStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Os parâmetros 'dataInicio' e 'dataFim' são obrigatórios (formato YYYY-MM-DD)",
		})
		return
	}

	// Converter as datas (formato "2006-01-02")
	dataInicio, err := time.Parse("2006-01-02", dataInicioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dataInicio inválida",
			"error":   err.Error(),
		})
		return
	}

	dataFim, err := time.Parse("2006-01-02", dataFimStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dataFim inválida",
			"error":   err.Error(),
		})
		return
	}

	// Chamar o service para buscar os pagamentos históricos com os filtros
	pagamentos, serviceErr := pc.service.FindPagamentosHistoricoService(dataInicio, dataFim, status, metodo)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pagamentos_historico": pagamentos})
}
