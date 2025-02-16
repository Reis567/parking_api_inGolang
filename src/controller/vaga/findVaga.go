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



// BuscarVagasDisponiveis retorna a vaga disponível para o tipo informado.
// Exemplo de uso: GET /vagas/disponiveis?tipo=carro
func (vc *vagaControllerInterface) BuscarVagasDisponiveis(c *gin.Context) {
	tipo := c.Query("tipo")
	if tipo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O parâmetro 'tipo' é obrigatório"})
		return
	}

	vaga, err := vc.service.BuscarVagaDisponivelService(tipo)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vaga_disponivel": vaga})
}
