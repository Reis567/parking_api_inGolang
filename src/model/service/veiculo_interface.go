package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
	"go.uber.org/zap"
	"time"
	"meu-novo-projeto/src/configuration/logger"
)

// NewVehicleDomainService cria uma instância de vehicleDomainService
func NewVehicleDomainService(vehicleRepository repository.VeiculoRepository) VehicleDomainService {
	return &vehicleDomainService{vehicleRepository}
}

type vehicleDomainService struct {
	vehicleRepository repository.VeiculoRepository
}

// VehicleDomainService define os métodos para o serviço de veículo
type VehicleDomainService interface {
	CreateVehicleService(vehicle model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr)
	FindVehicleByIDService(id uint) (model.VehicleDomainInterface, *rest_err.RestErr)
	FindAllVehiclesService() ([]model.VehicleDomainInterface, *rest_err.RestErr)
	UpdateVehicleService(vehicle model.VehicleDomainInterface) (model.VehicleDomainInterface, *rest_err.RestErr)
	DeleteVehicleService(id uint) *rest_err.RestErr
}

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


func (s *vehicleDomainService) FindAllVeiculosAtivosService() ([]model.VehicleDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando busca por veículos ativos")
	veiculos, err := s.vehicleRepository.FindVeiculosAtivos()
	if err != nil {
		logger.Error("Erro ao buscar veículos ativos", zap.Error(err))
		return nil, err
	}
	logger.Info("Veículos ativos encontrados", zap.Int("total", len(veiculos)))
	return veiculos, nil
}
