package main
import ("fmt")

func main() {
	var i = true
	var j = false

	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)
}

/*
	Verbos de formatação booleana em Go:
		- O seguinte verbo pode ser usado com o tipo de dados booleado:
			Verb		Description
			%t 	   -	Mostra o valor do operador booleano no formato verdadeiro ou falso (o mesmo que usar %v)
*/