package main

import "fmt"

func main() {
	var pessoas = map[string]int{}

	pessoas["lais"] = 26 //Chave
	pessoas["leo"] = 32  //Valor

	fmt.Println(pessoas)
	fmt.Println(pessoas["lais"])

	//Validando se existe uma chave para mostrar
	//if idade, ok := pessoas["fabio"]; ok {
	if idade, ok := pessoas["leo"]; ok {
		fmt.Println("Pessoa existe no map:", idade, ok)
	} else {
		fmt.Println("Pessoa NÃO existe no map")
	}

	//Remoção de um item do map.
	/* delete(pessoas, "leo")
	fmt.Println(pessoas)
	*/
}
