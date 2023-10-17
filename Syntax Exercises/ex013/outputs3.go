package main
import ("fmt")

func main() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i tem o valor: '%v' e o tipo: %T\n", i, i)
	fmt.Printf("j tem o valor: '%v' e o tipo: %T\n", j, j)
}

/*
	A função Printf() em Go:
		- A função "Printf()" primeiro formata o seu argumento com base no verbo de formatação fornecido (ex: %v) e depois os imprime.
	
	Foram utilizados dois verbos de formatação:
		- O verbo %v: é usado para imprimir o valor dos argumentos (ex: imprimir o valor que "i" recebe).
		- O verbo %T: é usado para imprimir o tipo dos argumentos (ex: imprimir o tipo de dado que "j" recebeu ao declarar a variável).
*/