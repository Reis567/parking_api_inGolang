package registro

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (rc *registroControllerInterface) DeleteRegistro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	// Chamar o serviço para excluir o registro
	serviceErr := rc.service.DeleteRegistroService(uint(id))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registro excluído com sucesso!"})
}
