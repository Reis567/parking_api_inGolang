package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// UpdateRegistroService atualiza os dados de um registro de estacionamento
func (s *registroEstacionamentoDomainService) UpdateRegistroService(registro model.RegistroEstacionamentoDomainInterface) (model.RegistroEstacionamentoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateRegistro service", zap.Uint("registro_id", registro.GetID()))

	existingRegistro, err := s.repo.FindRegistroByID(registro.GetID())
	if err != nil {
		logger.Error("Erro ao buscar registro para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar campos
	existingRegistro.(*model.RegistroEstacionamentoDomain).HoraSaida = registro.GetHoraSaida()
	existingRegistro.(*model.RegistroEstacionamentoDomain).Status = registro.GetStatus()
	existingRegistro.(*model.RegistroEstacionamentoDomain).ValorCobrado = registro.GetValorCobrado()
	existingRegistro.(*model.RegistroEstacionamentoDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	updatedRegistro, updateErr := s.repo.UpdateRegistro(existingRegistro.(*model.RegistroEstacionamentoDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar registro no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Registro atualizado com sucesso", zap.Uint("registro_id", updatedRegistro.GetID()))
	return updatedRegistro, nil
}
