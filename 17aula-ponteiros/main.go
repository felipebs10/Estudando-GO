package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main(){

	var p1 Pessoa = Pessoa{Nome: "Felipe"}
	var p2 Pessoa = Pessoa{Nome: "Teste"}

	//Exemplo Passando o endereço(Ponteiro), utilizando o & (E comercial)
	var p3 *Pessoa = &p1
	//Mostrando o endereço(Ponteiro) onde a variável está armazenada.
	fmt.Println(&p1.Nome) 
	//Após passar o endereço de p1 para p3, fica igual para passar o endereço precisa do *
	fmt.Println(&p3.Nome) 
	fmt.Println(p1.Nome)
	p3.Nome = "Teste2" // Nesse momento estou mudando o valor da variável p1, pois o p3 tem o endereço.
	fmt.Println(p1.Nome)
	fmt.Println(p2)
	
}