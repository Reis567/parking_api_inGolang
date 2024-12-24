package veiculo

import (
	"net/http"
	"meu-novo-projeto/src/model"
	"strconv"
	"github.com/gin-gonic/gin"
)

// UpdateVeiculo atualiza um veículo
func (vc *veiculoControllerInterface) UpdateVeiculo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var veiculoRequest model.VehicleDomainInterface
	if err := c.ShouldBindJSON(&veiculoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}
	veiculoRequest.SetID(uint(id))

	veiculo, updateErr := vc.service.UpdateVehicleService(veiculoRequest)
	if updateErr != nil {
		c.JSON(updateErr.Code, gin.H{"error": updateErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Veículo atualizado com sucesso!", "data": veiculo})
}
