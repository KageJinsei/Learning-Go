package main
import ("fmt")

func main() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5] // Slice array (uma fatia de uma array)

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Alterando o compriento fatiando a array novamente

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Alterando o comprimento anexando itens

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

/*
	Alterar o comprimento de uma Slice em Go:
		- Ao contrário das arrays, é possível alterar o comprimento de uma slice.

		Exemplo:
			- Em line 4, inicia um exemplo de como alterar o comprimento de uma slice.
*/