package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// BillingPlanDomainInterface define os métodos que a entidade Plano de Cobrança deve implementar
type BillingPlanDomainInterface interface {
	GetID() uint
	GetNomePlano() string
	GetDescricao() string
	GetValor() float64
	GetDuracao() int
	GetTipo() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	AtualizarPlano(nomePlano, descricao string, valor float64, duracao int, tipo string)
}

// NewBillingPlanDomain cria uma nova instância de BillingPlanDomain
func NewBillingPlanDomain(nomePlano, descricao string, valor float64, duracao int, tipo string) BillingPlanDomainInterface {
	plan := &BillingPlanDomain{
		NomePlano:  nomePlano,
		Descricao:  descricao,
		Valor:      valor,
		Duracao:    duracao,
		Tipo:       tipo,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	return plan
}

// BillingPlanDomain representa a estrutura de um plano de cobrança no sistema
type BillingPlanDomain struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	NomePlano  string  `gorm:"not null" json:"nome_plano"`   // Ex.: "6 horas", "Mensal"
	Descricao  string  `json:"descricao"`                   // Detalhes sobre o plano
	Valor      float64 `gorm:"not null" json:"valor"`       // Custo do plano
	Duracao    int     `gorm:"not null" json:"duracao"`     // Tempo associado ao plano (em horas ou dias)
	Tipo       string  `gorm:"not null" json:"tipo"`        // Ex.: "Por hora", "Por período"
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura BillingPlanDomain
func (bp *BillingPlanDomain) GetID() uint {
	return bp.ID
}

func (bp *BillingPlanDomain) GetNomePlano() string {
	return bp.NomePlano
}

func (bp *BillingPlanDomain) GetDescricao() string {
	return bp.Descricao
}

func (bp *BillingPlanDomain) GetValor() float64 {
	return bp.Valor
}

func (bp *BillingPlanDomain) GetDuracao() int {
	return bp.Duracao
}

func (bp *BillingPlanDomain) GetTipo() string {
	return bp.Tipo
}

func (bp *BillingPlanDomain) GetCreatedAt() string {
	return bp.CreatedAt
}

func (bp *BillingPlanDomain) GetUpdatedAt() string {
	return bp.UpdatedAt
}

// GetJSONValue retorna o BillingPlanDomain em formato JSON
func (bp *BillingPlanDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(bp)
	if err != nil {
		return "", fmt.Errorf("erro ao converter BillingPlanDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do plano
func (bp *BillingPlanDomain) SetID(id uint) {
	bp.ID = id
}

// AtualizarPlano atualiza os dados do plano
func (bp *BillingPlanDomain) AtualizarPlano(nomePlano, descricao string, valor float64, duracao int, tipo string) {
	bp.NomePlano = nomePlano
	bp.Descricao = descricao
	bp.Valor = valor
	bp.Duracao = duracao
	bp.Tipo = tipo
	bp.UpdatedAt = time.Now().Format(time.RFC3339)
}
