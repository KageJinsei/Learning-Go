package main
import ("fmt")

func main() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1,2,3,4,5,6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}

/*
	Encontrando o comprimento de um Array em Go:
		- A função "len()" é usada para encontrar o comprimento de um array. Basicamente, ele retorna em número inteiro a quantidade de elementos dentro de um array (ex: linhe 8 and 9).
*/