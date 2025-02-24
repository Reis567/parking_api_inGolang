package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
	"time"
)

// pagamentoRepository é a estrutura que implementa a interface PagamentoRepository
type pagamentoRepository struct {
	db *gorm.DB
}

// NewPagamentoRepository cria uma nova instância de pagamentoRepository
func NewPagamentoRepository() PagamentoRepository {
	return &pagamentoRepository{db: database.DB}
}

// PagamentoRepository interface define os métodos para gerenciar pagamentos
type PagamentoRepository interface {
	CreatePagamento(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr)
	UpdatePagamento(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr)
	FindPagamentosPorPeriodo(inicio, fim time.Time, status, metodo string) ([]model.PagamentoDomainInterface, *rest_err.RestErr)
}

// CreatePagamento insere um novo pagamento no banco de dados
func (r *pagamentoRepository) CreatePagamento(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr) {
	if err := r.db.Create(pagamento).Error; err != nil {
		log.Printf("Erro ao inserir pagamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao criar pagamento", err)
	}
	return pagamento, nil
}

// UpdatePagamento atualiza um pagamento existente no banco de dados,
// permitindo, por exemplo, alterar o status para 'Concluído' ou 'Cancelado'.
func (r *pagamentoRepository) UpdatePagamento(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr) {
	if err := r.db.Save(pagamento).Error; err != nil {
		log.Printf("Erro ao atualizar pagamento no banco de dados: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao atualizar pagamento", err)
	}
	return pagamento, nil
}


func (r *pagamentoRepository) FindPagamentosPorPeriodo(inicio, fim time.Time, status, metodo string) ([]model.PagamentoDomainInterface, *rest_err.RestErr) {
	var pagamentos []model.PagamentoDomain

	// Inicia a query sobre a tabela de pagamentos
	query := r.db.Model(&model.PagamentoDomain{}).Where("created_at BETWEEN ? AND ?", inicio.Format(time.RFC3339), fim.Format(time.RFC3339))

	// Se o status for informado, aplica o filtro (note que o campo é do tipo PaymentStatus)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Se o método de pagamento for informado, aplica o filtro
	if metodo != "" {
		query = query.Where("metodo_pagamento = ?", metodo)
	}

	if err := query.Find(&pagamentos).Error; err != nil {
		log.Printf("Erro ao buscar pagamentos por período: %v", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar pagamentos", err)
	}

	// Converter para a interface
	pagamentosInterfaces := make([]model.PagamentoDomainInterface, len(pagamentos))
	for i := range pagamentos {
		pagamentosInterfaces[i] = &pagamentos[i]
	}

	return pagamentosInterfaces, nil
}
