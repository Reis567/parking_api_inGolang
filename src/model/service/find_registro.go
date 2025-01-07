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
func (s *registroEstacionamentoDomainService) FindRegistroByIDService(id uint) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindRegistroByID service", zap.Uint("registro_id", id))

	registro, err := s.repo.FindRegistroByID(id)
	if err != nil {
		logger.Error("Erro ao buscar registro no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Registro encontrado com sucesso", zap.Uint("registro_id", registro.GetID()))
	return registro, nil
}