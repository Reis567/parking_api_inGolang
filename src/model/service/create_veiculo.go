package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// CreateVehicleService cria um novo veículo no sistema
func (s *vehicleDomainService) CreateVehicleService(vehicle model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateVehicle service", zap.String("journey", "Create vehicle"))

	createdVehicle, err := s.vehicleRepository.CreateVeiculo(vehicle)
	if err != nil {
		logger.Error("Erro ao criar veículo no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Veículo criado com sucesso", zap.Uint("vehicle_id", createdVehicle.GetID()))
	return createdVehicle, nil
}
