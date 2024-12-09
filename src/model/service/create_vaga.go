package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)

// CreateVagaService cria uma nova vaga
func (s *vagaDomainService) CreateVagaService(vaga model.VagaDomainInterface) (model.VagaDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateVagaService", zap.String("journey", "Create vaga"))

	// Atribuir timestamps de criação e atualização
	vaga.(*model.VagaDomain).CreatedAt = time.Now().Format(time.RFC3339)
	vaga.(*model.VagaDomain).UpdatedAt = time.Now().Format(time.RFC3339)

	// Criar a vaga no repositório
	createdVaga, err := s.vagaRepository.CreateVaga(vaga)
	if err != nil {
		logger.Error("Erro ao criar vaga no repositório", zap.Error(err))
		return nil, err
	}

	logger.Info("Vaga criada com sucesso", zap.String("vaga_id", createdVaga.GetSerial()))
	return createdVaga, nil
}
