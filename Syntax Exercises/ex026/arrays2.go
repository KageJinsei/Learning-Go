package main
import ("fmt")

func main() {
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	prices[2] = 50

	fmt.Println(prices)
}

/*
	Acessando elementos em um Array em Go:
		- É possível acessar um elemento dentro de um array específico referindo-se ao número do indice.
		Em Go, os índices nos arrays começam em 0. Isso significa que [0] é o primeiro elemento, [1] é o segundo e etc.
		- Em line 7 and 8 é possível ver um exemplo de acesso à um elemento através do índice no array que está armazenado na variável "prices" (line 5).

	Alterando Elementos em um Array em Go:
		- É possível alterar o valor de um elemento específico dentro de um array, referindo-se ao número do índice.
		O exemplo em line 10 mostra como alterar o valor do terceiro elemento no array "prices", e mostra na tela com a função "Println" em line 12.
*/