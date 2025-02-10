package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type PaymentStatus string

const (
	PaymentStatusAberto     PaymentStatus = "Aberto"
	PaymentStatusConcluido  PaymentStatus = "Concluido"
	PaymentStatusCancelado  PaymentStatus = "Cancelado"
)
type PagamentoDomainInterface interface {
	GetID() uint
	GetRegistroID() uint
	GetValorTotal() float64
	GetMetodoPagamento() string
	GetStatus() PaymentStatus
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	AtualizarStatus(status PaymentStatus)
	AtualizarMetodoPagamento(metodo string)
}

// NewPagamentoDomain cria uma nova instância de PagamentoDomain
func NewPagamentoDomain(registroID uint, valorTotal float64, metodoPagamento string, status PaymentStatus) PagamentoDomainInterface {
	pagamento := &PagamentoDomain{
		RegistroID:      registroID,
		ValorTotal:      valorTotal,
		MetodoPagamento: metodoPagamento,
		Status:          status,
		CreatedAt:       time.Now().Format(time.RFC3339),
		UpdatedAt:       time.Now().Format(time.RFC3339),
	}
	return pagamento
}

// PagamentoDomain representa a estrutura de um pagamento no sistema
type PagamentoDomain struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	RegistroID      uint    `json:"registro_id"`          // ID do registro relacionado
	ValorTotal      float64 `json:"valor_total"`          // Valor total pago
	MetodoPagamento string  `json:"metodo_pagamento"`    // Método de pagamento: "cartão", "dinheiro", etc.
	Status          PaymentStatus   `json:"status"`              // Status do pagamento: "pendente", "concluído"
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura PagamentoDomain
func (pd *PagamentoDomain) GetID() uint {
	return pd.ID
}

func (pd *PagamentoDomain) GetRegistroID() uint {
	return pd.RegistroID
}

func (pd *PagamentoDomain) GetValorTotal() float64 {
	return pd.ValorTotal
}

func (pd *PagamentoDomain) GetMetodoPagamento() string {
	return pd.MetodoPagamento
}

func (pd *PagamentoDomain) GetStatus() PaymentStatus  {
	return pd.Status
}

func (pd *PagamentoDomain) GetCreatedAt() string {
	return pd.CreatedAt
}

func (pd *PagamentoDomain) GetUpdatedAt() string {
	return pd.UpdatedAt
}

// GetJSONValue retorna o PagamentoDomain em formato JSON
func (pd *PagamentoDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(pd)
	if err != nil {
		return "", fmt.Errorf("erro ao converter PagamentoDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do pagamento
func (pd *PagamentoDomain) SetID(id uint) {
	pd.ID = id
}

// AtualizarStatus atualiza o status do pagamento
func (pd *PagamentoDomain) AtualizarStatus(status PaymentStatus ) {
	pd.Status = status
	pd.UpdatedAt = time.Now().Format(time.RFC3339)
}

// AtualizarMetodoPagamento atualiza o método de pagamento
func (pd *PagamentoDomain) AtualizarMetodoPagamento(metodo string) {
	pd.MetodoPagamento = metodo
	pd.UpdatedAt = time.Now().Format(time.RFC3339)
}
