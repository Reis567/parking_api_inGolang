package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindAllRegistrosService retorna todos os registros de estacionamento
func (s *registroEstacionamentoDomainService) FindAllRegistrosService() ([]model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllRegistros service")

	registros, err := s.repo.FindAllRegistros()
	if err != nil {
		logger.Error("Erro ao buscar todos os registros no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registros retornados com sucesso", zap.Int("count", len(registros)))
	return registros, nil
}
