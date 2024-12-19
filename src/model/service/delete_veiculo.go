package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"go.uber.org/zap"
)

// DeleteVehicleService exclui um veículo pelo ID
func (s *vehicleDomainService) DeleteVehicleService(id uint) *rest_err.RestErr {
	logger.Info("Init DeleteVehicle service", zap.Uint("vehicle_id", id))

	err := s.vehicleRepository.DeleteVeiculo(id)
	if err != nil {
		logger.Error("Erro ao excluir veículo no repositório", zap.Error(err))
		return err
	}

	logger.Info("Veículo excluído com sucesso", zap.Uint("vehicle_id", id))
	return nil
}
