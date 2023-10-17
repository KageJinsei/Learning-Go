package main
import ("fmt")

func main() {
	var a bool = true // Boolean
	var b int = 5 // Integer
	var c float32 = 3.14 // Float point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}

/*
	Tipos de dados em Go:
		- O tipo de dados é um conceito importante na programação. O tipo de dados especifica o tamanho e o tipo de valores que as variáveis podem receber.
		- Go é uma linguagem estaticamente tipada, o que significa que uma vez que um tipo de variável é definido, ele só pode armazenar dados desse tipo.
	
	Go tem três tipos de dados básicos:
		bool: representa um valor booleano (verdadeiro ou falso);
		numeric: representa números inteiros, valores de ponto flutuante, tipos complexos;
		string: representa um valor de cadeia de caracteres.
*/