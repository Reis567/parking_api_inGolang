package vaga

import (
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
)

// NewVagaControllerInterface cria uma nova instância de VagaControllerInterface
func NewVagaControllerInterface(serviceInterface service.VagaDomainService) VagaControllerInterface {
	return &vagaControllerInterface{
		service: serviceInterface,
	}
}

// VagaControllerInterface define os métodos do controlador de vaga
type VagaControllerInterface interface {
	CreateVaga(c *gin.Context)
	FindVagaByID(c *gin.Context)
	FindAllVagas(c *gin.Context)
	UpdateVaga(c *gin.Context)
	DeleteVaga(c *gin.Context)
	BuscarVagasDisponiveis(c *gin.Context)
}

// vagaControllerInterface implementa VagaControllerInterface
type vagaControllerInterface struct {
	service service.VagaDomainService
}
