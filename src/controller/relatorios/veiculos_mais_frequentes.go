package relatorios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// VeiculosMaisFrequentes retorna os veículos mais frequentes em um intervalo de datas
func (rc *relatoriosController) VeiculosMaisFrequentes(c *gin.Context) {
	inicioStr := c.Query("inicio")
	fimStr := c.Query("fim")

	inicio, err := time.Parse(time.RFC3339, inicioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data inicial inválida"})
		return
	}

	fim, err := time.Parse(time.RFC3339, fimStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data final inválida"})
		return
	}

	veiculos, errRes := rc.service.VeiculosMaisFrequentes(inicio, fim)
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"veiculos": veiculos})
}
