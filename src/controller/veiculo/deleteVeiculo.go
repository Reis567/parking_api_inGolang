package veiculo

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// DeleteVeiculo exclui um veículo pelo ID
func (vc *veiculoControllerInterface) DeleteVeiculo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	deleteErr := vc.service.DeleteVehicleService(uint(id))
	if deleteErr != nil {
		c.JSON(deleteErr.Code, gin.H{"error": deleteErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Veículo excluído com sucesso!"})
}
