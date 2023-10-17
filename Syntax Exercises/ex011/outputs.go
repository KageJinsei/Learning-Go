package main
import ("fmt")

func main() {
	var i, j string = "Hello", "World"
	var k, l string = "Olá", "Mundo"
	var m, n = 10, 20

	// imprimindo os valores das variáveis "i" e "j" separadamente e de forma padrão:
	fmt.Print(i)
	fmt.Print(j)

	// imprimindo os valores das variáveis "k" e "l" separadamente e com quebra de linha (\n):
	fmt.Print("\n", k, "\n")
	fmt.Print(l, "\n")

	// imprimindo múltiplos valores (variáveis "i" e "j") numa função e com quebra de linha (\n):
	fmt.Print(i, "\n", j, "\n")

	// imprimindo múltiplos valores (variáveis "i" e "j") numa função com espaçamento no meio e com quebra de linha (\n):
	fmt.Print(i, " ", j, "\n")

	// imprimindo múltiplos valores na tela (variáveis "m" e "n")
	fmt.Print(m,n)
}

/*
	A função Print() em Go:
		- A função "Print()"" imprime seus argumentos com seu formato padrão (ex: line 10 and 11).
		- Se quisermos imprimir os argumentos com quebra de linha, é preciso utilizar o "\n" (ex: line 14 and 15).
		- Com apenas uma função "Print()" é possível imprimir múltiplas variáveis (ex: line 18).
		- Se quisermos adicionar um espaço entre os argumentos de string, precisamos usar " " (ex: line 21).
		- A função "Print()" insere um espaço entre os argumentos caso não sejam strings (ex: line 24).
		Obs: "\n" cria novas linhas.
*/