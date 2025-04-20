package main

import "fmt"

func main() {

	var floatNumber float32 = 1.1 //Utiliza o espaço de 32 bits
	var pi float64 = 3.14         //Utiliza o espaço de 64 bits, mais utilizado
	var raio float64 = 2.5
	var area = pi * raio * raio

	// Não é possível calcular float32 com float64
	fmt.Println("floatNumber: ", floatNumber)
	fmt.Println("PI: ", pi)
	fmt.Println("raio: ", raio)
	fmt.Println("Área do círculo: ", area)

}
