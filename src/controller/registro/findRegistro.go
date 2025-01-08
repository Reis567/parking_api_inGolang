package registro

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (rc *registroControllerInterface) FindRegistroByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	// Chamar o serviço para buscar o registro
	registro, serviceErr := rc.service.FindRegistroByIDService(uint(id))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registro": registro})
}


func (rc *registroControllerInterface) FindAllRegistros(c *gin.Context) {
	// Chamar o serviço para buscar todos os registros
	registros, err := rc.service.FindAllRegistrosService()
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registros": registros})
}