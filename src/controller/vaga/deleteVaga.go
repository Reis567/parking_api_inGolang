package vaga

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (vc *vagaControllerInterface) DeleteVaga(c *gin.Context) {
	// Obter o ID da vaga pela URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	// Chamar o serviço para excluir a vaga
	errRes := vc.service.DeleteVagaService(uint(id))
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga excluída com sucesso!"})
}
