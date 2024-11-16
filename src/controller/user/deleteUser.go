package user

import (
	"meu-novo-projeto/src/configuration/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteUser exclui um usuário pelo ID
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userID := c.Param("id") // Obtém o ID do usuário a partir dos parâmetros da URL
	logger.Info("Iniciando DeleteUserController", zap.String("user_id", userID))

	// Valida o ID para garantir que é um número inteiro
	id, err := strconv.Atoi(userID)
	if err != nil || id <= 0 {
		logger.Error("ID inválido para exclusão", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido. Deve ser um número inteiro maior que zero.",
		})
		return
	}

	// Chama o serviço para excluir o usuário
	deleteErr := uc.service.DeleteUser(userID)
	if deleteErr != nil {
		logger.Error("Erro ao excluir usuário no serviço", zap.Error(deleteErr))
		c.JSON(deleteErr.Code, gin.H{
			"error": deleteErr.Message,
		})
		return
	}

	// Retorna uma resposta de sucesso
	logger.Info("Usuário excluído com sucesso", zap.String("user_id", userID))
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário excluído com sucesso!",
	})
}
