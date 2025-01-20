package relatorios

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CalcularOcupacaoAtual calcula a ocupação atual do estacionamento
func (rc *relatoriosController) CalcularOcupacaoAtual(c *gin.Context) {
	ocupacao, errRes := rc.service.CalcularOcupacaoAtual()
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ocupacao": ocupacao})
}
