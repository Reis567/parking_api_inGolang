package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// ConfiguracaoDomainInterface define os métodos que a entidade Configuracao deve implementar
type ConfiguracaoDomainInterface interface {
	GetID() uint
	GetTipoVaga() string
	GetValorPorHora() float64
	GetTempoMinimo() int
	GetValorAdicional() float64
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	AtualizarValores(valorPorHora float64, tempoMinimo int, valorAdicional float64)
}

// NewConfiguracaoDomain cria uma nova instância de ConfiguracaoDomain
func NewConfiguracaoDomain(tipoVaga string, valorPorHora float64, tempoMinimo int, valorAdicional float64) ConfiguracaoDomainInterface {
	configuracao := &ConfiguracaoDomain{
		TipoVaga:      tipoVaga,
		ValorPorHora:  valorPorHora,
		TempoMinimo:   tempoMinimo,
		ValorAdicional: valorAdicional,
		CreatedAt:     time.Now().Format(time.RFC3339),
		UpdatedAt:     time.Now().Format(time.RFC3339),
	}
	return configuracao
}

// ConfiguracaoDomain representa a estrutura de uma configuração no sistema
type ConfiguracaoDomain struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TipoVaga      string  `gorm:"not null" json:"tipo_vaga"` // Tipo de vaga: "carro", "moto"
	ValorPorHora  float64 `gorm:"not null" json:"valor_por_hora"` // Valor cobrado por hora
	TempoMinimo   int     `gorm:"not null" json:"tempo_minimo"`    // Tempo mínimo cobrado (em horas)
	ValorAdicional float64 `gorm:"not null" json:"valor_adicional"` // Valor adicional por hora extra
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura ConfiguracaoDomain
func (cd *ConfiguracaoDomain) GetID() uint {
	return cd.ID
}

func (cd *ConfiguracaoDomain) GetTipoVaga() string {
	return cd.TipoVaga
}

func (cd *ConfiguracaoDomain) GetValorPorHora() float64 {
	return cd.ValorPorHora
}

func (cd *ConfiguracaoDomain) GetTempoMinimo() int {
	return cd.TempoMinimo
}

func (cd *ConfiguracaoDomain) GetValorAdicional() float64 {
	return cd.ValorAdicional
}

func (cd *ConfiguracaoDomain) GetCreatedAt() string {
	return cd.CreatedAt
}

func (cd *ConfiguracaoDomain) GetUpdatedAt() string {
	return cd.UpdatedAt
}

// GetJSONValue retorna o ConfiguracaoDomain em formato JSON
func (cd *ConfiguracaoDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(cd)
	if err != nil {
		return "", fmt.Errorf("erro ao converter ConfiguracaoDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID da configuração
func (cd *ConfiguracaoDomain) SetID(id uint) {
	cd.ID = id
}

// AtualizarValores atualiza os valores configurados
func (cd *ConfiguracaoDomain) AtualizarValores(valorPorHora float64, tempoMinimo int, valorAdicional float64) {
	cd.ValorPorHora = valorPorHora
	cd.TempoMinimo = tempoMinimo
	cd.ValorAdicional = valorAdicional
	cd.UpdatedAt = time.Now().Format(time.RFC3339)
}
