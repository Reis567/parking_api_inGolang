package veiculo

import (
	"net/http"
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
)

// CreateVeiculo cria um novo veículo
func (vc *veiculoControllerInterface) CreateVeiculo(c *gin.Context) {
	var veiculoRequest model.VehicleDomainInterface
	if err := c.ShouldBindJSON(&veiculoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	veiculo, err := vc.service.CreateVehicleService(veiculoRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Veículo criado com sucesso!", "data": veiculo})
}
