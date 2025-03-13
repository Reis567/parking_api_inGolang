package relatorios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func (rc *relatoriosController) CalcularTempoMedioPermanencia(c *gin.Context) {
	dataInicioStr := c.Query("dataInicio")
	dataFimStr := c.Query("dataFim")

	// Conversão das datas
	inicio, err := time.Parse("2006-01-02", dataInicioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato de data inválido. Use AAAA-MM-DD"})
		return
	}

	fim, err := time.Parse("2006-01-02", dataFimStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato de data inválido. Use AAAA-MM-DD"})
		return
	}

	tempoMedio, errResp := rc.service.CalcularTempoMedioPermanencia(inicio, fim)
	if errResp != nil {
		c.JSON(errResp.Code, gin.H{"erro": errResp.Message})
		return
	}

	// Responde com o resultado
	c.JSON(http.StatusOK, gin.H{"tempoMedioMinutos": tempoMedio})
}
