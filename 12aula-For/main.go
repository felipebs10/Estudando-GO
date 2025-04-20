package main

import "fmt"

func main() {
	//Exemplo 01  de for
	/* sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i //Acumando os valores do contador.
	}

	fmt.Println("O valor da soma do contador é:", sum)
	*/
	/*******************************************************************/
	//Exemplo 02 de for, como se fosse um while.
	//For no formato de while.
	/* sum := 0
	for sum < 20 {
		sum += 2
		fmt.Println("Dentro do For/While:", sum)
	}
	fmt.Println("Fora do For/While", sum) */
	/*******************************************************************/
	//Exemplo 03 for em um slice, ou seja lendo a estrutura de um array, ou slice.
	//Famos percorer o slice, ou array.
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}
