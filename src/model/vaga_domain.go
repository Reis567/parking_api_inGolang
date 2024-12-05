package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// VagaDomainInterface define os métodos que a entidade Vaga deve implementar
type VagaDomainInterface interface {
	GetID() uint
	GetTipo() string
	GetStatus() string
	GetLocalizacao() string
	GetSerial() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetJSONValue() (string, error)
	SetID(id uint)
	GenerateSerial()
}

// NewVagaDomain cria uma nova instância de VagaDomain
func NewVagaDomain(tipo, status, localizacao string) VagaDomainInterface {
	vaga := &VagaDomain{
		Tipo:        tipo,
		Status:      status,
		Localizacao: localizacao,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	return vaga
}

// VagaDomain representa a estrutura de uma vaga no sistema
type VagaDomain struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Tipo        string `json:"tipo"`        // Ex: "carro", "moto"
	Status      string `json:"status"`      // Ex: "disponível", "ocupada", "reservada"
	Localizacao string `json:"localizacao"` // Ex: "Setor A, Número 15"
	Serial      string `json:"serial"`      // Gerado automaticamente: "A<ID>"
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Métodos Get para cada campo da estrutura VagaDomain
func (vd *VagaDomain) GetID() uint {
	return vd.ID
}

func (vd *VagaDomain) GetTipo() string {
	return vd.Tipo
}

func (vd *VagaDomain) GetStatus() string {
	return vd.Status
}

func (vd *VagaDomain) GetLocalizacao() string {
	return vd.Localizacao
}

func (vd *VagaDomain) GetSerial() string {
	return vd.Serial
}

func (vd *VagaDomain) GetCreatedAt() string {
	return vd.CreatedAt
}

func (vd *VagaDomain) GetUpdatedAt() string {
	return vd.UpdatedAt
}

// GenerateSerial gera o serial da vaga no formato "A<ID>"
func (vd *VagaDomain) GenerateSerial() {
	vd.Serial = fmt.Sprintf("A%d", vd.ID)
}

// GetJSONValue retorna a VagaDomain em formato JSON
func (vd *VagaDomain) GetJSONValue() (string, error) {
	jsonData, err := json.Marshal(vd)
	if err != nil {
		return "", fmt.Errorf("erro ao converter VagaDomain para JSON: %w", err)
	}
	return string(jsonData), nil
}

// SetID define o ID da vaga
func (vd *VagaDomain) SetID(id uint) {
	vd.ID = id
	vd.GenerateSerial() // Atualiza o serial ao definir o ID
}
