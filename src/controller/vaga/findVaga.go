package vaga

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FindVagaByID busca uma vaga pelo ID
func (vc *vagaControllerInterface) FindVagaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	vaga, errRes := vc.service.FindVagaByIDService(uint(id))
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vaga": vaga})
}

// FindAllVagas busca todas as vagas
func (vc *vagaControllerInterface) FindAllVagas(c *gin.Context) {
	vagas, errRes := vc.service.FindAllVagasService()
	if errRes != nil {
		c.JSON(errRes.Code, gin.H{"message": errRes.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vagas": vagas})
}
