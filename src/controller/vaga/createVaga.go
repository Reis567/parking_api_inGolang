package vaga

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (vc *vagaControllerInterface) CreateVaga(c *gin.Context) {
	var vagaRequest model.VagaDomain

	// Fazer o binding do corpo da requisição
	if err := c.ShouldBindJSON(&vagaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// Chamar o serviço para criar a vaga
	createdVaga, err := vc.service.CreateVagaService(&vagaRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vaga criada com sucesso!", "vaga": createdVaga})
}
