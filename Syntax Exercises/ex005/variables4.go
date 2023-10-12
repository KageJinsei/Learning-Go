package main
import ("fmt")

var a int
var b int = 2
var c = 3
func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
	Diferença entre "var" e ":=":
		- A palavra-chave "var" pode ser utilizada dentro e fora das funções.
		- A declaração de variáveis e a atribuição de valores podem ser feitas separadamente (ex: line 5 and line 6).

		- O operador ":=" pode apenas ser utilizado dentro de funções.
		- A declaração de variáveis e a atribuição de valores não podem ser feitas separadamente (devem ser feitas na mesma linha).

	Obs: O exemplo do código acima mostra a declaração de variáveis fora de uma função com a palavra-chave "var".
*/