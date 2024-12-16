package repository

import (
	"log"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"gorm.io/gorm"
)

// DeleteVeiculo exclui um veículo pelo ID
func (r *veiculoRepository) DeleteVeiculo(id uint) *rest_err.RestErr {
	var veiculo model.VehicleDomain

	if err := r.db.First(&veiculo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Veículo não encontrado para exclusão: ID %d", id)
			return rest_err.NewNotFoundError("Veículo não encontrado para exclusão")
		}
		log.Printf("Erro ao buscar veículo para exclusão: %v", err)
		return rest_err.NewInternalServerError("Erro ao buscar veículo para exclusão", err)
	}

	if err := r.db.Delete(&veiculo).Error; err != nil {
		log.Printf("Erro ao excluir veículo do banco de dados: %v", err)
		return rest_err.NewInternalServerError("Erro ao excluir veículo", err)
	}

	log.Printf("Veículo excluído com sucesso: ID %d", id)
	return nil
}
