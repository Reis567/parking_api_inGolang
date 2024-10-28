package model

import (
	"fmt"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"time"

	"go.uber.org/zap"
)

// CreateUser cria um novo usuário no sistema
func (ud *UserDomain) CreateUser(user UserDomain) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init CreateUser model", zap.String("journey", "Create user"))

	// Atribuir ID e datas
	user.ID = generateID()
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	// Encriptar a senha
	user.EncryptPassword()

	// Log e retorno do usuário criado (simulação)
	logger.Info("Usuário criado com sucesso", zap.String("user_id", user.ID))
	return &user, nil
}

// Exemplo de função para geração de ID
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
