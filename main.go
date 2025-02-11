package main

import (
	"fmt"
	"log"
	"os"

	"meu-novo-projeto/src/configuration/database"
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/controller/agendamento" // Import para o controlador de agendamentos
	"meu-novo-projeto/src/controller/registro"
	"meu-novo-projeto/src/controller/routes"
	"meu-novo-projeto/src/controller/user"
	"meu-novo-projeto/src/controller/vaga"
	"meu-novo-projeto/src/controller/veiculo"
	"meu-novo-projeto/src/controller/relatorios"
	"meu-novo-projeto/src/middleware"
	"meu-novo-projeto/src/model/repository"
	"meu-novo-projeto/src/model/service"
	"meu-novo-projeto/src/controller/calendario"
	"meu-novo-projeto/src/controller/pagamento"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting application...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

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

	registroRepo := repository.NewRegistroEstacionamentoRepository()
	registroService := service.NewRegistroEstacionamentoDomainService(registroRepo,vagaRepo)
	registroController := registro.NewRegistroControllerInterface(registroService)

	agendamentoRepo := repository.NewAgendamentoRepository() // Novo repositório
	agendamentoService := service.NewAgendamentoDomainService(agendamentoRepo,vagaRepo,registroRepo) // Novo serviço
	agendamentoController := agendamento.NewAgendamentoControllerInterface(agendamentoService, vagaService,registroService)

	relatoriosService := service.NewRelatoriosService(registroRepo, vagaRepo)
	relatoriosController := relatorios.NewRelatoriosController(relatoriosService)


	calendarioController := calendario.NewCalendarioController(registroService)

	pagamentoRepo := repository.NewPagamentoRepository()
	pagamentoService := service.NewPagamentoDomainService(pagamentoRepo)
	pagamentoController := pagamento.NewPagamentoControllerInterface(pagamentoService)


	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(middleware.ErrorHandlingMiddleware())

	api := router.Group("/api/v1")
	routes.InitRoutes(api,
		 userController,
		  vagaController,
		   veiculoController,
		    registroController,
			 agendamentoController,
			  relatoriosController,
			  calendarioController,
			  pagamentoController,

			)

	if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
