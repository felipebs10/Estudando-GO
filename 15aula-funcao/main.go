package main	

import "fmt"

func main(){
	//Função com retorno
	/* resultado := soma(3, 2)
	fmt.Println(resultado) 
	*/
	
	//Função sem o retorno
	// soma(3, 2)

	//Exemplo de atribuindo o retorno de uma função a uma variável.
	var fixo = 4
	multiplica := func(x int) int {
		return x * fixo
	}

	resultado := multiplica(5)
	fmt.Println(resultado)

}

//Exemplo de chamada de outra função completa
//func soma(a, b int, nome string, x Cliente) int {

func soma(a, b int) /*int*/ { //caso não tenha o tipo do retorno, se é inteiro ou string, é uma função void, que não retorna ada.
	//Atenção Soma, com o S maiusculo, é possível outros pacotes acessar, se for minusculo, exemplo soma, é um função privada desse pacote.
	// return a + b
	fmt.Println(a + b)
}