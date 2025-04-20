package main

import "fmt"


//Obs.: O underline no for é ocultar, porque no go não deixa executar se não estiver usando a variável.
func main(){
	//exemplo de um slice.
	//noms := []string{"lais","joao", "nath"} 

	//Exemplo de map
	users := map[string]string{
		"nome": "João",
		"idade":"30",
	}

	//Verificando todos os dados do map
    for key, value := range users {
		fmt.Println(key, value)
	}

	//Verificando apenas o valor
	for _, value := range users {
		fmt.Println(value)
	}

	//Verificando apenas a chave
	for key, _ := range users {
		fmt.Println(key)
	}
}