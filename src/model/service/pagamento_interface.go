package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"go.uber.org/zap"

)

// PagamentoDomainService define os métodos para o serviço de pagamento
type PagamentoDomainService interface {
	// Cria um novo pagamento
	CreatePagamentoService(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr)
	// Atualiza um pagamento existente, podendo alterar o status (por exemplo, para "Concluído" ou "Cancelado")
	UpdatePagamentoService(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr)
}

// pagamentoDomainService é a implementação da interface PagamentoDomainService
type pagamentoDomainService struct {
	pagamentoRepo repository.PagamentoRepository
}

// NewPagamentoDomainService cria uma nova instância de pagamentoDomainService
func NewPagamentoDomainService(repo repository.PagamentoRepository) PagamentoDomainService {
	return &pagamentoDomainService{
		pagamentoRepo: repo,
	}
}

// CreatePagamentoService cria um novo pagamento utilizando o repositório
func (s *pagamentoDomainService) CreatePagamentoService(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr) {
	zap.L().Info("Criando pagamento", zap.Any("pagamento", pagamento))

	createdPagamento, err := s.pagamentoRepo.CreatePagamento(pagamento)
	if err != nil {
		zap.L().Error("Erro ao criar pagamento", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Pagamento criado com sucesso", zap.Uint("pagamento_id", createdPagamento.GetID()))
	return createdPagamento, nil
}

// UpdatePagamentoService atualiza um pagamento existente no repositório
func (s *pagamentoDomainService) UpdatePagamentoService(pagamento model.PagamentoDomainInterface) (model.PagamentoDomainInterface, *rest_err.RestErr) {
	zap.L().Info("Atualizando pagamento", zap.Any("pagamento", pagamento))

	updatedPagamento, err := s.pagamentoRepo.UpdatePagamento(pagamento)
	if err != nil {
		zap.L().Error("Erro ao atualizar pagamento", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Pagamento atualizado com sucesso", zap.Uint("pagamento_id", updatedPagamento.GetID()))
	return updatedPagamento, nil
}
