// Outro exemplo de if e else com map
package main

import "fmt"

func main() {
	players := map[string]int{
		"Teste01": 26,
		"Teste02": 30,
		"Teste03": 35,
	}

	if value, ok := players["Teste01"]; ok { //Expressão se existe
		fmt.Println("Pontos:", value, ok)
	}
	value, ok := players["Teste04"] //Mostrando os valores quando não tem o dado no map
	fmt.Println("Pontos:", value, ok)

}
