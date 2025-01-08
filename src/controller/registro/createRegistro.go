package registro

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rc *registroControllerInterface) CreateRegistro(c *gin.Context) {
	var registroRequest model.RegistroEstacionamentoDomain

	// Fazer o binding do corpo da requisição
	if err := c.ShouldBindJSON(&registroRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// Chamar o serviço para criar o registro
	createdRegistro, err := rc.service.CreateRegistroService(&registroRequest)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registro criado com sucesso!", "registro": createdRegistro})
}
