package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// CreateRegistroService cria um novo registro de estacionamento
func (s *registroEstacionamentoDomainService) CreateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateRegistro service", zap.String("journey", "Create registro"))

	createdRegistro, err := s.repo.CreateRegistro(registro)
	if err != nil {
		logger.Error("Erro ao criar registro no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registro criado com sucesso", zap.Uint("registro_id", createdRegistro.GetID()))
	return createdRegistro, nil
}
