package main 

import ("fmt"
		"time")

func main(){
	//Declarando um channel que vai ler um inteiro.
	ch := make (chan int, 3) //esse 3 configura par que seja possível escrever 3 vezes, caso tenha uma quarta vez vai esperar liberar para escrever. 

	//Escrevendo a seta fica do lado direito (exemplo: ch <-) e se quer ler fica do lado esquerdo(exemplo: <-ch)
	go func(){ 
		//Exemplo de atribuição de valor simpes no channel
		/*	ch <- 3
		ch <- 2
		ch <- 1*/ 

		//Exemplo de atribuição de valor com for no channel
		for i:= 0; i < 5; i++{
			ch <- i
		}
		close(ch) //Ao abrir um channel, deve ser fechado para não causar o deadlock.
		//Não apresentou esse print, 
		fmt.Println("Escrita Finalizada")
	}()

	time.Sleep(time.Second * 1)
	/*1º Exemplo */
	/*<-ch //Retirando o valor mostrar que vai pegar o 3
	<-ch
	valor := <-ch
	fmt.Println("Valor do canal:", valor) */
	/* 2ª Exemplo com For*/ 
	for valor := range ch {
		fmt.Println("Leitura", valor)
	}

}