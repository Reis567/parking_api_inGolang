package registro

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (rc *registroControllerInterface) FindRegistroByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	// Chamar o serviço para buscar o registro
	registro, serviceErr := rc.service.FindRegistroByIDService(uint(id))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registro": registro})
}


func (rc *registroControllerInterface) FindAllRegistros(c *gin.Context) {
	// Chamar o serviço para buscar todos os registros
	registros, err := rc.service.FindAllRegistrosService()
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registros": registros})
}


func (rc *registroControllerInterface) HistoricoRegistros(c *gin.Context) {
	// Ler os parâmetros de query
	dataInicioStr := c.Query("dataInicio")
	dataFimStr := c.Query("dataFim")
	placa := c.Query("placa")
	status := c.Query("status")

	// Verificar se dataInicio e dataFim foram fornecidos
	if dataInicioStr == "" || dataFimStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Os parâmetros 'dataInicio' e 'dataFim' são obrigatórios (formato YYYY-MM-DD)"})
		return
	}

	// Converter as datas
	dataInicio, err := time.Parse("2006-01-02", dataInicioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "dataInicio inválida", "error": err.Error()})
		return
	}
	dataFim, err := time.Parse("2006-01-02", dataFimStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "dataFim inválida", "error": err.Error()})
		return
	}

	// Chamar o service para buscar os registros filtrados
	registros, serviceErr := rc.service.FindHistoricoRegistrosService(dataInicio, dataFim, placa, status)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"historico_registros": registros})
}