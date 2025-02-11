package pagamento

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"meu-novo-projeto/src/model"
)

// UpdatePagamento atualiza um pagamento existente
func (pc *pagamentoControllerInterface) UpdatePagamento(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Crie uma struct para receber os dados de atualização
	var pagamentoPayload struct {
		RegistroID      uint    `json:"registro_id" binding:"required"`
		ValorTotal      float64 `json:"valor_total" binding:"required"`
		MetodoPagamento string  `json:"metodo_pagamento" binding:"required"`
		Status          string  `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&pagamentoPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	// Crie o domínio de pagamento e defina o ID
	pagamento := model.NewPagamentoDomain(
		pagamentoPayload.RegistroID,
		pagamentoPayload.ValorTotal,
		pagamentoPayload.MetodoPagamento,
		model.PaymentStatus(pagamentoPayload.Status), // Conversão explícita para PaymentStatus
	)
	
	pagamento.SetID(uint(id))

	updatedPagamento, updateErr := pc.service.UpdatePagamentoService(pagamento)
	if updateErr != nil {
		c.JSON(updateErr.Code, gin.H{"error": updateErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pagamento atualizado com sucesso", "pagamento": updatedPagamento})
}
