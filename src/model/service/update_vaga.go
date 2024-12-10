package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// UpdateVagaService atualiza os dados de uma vaga
func (s *vagaDomainService) UpdateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateVagaService", zap.Uint("vaga_id", vaga.GetID()))

	// Buscar vaga existente
	existingVaga, err := s.vagaRepository.FindVagaByID(vaga.GetID())
	if err != nil {
		logger.Error("Erro ao buscar vaga para atualização", zap.Error(err))
		return nil, err
	}

	// Atualizar os campos necessários
	existingVaga.(*model.VagaDomain).Tipo = vaga.GetTipo()
	existingVaga.(*model.VagaDomain).Status = vaga.GetStatus()
	existingVaga.(*model.VagaDomain).Localizacao = vaga.GetLocalizacao()
	existingVaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Salvar as alterações
	updatedVaga, updateErr := s.vagaRepository.UpdateVaga(existingVaga.(*model.VagaDomain))
	if updateErr != nil {
		logger.Error("Erro ao atualizar vaga no repositório", zap.Error(updateErr))
		return nil, updateErr
	}

	logger.Info("Vaga atualizada com sucesso", zap.Uint("vaga_id", updatedVaga.GetID()))
	return updatedVaga, nil
}
