package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

/*Dois exemplos para os métodos*/
/*Nesse formato recebe uma cópia, e não muda o valor original*/
 func (p Pessoa) Apresentar(){
	p.Nome = "Joana"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

/*Mese formato com *(asterisco), ele passa como referência e permite alterar o valor original.*/
/*func (p *Pessoa) Apresentar(){
	p.Nome = "Joana"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}*/ 
func main(){
	p1 := Pessoa{Nome: "Felipe", Idade: 37}
	p2 := Pessoa{Nome: "Teste", Idade: 22}
	p1.Apresentar()
	p2.Apresentar()
	fmt.Println(p1.Nome)
	fmt.Println(p2.Nome)
}