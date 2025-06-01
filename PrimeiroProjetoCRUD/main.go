package main

import (
	"log"

	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/logger"
	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//log.Println("Está começando") /*Exemplo mostrando como adicionar a mensagem no log*/
	logger.Info("About to start application ou Começando a usar a aplicação.") /*Exemplo mostrando como aplicar a função criada para mais nível */
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup) //Passando o ponteiro, o local onde está a armazenado o valor.

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	/*A instrução router := gin.New() criará um novo roteador gin. Os
	roteadores podem ser inicializados de duas maneiras, uma usando gin.New()
	e a outra usando gin.Default().

	A diferença é que gin.New() inicializa um roteador sem nenhum
	middleware enquanto gin.Default() inicializa o roteador com logger
	e middlewares de recovery.
	*/

}
