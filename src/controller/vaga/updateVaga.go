package vaga

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (vc *vagaControllerInterface) UpdateVaga(c *gin.Context) {
	var vagaRequest model.VagaDomain

	// Fazer o binding do corpo da requisição
	if err := c.ShouldBindJSON(&vagaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// Obter o ID da vaga pela URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	vagaRequest.ID = uint(id)

	// Chamar o serviço para atualizar a vaga
	updatedVaga, errRes := vc.service.UpdateVagaService(&vagaRequest)
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga atualizada com sucesso!", "vaga": updatedVaga})
}
