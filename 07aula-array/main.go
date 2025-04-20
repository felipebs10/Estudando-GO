package main

import "fmt"

func main() {

	//O espaço de um array no go, já é definido, se for mair, que o definido, ele não deixa nem compilar.
	// Lembrando que o array e variaveis são case sensitive.

	var Pessoas [2]string
	Pessoas[0] = "Felipe"
	Pessoas[1] = "Lis"
	// Pessoas[2] = "Teste" // Não deixa executar porque o tamanho do do array.
	fmt.Println(Pessoas)    // Mostra dos os valores do array
	fmt.Println(Pessoas[0]) // Mostra o valor da primeira posição do array.

}
