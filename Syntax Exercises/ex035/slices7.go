package main
import ("fmt")

func main() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
}

/*
	Anexar uma Slice a outra Slice em Go:
		- É possível anexar todos os elementos de uma slice a outra slice usando a função "append()".

		Syntax: slice3 := append(slice1, slice2...)

		Note: O '...' depois de slice2 é necessário ao anexar os elementos de uma slice a outra.

		Exemplo:
			- Em line 7, mostra como anexar uma slice a outra slice.
*/