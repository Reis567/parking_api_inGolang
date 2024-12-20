package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// UpdateVehicleService atualiza os dados de um veículo
func (s *vehicleDomainService) UpdateVehicleService(vehicle model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateVehicleService", zap.Uint("vehicle_id", vehicle.GetID()))

	// Buscar veículo existente no repositório
	existingVehicle, err := s.vehicleRepository.FindVeiculoByID(vehicle.GetID())
	if err != nil {
		logger.Error("Erro ao buscar veículo para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos necessários no veículo existente
	existingVehicle.(*model.VehicleDomain).Placa = vehicle.GetPlaca()
	existingVehicle.(*model.VehicleDomain).Modelo = vehicle.GetModelo()
	existingVehicle.(*model.VehicleDomain).Cor = vehicle.GetCor()
	existingVehicle.(*model.VehicleDomain).Status = vehicle.GetStatus()
	existingVehicle.(*model.VehicleDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar as alterações no repositório
	updatedVehicle, updateErr := s.vehicleRepository.UpdateVeiculo(existingVehicle.(*model.VehicleDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar veículo no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Veículo atualizado com sucesso", zap.Uint("vehicle_id", updatedVehicle.GetID()))
	return updatedVehicle, nil
}
