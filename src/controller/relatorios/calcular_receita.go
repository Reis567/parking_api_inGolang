package relatorios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CalcularReceita calcula a receita total em um intervalo de datas
func (rc *relatoriosController) CalcularReceita(c *gin.Context) {
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

	receita, errRes := rc.service.CalcularReceita(inicio, fim)
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"receita": receita})
}
