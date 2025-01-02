package main

import (
	"fmt"
	"log"
	"os"

	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/controller/routes"
	"meu-novo-projeto/src/controller/user"
	"meu-novo-projeto/src/controller/vaga"
	"meu-novo-projeto/src/controller/veiculo" // Import para o controlador de veículos
	"meu-novo-projeto/src/middleware"
	"meu-novo-projeto/src/model/repository"
	"meu-novo-projeto/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Mensagem de log inicial
	logger.Info("Starting application...")

	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Conectar ao banco de dados
	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Usar a porta da aplicação definida nas variáveis de ambiente
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("A porta da aplicação (APP_PORT) não está definida no arquivo .env")
	}

	// Inicializar repositórios e serviços
	userRepo := repository.NewUserRepository()
	userService := service.NewUserDomainService(userRepo)
	userController := user.NewUserControllerInterface(userService)

	vagaRepo := repository.NewVagaRepository()
	vagaService := service.NewVagaDomainService(vagaRepo)
	vagaController := vaga.NewVagaControllerInterface(vagaService)

	veiculoRepo := repository.NewVeiculoRepository()
	veiculoService := service.NewVehicleDomainService(veiculoRepo)
	veiculoController := veiculo.NewVeiculoControllerInterface(veiculoService)

	// Configurar o servidor Gin
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Aplicar middleware de erros
	router.Use(middleware.ErrorHandlingMiddleware())

	// Inicializar rotas com as dependências
	api := router.Group("/api/v1")
	routes.InitRoutes(api, userController, vagaController, veiculoController)

	// Rodar a aplicação e tratar erro de inicialização
	if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
