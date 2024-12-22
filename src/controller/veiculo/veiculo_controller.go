package veiculo

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewVeiculoControllerInterface cria uma nova instância de VeiculoControllerInterface
func NewVeiculoControllerInterface(serviceInterface service.VehicleDomainService) VeiculoControllerInterface {
	return &veiculoControllerInterface{
		service: serviceInterface,
	}
}

// VeiculoControllerInterface define os métodos do controlador de veículo
type VeiculoControllerInterface interface {
	CreateVeiculo(c *gin.Context)
	FindVeiculoByID(c *gin.Context)
	FindAllVeiculos(c *gin.Context)
	UpdateVeiculo(c *gin.Context)
	DeleteVeiculo(c *gin.Context)
}

// veiculoControllerInterface implementa VeiculoControllerInterface
type veiculoControllerInterface struct {
	service service.VehicleDomainService
}
