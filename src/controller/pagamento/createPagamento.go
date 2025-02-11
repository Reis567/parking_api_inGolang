package pagamento

import (
	"net/http"
	"meu-novo-projeto/src/model"
	"github.com/gin-gonic/gin"
)

// CreatePagamento cria um novo registro de pagamento
func (pc *pagamentoControllerInterface) CreatePagamento(c *gin.Context) {
	var pagamentoRequest struct {
		RegistroID      uint    `json:"registro_id" binding:"required"`
		ValorTotal      float64 `json:"valor_total" binding:"required"`
		MetodoPagamento string  `json:"metodo_pagamento" binding:"required"`
	}

	if err := c.ShouldBindJSON(&pagamentoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	// Cria um novo pagamento com status "Aberto" (ou utilizando a constante, ex: model.PaymentStatusAberto)
	pagamento := model.NewPagamentoDomain(
		pagamentoRequest.RegistroID,
		pagamentoRequest.ValorTotal,
		pagamentoRequest.MetodoPagamento,
		model.PaymentStatusAberto,
	)

	createdPagamento, serviceErr := pc.service.CreatePagamentoService(pagamento)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, gin.H{"message": serviceErr.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Pagamento criado com sucesso",
		"pagamento": createdPagamento,
	})
}
