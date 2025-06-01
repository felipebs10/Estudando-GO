package controller

import (
	"net/http"

	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/logger"
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/validation"
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/controller/model/request"
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//Função antes
/*
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
*/
// Função atual
func CreateUser(c *gin.Context) {
	// log.Println("Init CreateUser controller") /*Como estava antes da função logger.info criada*/
	//Um exemplo para identificar qual rotina está sendo executada abaixo
	///logger.Info("Init CreateUser Controller ou Iniciando criação de usuário no controle, journey=CriacaoUsuario")
	//Exemplo mais detalhado do Professor mas não funcionou  dá erro.
	/*	logger.Info("Init CreateUser Controller ou Iniciando criação de usuário no controle",
		//Forma abaixo de passar os valores
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	) */

	//último exemplo funcionando para mostrar os logs no consolle.
	logger.Info("Init Create User controller",
		zap.String("journey", "createuser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// log.Printf("Error trying to marshal object, error=%s\n", err.Error()) /*como estava antes da função logger.info*/
		//Exemplo mais detalhado do Professor mas não funcionou  dá erro.
		/* logger.Error("Error trying to validate user info", err, zapcore.Field{
			Key:    "journey",
			String: "createUser",
		}) */

		//último exemplo funcionando para mostrar os logs no consolle.
		logger.Error("Error Trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// fmt.Println(userRequest) /*Trocado para o logger.info, para dizer que deu certo*/
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	//Exemplo mais detalhado do Professor mas não funcionou  dá erro.
	/*	logger.Info("User create successfuly ou Usuário criado",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	*/

	//último exemplo funcionando para mostrar os logs no consolle.
	logger.Info("Init Create Successfully",
		zap.String("journey", "createuser"),
	)
	c.JSON(http.StatusOK, response)

}
