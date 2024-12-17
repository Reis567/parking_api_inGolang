package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"meu-novo-projeto/src/model/repository"
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
