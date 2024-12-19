package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// UpdateVehicleService atualiza um veículo
func (s *vehicleDomainService) UpdateVehicleService(vehicle model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateVehicle service", zap.Uint("vehicle_id", vehicle.GetID()))

	// Buscar o veículo existente
	existingVehicle, err := s.vehicleRepository.FindVeiculoByID(vehicle.GetID())
	if err != nil {
		logger.Error("Erro ao buscar veículo no repositório", zap.Error(err))
		return nil, err
	}

	// Atualizar campos
	existingVehicle.(*model.VehicleDomain).Plate = vehicle.GetPlate()
	existingVehicle.(*model.VehicleDomain).Type = vehicle.GetType()
	existingVehicle.(*model.VehicleDomain).Owner = vehicle.GetOwner()
	existingVehicle.(*model.VehicleDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar no repositório
	updatedVehicle, updateErr := s.vehicleRepository.UpdateVeiculo(existingVehicle.(*model.VehicleDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar veículo no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Veículo atualizado com sucesso", zap.Uint("vehicle_id", updatedVehicle.GetID()))
	return updatedVehicle, nil
}
