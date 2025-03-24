package model

import (
	"encoding/json"
	"fmt"

)

// ParkingHistoryDomainInterface define os métodos para o histórico de estacionamento
type ParkingHistoryDomainInterface interface {
	GetID() uint
	GetUserID() uint
	GetVehicleID() uint
	GetEntrada() string
	GetSaida() string
	GetValorCobrado() float64
	GetJSONValue() (string, error)
	SetID(id uint)
}

// ParkingHistoryDomain representa o histórico de estacionamento de um usuário
type ParkingHistoryDomain struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	UserID      uint    `json:"user_id"`
	VehicleID   uint    `json:"vehicle_id"`
	Entrada     string  `json:"entrada"`
	Saida       string  `json:"saida"`
	ValorCobrado float64 `json:"valor_cobrado"`
	CreatedAt   string  `json:"created_at"`
}

// Métodos Get
func (ph *ParkingHistoryDomain) GetID() uint {
	return ph.ID
}

func (ph *ParkingHistoryDomain) GetUserID() uint {
	return ph.UserID
}

func (ph *ParkingHistoryDomain) GetVehicleID() uint {
	return ph.VehicleID
}

func (ph *ParkingHistoryDomain) GetEntrada() string {
	return ph.Entrada
}

func (ph *ParkingHistoryDomain) GetSaida() string {
	return ph.Saida
}

func (ph *ParkingHistoryDomain) GetValorCobrado() float64 {
	return ph.ValorCobrado
}

// GetJSONValue retorna o histórico de estacionamento em formato JSON
func (ph *ParkingHistoryDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(ph)
	if err != nil {
		return "", fmt.Errorf("erro ao converter ParkingHistoryDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID do histórico de estacionamento
func (ph *ParkingHistoryDomain) SetID(id uint) {
	ph.ID = id
}
