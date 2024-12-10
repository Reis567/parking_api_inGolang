package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindVagaByIDService busca uma vaga pelo ID
func (s *vagaDomainService) FindVagaByIDService(id uint) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindVagaByIDService", zap.Uint("vaga_id", id))

	vaga, err := s.vagaRepository.FindVagaByID(id)
	if err != nil {
		logger.Error("Erro ao buscar vaga no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Vaga encontrada com sucesso", zap.Uint("vaga_id", id))
	return vaga, nil
}

