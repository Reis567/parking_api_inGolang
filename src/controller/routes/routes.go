package routes

import (
	"meu-novo-projeto/src/controller/user"
	"meu-novo-projeto/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController user.UserControllerInterface) {
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
}
