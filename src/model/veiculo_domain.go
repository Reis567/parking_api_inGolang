package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// VehicleDomainInterface define os métodos que a entidade Veículo deve implementar
type VehicleDomainInterface interface {
	GetID() uint
	GetPlaca() string
	GetModelo() string
	GetCor() string
	GetStatus() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	UpdateStatus(status string)
}

// NewVehicleDomain cria uma nova instância de VehicleDomain
func NewVehicleDomain(placa, modelo, cor, status string) VehicleDomainInterface {
	vehicle := &VehicleDomain{
		Placa:     placa,
		Modelo:    modelo,
		Cor:       cor,
		Status:    status,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	return vehicle
}

// VehicleDomain representa a estrutura de um veículo no sistema
type VehicleDomain struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Placa     string `json:"placa"`     // Ex: "ABC-1234"
	Modelo    string `json:"modelo"`    // Ex: "Honda Civic"
	Cor       string `json:"cor"`       // Ex: "Preto"
	Status    string `json:"status"`    // Ex: "estacionado", "fora"
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura VehicleDomain
func (vd *VehicleDomain) GetID() uint {
	return vd.ID
}

func (vd *VehicleDomain) GetPlaca() string {
	return vd.Placa
}

func (vd *VehicleDomain) GetModelo() string {
	return vd.Modelo
}

func (vd *VehicleDomain) GetCor() string {
	return vd.Cor
}

func (vd *VehicleDomain) GetStatus() string {
	return vd.Status
}

func (vd *VehicleDomain) GetCreatedAt() string {
	return vd.CreatedAt
}

func (vd *VehicleDomain) GetUpdatedAt() string {
	return vd.UpdatedAt
}

// GetJSONValue retorna a VehicleDomain em formato JSON
func (vd *VehicleDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(vd)
	if err != nil {
		return "", fmt.Errorf("erro ao converter VehicleDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do veículo
func (vd *VehicleDomain) SetID(id uint) {
	vd.ID = id
}

// UpdateStatus atualiza o status do veículo
func (vd *VehicleDomain) UpdateStatus(status string) {
	vd.Status = status
	vd.UpdatedAt = time.Now().Format(time.RFC3339)
}
