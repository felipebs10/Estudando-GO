package main

import (
	"errors"
	"fmt"
)

func main() { //Inicio do main

	//Validando se nota está aprovado ou não
	/* nota := 75

	if nota >= 90 {
		fmt.Println("Aprovado com Estrelinhas", nota)
	} else if nota >= 70 {
		fmt.Println("Aprovado", nota)
	} else {
		fmt.Println("Reprovado", nota)
	}
	*/
	//Possibilidade de declarar uma variável dentro da expressão.
	if err := thisIsAnError(); err != nil { //Declarando uma váriavel que só funciona no escopo desse if, conforme informado muito comum no golang.
		fmt.Println(err.Error())
	}

} //Fim da função Main

// Possibilidade de declarar uma variável dentro da expressão.
func thisIsAnError() error { // A função tem que ficar fora da função do main.
	return errors.New("Isto é um erro")
}
