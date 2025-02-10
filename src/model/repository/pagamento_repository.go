package repository

import (
	"log"

	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
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
	// Outros métodos (Find, Delete, etc.) podem ser adicionados conforme necessário.
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
