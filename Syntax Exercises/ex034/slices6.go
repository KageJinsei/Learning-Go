package main
import ("fmt")

func main() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

/*
	Anexar elementos a uma Slice em Go:
		- É possível anexar elementos ao final de uma slice usando a função "append()".

		Syntax: slice_name := append(slice_name, element1, element2, ...)

		Exemplo:
			- Em line 10, mostra como anexar elementos ao final de uma slice.
*/