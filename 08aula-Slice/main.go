package main

import "fmt"

func main() {

	var pessoas []string
	pessoas = append(pessoas, "Felipe", "Lis", "Teste", "Teste2")
	fmt.Println(pessoas)
	//Acessando o valor da posição
	fmt.Println(pessoas[2])

	//Verificando o tamanho do slice.
	fmt.Println(len(pessoas))

	//Operação de divisão de um slice, Monstrando os valores a partir de um intervalo.
	fmt.Println(pessoas[1:3]) //Caso não passe nada no primeiro índice, passa a primeira posição no caso, se não passar nada na segunda posição ele vai até o final

}
