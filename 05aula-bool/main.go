package main

import "fmt"

func main() {
	// var novotipo bool //Inicialmente o go assume que é false.
	// var novotipo bool = true // atribuindo o valor
	// fmt.Println(novotipo)

	/*tratando para saber se é maior ou menor do valor.*/
	var maior bool = 10 > 5
	fmt.Println("10 é maior que 5?", maior)
	var menor bool = 10 < 5
	fmt.Println("10 é maior que 5?", menor)

}
