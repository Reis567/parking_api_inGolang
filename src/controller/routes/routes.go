package routes

import (
    "meu-novo-projeto/src/controller/user"
    "meu-novo-projeto/src/model/service"
    "github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
    // Instancia o serviço e o controlador
    userService := service.NewUserDomainService()
    userController := user.NewUserControllerInterface(userService)

    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/", userController.CreateUser)
        userRoutes.GET("/:id", userController.FindUserByID)
        userRoutes.GET("/email/:email", userController.FindUserByEmail)
        userRoutes.PUT("/:id", userController.UpdateUser)
        userRoutes.DELETE("/:id", userController.DeleteUser)
    }
}
