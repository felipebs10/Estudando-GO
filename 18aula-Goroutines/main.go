package main

import ("fmt"
		"time"
		)

func exibirMsg(){ //Goroutine
	fmt.Println("Olá de uma goroutine!")
}

/*Nesse exemplo, disparamos 3 processos, ou seja, 3 gourintes, ao mesmo 
tempo e eles estão consumindo o mesmo recurso do processador.*/
func main() {
	go exibirMsg() //Executando uma Goroutine 1 vez.
	go exibirMsg() //Executando uma Goroutine 2 vez.
	go exibirMsg() //Executando uma Goroutine 3 vez.
	time.Sleep(1 * time.Second) //Tem que esperar um segudo para mostrar a função.
	fmt.Println("Olá main function")
}
