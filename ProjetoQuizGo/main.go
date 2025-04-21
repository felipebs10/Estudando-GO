package main

import ("fmt"
		"os"
		"bufio"
		"encoding/csv"
		"strconv"
		"errors"
		)


type Question struct {
	Text    string
	Options  []string
	Answer int 
}

type GameState struct {
	Name string
	Points int
	Questions []Question
}

func (g *GameState) ProccessaCSV(){
	//Abrindo o arquivo em go.
	f, err := os.Open("quizz.csv") 
	if err != nil {
		panic("Erro ao ler arquivo")
	}
	
	//Fechando o arquivo, deve ser fechado, porque ocupa muito espaço em memória.
	//defer, é uma funcionalidade do GO para depois que concluir todas as funções programadas, ele vai executar aquela por último, idependente da parte do código.
	defer f.Close()

	reader := csv.NewReader(f) //Lendo o Arquivo CSV
	records, err := reader.ReadAll() //Lendo todas as linhas do CSV.
	if err != nil {
		panic ("Error ao ler CSV")
	}

	for index, record := range records {
		//fmt.Println(index, record) //Foi utilizado o index, que é uma estrutura de array, mas antes foi utilizado o underline (_), que ai informa como não precisa.
		if index > 0 {//Primeira linha é o cabeçalho, por isso da primeira posição para frente.
			correctAnwser, _ := toInt(record[5])
			question := Question {
				Text: record[0],
				Options: record[1:5], //No slice, o final sempre subtrai 1 
				Answer: correctAnwser, //No meu não apresentou erro de conversão, mas fiz mesmo assim, depois foi adaptado devido ao usuário poder escolher uma resposta errda.
			}
			g.Questions = append(g.Questions, question) //Adicionando um intem no slice com append.	
		}
	}
}

func toInt(s string) (int, error){
	i, err := strconv.Atoi(s) //Pacote utilizado para converter um texto para inteiro
	if err != nil {
		return 0, errors.New("Não é permitido caracteres diferente de número, por favor insira um número")
	}

	return i, nil
}

func (g *GameState) Init(){
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome:")
	reader := bufio.NewReader(os.Stdin) //OS é o pacote, o Stdin é a entrada, o Stdout é a saída.
	name, err  := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a String")
	}

	g.Name = name
	fmt.Println("Vamos ao jogo", g.Name) //Não precisei colocar o %s, funcionou sem.
}



func (g *GameState) Run() {
	//Exibir a pergunta pro usuário
	for index, question := range g.Questions {
		fmt.Println("\033[33m", index+1, question.Text ,"\033[0m")
		for j, option := range question.Options {
			fmt.Println("[",j+1,"]", option)
		}

		fmt.Println("Digite uma alternativa: ")

		var answer int
		var err error

			for {
				reader := bufio.NewReader(os.Stdin)
				read, _ := reader.ReadString('\n')
			
				answer, err = toInt(read[:len(read)-2]) //Tive de alterar para -2, pois estava caindo no erro.
				if err != nil {
					fmt.Println(err.Error())
					continue // Vai começar o for novamente, até ser informada a opção certa.
				}
				break // Comando para sair do for quando digitar a opção correta.
			}
			
		
		if answer == question.Answer {
			fmt.Println("Parabéns você acertou!!")
			g.Points +=  10
		}else {
			fmt.Println("Ops ! Errou!" )
			fmt.Println("---------------------")
		}
	}
}

func main(){
	game := &GameState{}

	go game.ProccessaCSV()

	game.Init()
	game.Run()

	fmt.Println("Fim de jogo, você fez ", game.Points, "pontos")
	
}


