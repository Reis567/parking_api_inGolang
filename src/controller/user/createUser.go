package user

import (
    "fmt"
    "meu-novo-projeto/src/configuration/rest_err"
    "meu-novo-projeto/src/controller/model/request"
    "github.com/go-playground/validator/v10"

    "github.com/gin-gonic/gin"
)

// Instancia do validador
var validate = validator.New()

// CreateUser é responsável por criar um novo usuário
func CreateUser(c *gin.Context) {
    var userRequest request.UserRequest
    // Tentar fazer o binding do JSON para o struct UserRequest
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        restErr := rest_err.NewBadRequestError(
            fmt.Sprintf("Existem campos incorretos, erro=%s", err),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

    // Validar os dados usando o validator
    if err := validate.Struct(userRequest); err != nil {
        restErr := rest_err.NewBadRequestError(
            fmt.Sprintf("Erro de validação: %s", err),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

    // Simulação de lógica para criar o usuário
    fmt.Println(userRequest)

    // Retornar resposta de sucesso
    c.JSON(201, gin.H{
        "message": "Usuário criado com sucesso!",
        "user":    userRequest,
    })
}
