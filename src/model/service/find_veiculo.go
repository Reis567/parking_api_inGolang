package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
)

// FindVehicleByIDService busca um veículo pelo ID
func (s *vehicleDomainService) FindVehicleByIDService(id uint) (model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindVehicleByID service", zap.Uint("vehicle_id", id))

	vehicle, err := s.vehicleRepository.FindVeiculoByID(id)
	if err != nil {
		logger.Error("Erro ao buscar veículo no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Veículo encontrado com sucesso", zap.Uint("vehicle_id", id))
	return vehicle, nil
}

// FindAllVehiclesService busca todos os veículos
func (s *vehicleDomainService) FindAllVehiclesService() ([]model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllVehicles service")

	vehicles, err := s.vehicleRepository.FindAllVeiculos()
	if err != nil {
		logger.Error("Erro ao buscar todos os veículos no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Todos os veículos encontrados com sucesso", zap.Int("total", len(vehicles)))
	return vehicles, nil
}
