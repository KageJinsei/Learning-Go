package main
import ("fmt")

func main() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", (myslice1))
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// com capacidade omitida
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", (myslice2))
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

/*
	Criando uma Slice com a função make() em Go:
		- A função "make()" em Go também pode ser usada para criar uma slice.
		- Syntax: slice_name := make([]type, length, capacity)

		Note: Se o parâmetro capacidade não estiver definido, será igual ao comprimento.

		Exemplo: Em line 4 mostra como criar slices usando a função "make()".
*/