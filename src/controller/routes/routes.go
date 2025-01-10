package routes

import (
	"meu-novo-projeto/src/controller/registro" // Import para o controlador de registro
	"meu-novo-projeto/src/controller/user"
	"meu-novo-projeto/src/controller/vaga"
	"meu-novo-projeto/src/controller/veiculo"
	"meu-novo-projeto/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController user.UserControllerInterface,
	vagaController vaga.VagaControllerInterface,
	veiculoController veiculo.VeiculoControllerInterface,
	registroController registro.RegistroControllerInterface, // Novo parâmetro
) {
	// Rotas de usuário
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/login", userController.LoginUser)

		// Aplica o middleware AuthMiddleware às rotas protegidas
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
		// Aplica o middleware AuthMiddleware às rotas protegidas de vagas
		vagaRoutes.Use(middleware.AuthMiddleware())
		{
			vagaRoutes.POST("/", vagaController.CreateVaga)
			vagaRoutes.GET("/:id", vagaController.FindVagaByID)
			vagaRoutes.GET("/", vagaController.FindAllVagas)
			vagaRoutes.PUT("/:id", vagaController.UpdateVaga)
			vagaRoutes.DELETE("/:id", vagaController.DeleteVaga)
		}
	}

	// Rotas de veículo
	vehicleRoutes := r.Group("/veiculos")
	{
		// Aplica o middleware AuthMiddleware às rotas protegidas de veículos
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
		// Aplica o middleware AuthMiddleware às rotas protegidas de registros
		registroRoutes.Use(middleware.AuthMiddleware())
		{
			registroRoutes.POST("/", registroController.CreateRegistro)
			registroRoutes.GET("/:id", registroController.FindRegistroByID)
			registroRoutes.GET("/", registroController.FindAllRegistros)
			registroRoutes.PUT("/:id", registroController.UpdateRegistro)
			registroRoutes.DELETE("/:id", registroController.DeleteRegistro)
		}
	}
}
