package veiculo

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// FindVeiculoByID busca um veículo pelo ID
func (vc *veiculoControllerInterface) FindVeiculoByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	veiculo, fetchErr := vc.service.FindVeiculoByIDService(uint(id))
	if fetchErr != nil {
		c.JSON(fetchErr.Code, gin.H{"error": fetchErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": veiculo})
}
