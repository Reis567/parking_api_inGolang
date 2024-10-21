package user

import (
	"fmt"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
   var userRequest request.UserRequest
   if err := c.ShouldBindJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, erro=%s",err)
		)
		c.JSON(restErr.Code,restErr)
		return
   }
   fmt.Println(userRequest)
}
