package routes

import (
	"meu-novo-projeto/src/controller/agendamento" // Import para o controlador de agendamentos
	"meu-novo-projeto/src/controller/registro"
	"meu-novo-projeto/src/controller/user"
	"meu-novo-projeto/src/controller/vaga"
	"meu-novo-projeto/src/controller/veiculo"
	"meu-novo-projeto/src/controller/relatorios"
	"meu-novo-projeto/src/controller/calendario"
	"meu-novo-projeto/src/controller/pagamento"
	"meu-novo-projeto/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController user.UserControllerInterface,
	vagaController vaga.VagaControllerInterface,
	veiculoController veiculo.VeiculoControllerInterface,
	registroController registro.RegistroControllerInterface,
	agendamentoController agendamento.AgendamentoControllerInterface,
	relatoriosController relatorios.RelatoriosController,
	calendarioController calendario.CalendarioControllerInterface,
	pagamentoController pagamento.PagamentoControllerInterface,
) {
	// Rotas de usuário
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/login", userController.LoginUser)

		userRoutes.Use(middleware.AuthMiddleware())
		{
			userRoutes.GET("/:id", userController.FindUserByID)
			userRoutes.GET("/email/:email", userController.FindUserByEmail)
			userRoutes.PUT("/:id", userController.UpdateUser)
			userRoutes.DELETE("/:id", userController.DeleteUser)
		}
	}

	// Rotas de vaga
	vagaRoutes := r.Group("/vagas")
	{
		vagaRoutes.Use(middleware.AuthMiddleware())
		{
			vagaRoutes.POST("/", vagaController.CreateVaga)
			vagaRoutes.GET("/:id", vagaController.FindVagaByID)
			vagaRoutes.GET("/", vagaController.FindAllVagas)
			vagaRoutes.PUT("/:id", vagaController.UpdateVaga)
			vagaRoutes.DELETE("/:id", vagaController.DeleteVaga)
			vagaRoutes.GET("/disponiveis", vagaController.BuscarVagasDisponiveis)
		}
	}

	// Rotas de veículo
	vehicleRoutes := r.Group("/veiculos")
	{
		vehicleRoutes.Use(middleware.AuthMiddleware())
		{
			vehicleRoutes.POST("/", veiculoController.CreateVeiculo)
			vehicleRoutes.GET("/:id", veiculoController.FindVeiculoByID)
			vehicleRoutes.GET("/", veiculoController.FindAllVeiculos)
			vehicleRoutes.PUT("/:id", veiculoController.UpdateVeiculo)
			vehicleRoutes.DELETE("/:id", veiculoController.DeleteVeiculo)
		}
	}

	// Rotas de registro de estacionamento
	registroRoutes := r.Group("/registros")
	{
		registroRoutes.Use(middleware.AuthMiddleware())
		{
			registroRoutes.POST("/", registroController.CreateRegistro)
			registroRoutes.GET("/:id", registroController.FindRegistroByID)
			registroRoutes.GET("/", registroController.FindAllRegistros)
			registroRoutes.PUT("/:id", registroController.UpdateRegistro)
			registroRoutes.DELETE("/:id", registroController.DeleteRegistro)
		}
	}

	// Rotas de agendamentos
	agendamentoRoutes := r.Group("/agendamentos")
	{
		agendamentoRoutes.Use(middleware.AuthMiddleware())
		{
			agendamentoRoutes.POST("/", agendamentoController.CreateAgendamento)
			agendamentoRoutes.GET("/:id", agendamentoController.FindAgendamentoByID)
			agendamentoRoutes.GET("/", agendamentoController.FindAllAgendamentos)
			agendamentoRoutes.PUT("/:id", agendamentoController.UpdateAgendamento)
			agendamentoRoutes.DELETE("/:id", agendamentoController.DeleteAgendamento)
			agendamentoRoutes.POST("/entrada", agendamentoController.RegistrarEntrada)
			agendamentoRoutes.POST("/saida/:registroID", agendamentoController.FinalizarEstacionamento)
		}
	}
	relatoriosRoutes := r.Group("/relatorios")
	{
		relatoriosRoutes.GET("/financeiro", relatoriosController.CalcularReceita)
		relatoriosRoutes.GET("/ocupacao", relatoriosController.CalcularOcupacaoAtual)
		relatoriosRoutes.GET("/veiculos", relatoriosController.VeiculosMaisFrequentes)
		relatoriosRoutes.GET("/lotacao", relatoriosController.CalcularLotacao)
	}
	calendarioRoutes := r.Group("/calendario")
	{
		calendarioRoutes.GET("/:data", calendarioController.ListarRegistrosPorData)
	}
	pagamentoRoutes := r.Group("/pagamento")
	{
		pagamentoRoutes.POST("/", pagamentoController.CreatePagamento)
		pagamentoRoutes.PUT("/:id", pagamentoController.UpdatePagamento)

	}
}
