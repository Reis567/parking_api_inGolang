package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// RegistroEstacionamentoDomainInterface define os métodos que a entidade RegistroEstacionamento deve implementar
type RegistroEstacionamentoDomainInterface interface {
	GetID() uint
	GetPlaca() string
	GetVagaID() uint
	GetHoraEntrada() string
	GetHoraSaida() string
	GetValorCobrado() float64
	GetStatus() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	RegistrarSaida(saida string, valor float64)
}

// NewRegistroEstacionamentoDomain cria uma nova instância de RegistroEstacionamentoDomain
func NewRegistroEstacionamentoDomain(placa string, vagaID uint, entrada string, status string) RegistroEstacionamentoDomainInterface {
	registro := &RegistroEstacionamentoDomain{
		Placa:       placa,
		VagaID:      vagaID,
		HoraEntrada: entrada,
		Status:      status,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	return registro
}

// RegistroEstacionamentoDomain representa a estrutura de um registro de estacionamento no sistema
type RegistroEstacionamentoDomain struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Placa       string  `json:"placa"`               // Placa do veículo estacionado
	VagaID      uint    `json:"vaga_id"`            // ID da vaga utilizada
	HoraEntrada string  `json:"hora_entrada"`       // Horário de entrada
	HoraSaida   string  `json:"hora_saida"`         // Horário de saída (opcional no momento da entrada)
	ValorCobrado float64 `json:"valor_cobrado"`      // Valor cobrado pelo estacionamento
	Status      string  `json:"status"`            // Status do registro: "entrada" ou "saida"
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura RegistroEstacionamentoDomain
func (rd *RegistroEstacionamentoDomain) GetID() uint {
	return rd.ID
}

func (rd *RegistroEstacionamentoDomain) GetPlaca() string {
	return rd.Placa
}

func (rd *RegistroEstacionamentoDomain) GetVagaID() uint {
	return rd.VagaID
}

func (rd *RegistroEstacionamentoDomain) GetHoraEntrada() string {
	return rd.HoraEntrada
}

func (rd *RegistroEstacionamentoDomain) GetHoraSaida() string {
	return rd.HoraSaida
}

func (rd *RegistroEstacionamentoDomain) GetValorCobrado() float64 {
	return rd.ValorCobrado
}

func (rd *RegistroEstacionamentoDomain) GetStatus() string {
	return rd.Status
}

func (rd *RegistroEstacionamentoDomain) GetCreatedAt() string {
	return rd.CreatedAt
}

func (rd *RegistroEstacionamentoDomain) GetUpdatedAt() string {
	return rd.UpdatedAt
}

// GetJSONValue retorna o RegistroEstacionamentoDomain em formato JSON
func (rd *RegistroEstacionamentoDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(rd)
	if err != nil {
		return "", fmt.Errorf("erro ao converter RegistroEstacionamentoDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do registro
func (rd *RegistroEstacionamentoDomain) SetID(id uint) {
	rd.ID = id
}

// RegistrarSaida atualiza os dados do registro no momento da saída
func (rd *RegistroEstacionamentoDomain) RegistrarSaida(saida string, valor float64) {
	rd.HoraSaida = saida
	rd.ValorCobrado = valor
	rd.Status = "saida"
	rd.UpdatedAt = time.Now().Format(time.RFC3339)
}
