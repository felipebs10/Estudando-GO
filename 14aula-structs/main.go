package main 

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco //Struct dentro de uma outr struct
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

func main() {
	cliente1 := Cliente{
		Nome: "Felipe",
		Idade: 37,
		Endereco: Endereco { //Inserido dados de um struct que será utilizado dentro de outro struct. Não é obrigatório ter todas os dados preenchidos.
			Rua: "Rua do teste",
			Numero: 123,
			Cep: "88800-000",
			Estado: "SP",
		},
	}

	 cliente2 := Cliente{
		Nome: "Teste01",
		Idade: 25,
	} 

	cliente2.Email = "teste@gmail.com" // Alteração da struct do cliente2, que não tinha e-mail ainda.
	cliente1.Endereco.Numero = 124

	fmt.Println(cliente1)
	fmt.Println(cliente1.Endereco.Cep)
	fmt.Println(cliente2)
	fmt.Println(cliente1.Nome)
	fmt.Println(cliente2.Idade)
	fmt.Println(cliente2.Email)
}