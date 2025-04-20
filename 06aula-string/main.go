package main

import (
	"fmt"
	"strings"
)

func main() {

	var hello string = "Olá Mundo!"
	var question string = "Como vai?"

	//Contatenação
	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.Contains(meet, "Mundo")) //Função para verificar se existe a palavra Mundo, é case sensitive, palavra tem que ser exatamente.
	fmt.Println(strings.ToUpper(meet))           //Função para converter toda a variável para maiusculo.

}
