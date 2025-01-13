package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// AgendamentoDomainInterface define os métodos que a entidade Agendamento deve implementar
type AgendamentoDomainInterface interface {
	GetID() uint
	GetPlaca() string
	GetTipoVaga() string
	GetDataHoraReserva() string
	GetStatus() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	AtualizarStatus(status string)
}

// NewAgendamentoDomain cria uma nova instância de AgendamentoDomain
func NewAgendamentoDomain(placa string, tipoVaga string, dataHoraReserva string, status string) AgendamentoDomainInterface {
	agendamento := &AgendamentoDomain{
		Placa:           placa,
		TipoVaga:        tipoVaga,
		DataHoraReserva: dataHoraReserva,
		Status:          status,
		CreatedAt:       time.Now().Format(time.RFC3339),
		UpdatedAt:       time.Now().Format(time.RFC3339),
	}
	return agendamento
}

// AgendamentoDomain representa a estrutura de um agendamento no sistema
type AgendamentoDomain struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Placa           string `json:"placa"`              // Placa do veículo
	TipoVaga        string `json:"tipo_vaga"`         // Tipo de vaga: "carro" ou "moto"
	DataHoraReserva string `json:"data_hora_reserva"` // Data e hora da reserva
	Status          string `json:"status"`           // Status: "confirmada", "cancelada", "concluida"
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura AgendamentoDomain
func (ad *AgendamentoDomain) GetID() uint {
	return ad.ID
}

func (ad *AgendamentoDomain) GetPlaca() string {
	return ad.Placa
}

func (ad *AgendamentoDomain) GetTipoVaga() string {
	return ad.TipoVaga
}

func (ad *AgendamentoDomain) GetDataHoraReserva() string {
	return ad.DataHoraReserva
}

func (ad *AgendamentoDomain) GetStatus() string {
	return ad.Status
}

func (ad *AgendamentoDomain) GetCreatedAt() string {
	return ad.CreatedAt
}

func (ad *AgendamentoDomain) GetUpdatedAt() string {
	return ad.UpdatedAt
}

// GetJSONValue retorna o AgendamentoDomain em formato JSON
func (ad *AgendamentoDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(ad)
	if err != nil {
		return "", fmt.Errorf("erro ao converter AgendamentoDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do agendamento
func (ad *AgendamentoDomain) SetID(id uint) {
	ad.ID = id
}

// AtualizarStatus atualiza o status do agendamento
func (ad *AgendamentoDomain) AtualizarStatus(status string) {
	ad.Status = status
	ad.UpdatedAt = time.Now().Format(time.RFC3339)
}
