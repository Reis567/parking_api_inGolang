package middleware

import (
    "net/http"
    "runtime/debug"

    "github.com/gin-gonic/gin"
    "meu-novo-projeto/src/configuration/rest_err"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                err := resterr.NewInternalServerError("Internal server error", nil)
                c.JSON(http.StatusInternalServerError, err)
                c.Abort()
                debug.PrintStack()
            }
        }()
        c.Next()
    }
}
