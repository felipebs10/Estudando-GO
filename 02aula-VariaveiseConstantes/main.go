package main

import "fmt"

func main() {

	/*Exemplo de variável em go*/
	//Passanto com o tipo da váriavel
	//var texto string = "Olá"

	//Não passando o tipo de váriavel
	//var texto = "Oi"

	//Atribuição
	// texto = "Tchau"

	//Passando um inteiro
	//var texto = 123

	//Passando inteiro com operação.
	//var texto = 1 + 1

	//Caso não seja inicializado, a ariável será inicializada no go, se for string com vazio, inteiro com zero e boolean com false.

	// fmt.Println(texto)

	//Fazendo operação  ou concatenando o texto
	/* texto := "Ola"
	texto += " Tudo bem"
	fmt.Println(texto) */
	/************************************************************************************************************************************************/
	// Exemplo de Constantes em go

	// const nome string = "Felipe" //Ela também declarou antes da função do main.
	// nome = "Teste" // Ele já deu erro, o compilador não deixa alterar e compilar o programa.

	//Declaração mais curta sem informar o tipo, ele vai pegar conforme for atribuído, igual a constante.
	const nome = "Felipe"

	fmt.Println(nome)

}
