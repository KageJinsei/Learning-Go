package main
import ("fmt")

func main() {
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

/*
	Acessar elementos de uma Slice em Go:
		- Você pode acessar um elemento de uma slice específica referindo-se ao número do índice.
		- Em Go, os índices começam em 0. Isso significa que [0] é o primeiro elemento, [1] é o segundo elemento e etc..

		Exemplo:
			- Em line 7 and 8, é mostrado como acessar o primeiro e o terceiro elemento da slice "prices".
*/