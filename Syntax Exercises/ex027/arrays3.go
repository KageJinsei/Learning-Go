package main
import ("fmt")

func main() {
	arr1 := [5]int{} // não inicializado
	arr2 := [5]int{1, 2} // parcialmente inicializado
	arr3 := [5]int{1, 2, 3, 4, 5} // completamente inicializado

	arr4 := [5]int{1:10,2:40}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}

/*
	Inicialização de Array em Go:
		- Se um array ou um de seus elementos não tiver sido inicializado no código, ele é atribuído o valor padrão do seu tipo (ex: valor padrão para "int" é "0" e o valor pardrão para string é "").

	Inicializar apenas elementos específicos em Go:
		- É possível inicializar apenas elementos específicos em um array (ex: em line 9, inicializa apenas o segundo e o terceiro elemento do array).

		Exemplo explicado:
			- O array em line 9 tem 5 elementos:
				- 1:10 significa: atribuir 10 ao índice 1 do array (segundo elemento).
				- 2:40 significa: atribuir 40 ao índice 2 do array (terceiro elemento).
*/