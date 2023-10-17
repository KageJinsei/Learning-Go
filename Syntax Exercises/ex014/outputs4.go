package main
import ("fmt")

func main() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}

/*
	Formatação de verbos para Printf() em Go:
		- Go oferece vários verbos de formatação que podem ser usados com a função "Printf()".
	
	Verbos de formatação geral:
		- Os seguintes verbos podem ser usados com todos os tipos de dados:

			Verb		Description
			%v 	   -	Imprime o valor no formato padrão
			%#v    -	Imprime o valor no formato Go-syntax
			%T 	   -	Imprime o tipo do valor
			%% 	   - 	Imprime o sinal "%"
*/