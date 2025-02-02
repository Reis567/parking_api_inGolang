package agendamento

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (ac *agendamentoControllerInterface) RegistrarEntrada(c *gin.Context) {
	var entradaRequest struct {
		Placa          string `json:"placa" binding:"required"`
		Modelo         string `json:"modelo" binding:"required"`
		Cor            string `json:"cor" binding:"required"`
		TipoVaga       string `json:"tipo_vaga" binding:"required"`
		PlanoCobrançaID *uint  `json:"plano_cobranca_id,omitempty"` // Opcional
	}

	// Fazer o binding da entrada JSON
	if err := c.ShouldBindJSON(&entradaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// 1. Verificar se existe uma reserva
	reserva, reservaErr := ac.VerificarReservaPorPlacaService(entradaRequest.Placa)
	var vagaID uint

	if reservaErr == nil && reserva != nil {
		// Se houver uma reserva, associar a vaga reservada
		vagaID = reserva.GetVagaID()
	} else {
		// Se não houver uma reserva, buscar a primeira vaga disponível
		vagaDisponivel, vagaErr := ac.BuscarVagaDisponivelService(entradaRequest.TipoVaga)
		if vagaErr != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Nenhuma vaga disponível para o tipo informado"})
			return
		}
		vagaID = vagaDisponivel.GetID()
	}

	// 2. Registrar a entrada do veículo
	registro := model.NewRegistroEstacionamentoDomain(
		entradaRequest.Placa,
		vagaID,
		time.Now().Format(time.RFC3339),
		"entrada",
	)

	// 3. Salvar o registro no banco de dados
	createdRegistro, err := ac.CreateRegistroService(registro)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	// 4. Atualizar o status da vaga para "ocupada"
	vagaUpdateErr := ac.AtualizarStatusVagaService(vagaID, "ocupada")
	if vagaUpdateErr != nil {
		c.JSON(vagaUpdateErr.Code, gin.H{"message": vagaUpdateErr.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Entrada registrada com sucesso", "registro": createdRegistro})
}
