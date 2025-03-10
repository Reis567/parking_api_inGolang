package agendamento

import (
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strconv"
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
	reserva, reservaErr := ac.service.VerificarReservaPorPlacaService(entradaRequest.Placa)
	var vagaID uint

	if reservaErr == nil && reserva != nil {
		// Se houver uma reserva, associar a vaga reservada
		vagaID = reserva.GetVagaID()
	} else {
		// Se não houver uma reserva, buscar a primeira vaga disponível
		vagaDisponivel, vagaErr := ac.vagaService.BuscarVagaDisponivelService(entradaRequest.TipoVaga)
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
	createdRegistro, err := ac.registroService.CreateRegistroService(registro)
	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	// 4. Atualizar o status da vaga para "ocupada"
	vagaUpdateErr := ac.vagaService.AtualizarStatusVagaService(vagaID, "ocupada")
	if vagaUpdateErr != nil {
		c.JSON(vagaUpdateErr.Code, gin.H{"message": vagaUpdateErr.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Entrada registrada com sucesso", "registro": createdRegistro})
}


func (ac *agendamentoControllerInterface) FinalizarEstacionamento(c *gin.Context) {
	// Obter o registroID da URL (exemplo: /estacionamento/saida/{registroID})
	registroIDParam := c.Param("registroID")
	registroID, err := strconv.ParseUint(registroIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "RegistroID inválido", "error": err.Error()})
		return
	}

	// Definir a hora de saída como a hora atual
	horaSaida := time.Now().Format(time.RFC3339)

	// Chamar o service para finalizar o estacionamento
	resultado, serviceErr := ac.service.FinalizarEstacionamentoService(uint(registroID), horaSaida)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	// O resultado deve ser o registro atualizado; faça a conversão
	registro, ok := resultado.(*model.RegistroEstacionamentoDomain)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao processar o registro finalizado"})
		return
	}

	// Agora, criar o registro de pagamento com status "Aberto"
	// (Utilizando a constante PaymentStatusAberto definida no domínio)
	pagamento := model.NewPagamentoDomain(
		registro.GetID(),            // Associação ao registro finalizado
		registro.GetValorCobrado(),  // Valor total calculado
		"indefinido",                // Método de pagamento (pode ser ajustado conforme a lógica)
		model.PaymentStatusAberto,   // Status do pagamento
	)
	createdPagamento, pagamentoErr := ac.pagamentoService.CreatePagamentoService(pagamento)
	if pagamentoErr != nil {
		c.JSON(pagamentoErr.Code, gin.H{"message": pagamentoErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Estacionamento finalizado com sucesso e pagamento aberto",
		"registro":   registro,
		"pagamento":  createdPagamento,
	})
}