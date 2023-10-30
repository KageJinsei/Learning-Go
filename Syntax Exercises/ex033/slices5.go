package main
import ("fmt")

func main() {
	prices := []int{10, 20, 30}
	prices[2] = 50

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

/*
	Alterar elementos de uma Slice em Go:
		- É possível alterar um elemento específico em uma slice referindo-se ao número do índice.

		Exemplo:
			- Em line 6, mostra como alterar o terceiro elemento na slice "prices".
*/