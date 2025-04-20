package main

import "fmt"

// - **int8:  ** Inteiro de  8 bits (valores de -128 a 127).
// - **int16: ** Inteiro de 16 bits (valores de -32768 a 3267).
// - **int32  ** Inteiro de 32 bits (valores de -2147483648 a 214783647).
// - **int64: ** Inteiro de 64 bits (valores de -9223372036854775808 a 9223372036854775807).

func main() {
	var idade int = 37 //Armazena o tamanho do sistema, no meu caso 64
	//var contador int32 = 2 //Armazena 32 bits na memória
	// var intice int8 = 1    //Armazena apenas 8 bits

	fmt.Println("Idade: ", idade)
}
