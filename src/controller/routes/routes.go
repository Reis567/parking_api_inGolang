package routes

import (
    "github.com/gin-gonic/gin"
    "meu-novo-projeto/src/controller/user"
)

func InitRoutes(r *gin.RouterGroup) {
    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/", user.CreateUser)
        userRoutes.GET("/:id", user.FindUserByID)
        userRoutes.GET("/email/:email", user.FindUserByEmail)
        userRoutes.PUT("/:id", user.UpdateUser)
        userRoutes.DELETE("/:id", user.DeleteUser)
    }
}
