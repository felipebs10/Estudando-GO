package controller

import (
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamnou a rota de forma errada")
	c.JSON(err.Code, err) //Retornando o erro.
}
