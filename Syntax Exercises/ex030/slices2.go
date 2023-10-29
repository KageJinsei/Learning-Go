package main
import ("fmt")

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}

/*
	Criando uma Slice de um Array em Go:
		- Você pode criar uma slice cortando um array:
			- var myarray = [length]datatype{values} // A array
			- myslice := myarray[start:end] // Uma slice feita da array
		
		Exemplo:
			- O exemplo em line 4 mostra como criar uma slice de uma array.
			- No exemplo citado acima, "myslice" é uma slice com 2 de comprimento. É feito com a array "arr1" que tem 6 de comprimento.
			A slice começa a partir do segundo elemento da array que tem valor 12. A slice pode crescer até o final da array. Isso significa que a capacidade da slice é 4.
			Se "myslice" for iniciada a partir do elemento 0, a capacidade da slice seria 6.

*/