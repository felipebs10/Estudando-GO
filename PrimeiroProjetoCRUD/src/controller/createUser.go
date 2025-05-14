package controller

import (
	"fmt"

	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/rest_err"
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError( //Tratando o erro
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))

		//Retorno do json
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(UserRequest)
}
