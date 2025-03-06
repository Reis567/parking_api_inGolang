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

	veiculo, fetchErr := vc.service.FindVehicleByIDService(uint(id))
	if fetchErr != nil {
		c.JSON(fetchErr.Code, gin.H{"error": fetchErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": veiculo})
}

func (vc *veiculoControllerInterface) FindAllVeiculos(c *gin.Context) {
	veiculos, err := vc.service.FindAllVehiclesService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": veiculos})
}



func (vc *veiculoControllerInterface) FindVeiculosAtivos(c *gin.Context) {
	veiculos, err := vc.service.FindAllVeiculosAtivosService()
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"veiculos_ativos": veiculos})
}