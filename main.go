package main

import (
	"fmt"
	"log"
	"os"

	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/controller/routes"
	"meu-novo-projeto/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    logger.Info("About to start application")
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Usar as variáveis de ambiente
    appPort := os.Getenv("APP_PORT")
    if appPort == "" {
        log.Fatal("A porta da aplicação (APP_PORT) não está definida no arquivo .env")
    }

    // Configurar o servidor Gin
    router := gin.Default()

    // Aplicar middleware de erros
    router.Use(middleware.ErrorHandlingMiddleware())

    // Inicializar rotas
    api := router.Group("/api/v1")
    routes.InitRoutes(api)

    // Rodar a aplicação e tratar erro de inicialização
    if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}
