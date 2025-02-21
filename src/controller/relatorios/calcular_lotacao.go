package relatorios

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *relatoriosController) CalcularLotacao(c *gin.Context) {
	periodo := c.Query("periodo")
	tipo := c.Query("tipo") // Opcional

	if periodo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O parâmetro 'periodo' é obrigatório (diario, semanal ou mensal)"})
		return
	}

	ocupacao, err := rc.service.CalcularLotacaoHistorica(periodo, tipo)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lotacao": ocupacao})
}
