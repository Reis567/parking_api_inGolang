package calendario

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


func (cc *calendarioController) ListarRegistrosPorData(c *gin.Context) {
	dataString := c.Param("data")
	data, err := time.Parse("2006-01-02", dataString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Formato de data inválido, use YYYY-MM-DD"})
		return
	}

	registros, errRes := cc.service.FindRegistrosPorDataService(data)
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dataString, "registros": registros})
}
