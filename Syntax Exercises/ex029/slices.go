package main
import ("fmt")

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

/*
	Slices em Go:
		- As slices são semelhantes às arrays, mas são mais poderosas e flexíveis.
		- Como arrays, slices também são usadas para armazenar vários valores do mesmo tipo em uma única variável. No entanto, ao contrário dos arrays, o comprimento de uma slice pode crescer e encolher como você achar melhor.
		- Em Go, existem várias maneiras de criar uma slice:
			- Usando o formato: []data type{values};
			- Criando uma slice de um array;
			- Usando a função "make()".

	Criando uma slice com "[]data type{values}":
		- Syntax: slice_name := []datatype{values}
		- Uma maneira comum de declarar uma slice é assim:
			- myslice := []int{}
		- O código acima declara uma slice vazia de 0 comprimento e 0 capacidade. Para inicializar uma slice durante a declaração, deverá usar:
			- myslice:= []int{1, 2, 3}

		- O código acima declara uma slice de inteiros com 3 de comprimento e 3 de capacidade.
		- Em Go, existem duas funções que podem ser usadas para retornar o comprimento e a capacidade de uma slice:
			- Função "len()" | retorna o comprimento da slice (número de elementos na slice)
			- Função "cap()" | retorna a capacidade da slice (número de elementos que a slice pode aumentar ou diminuir)

		Exemplo em: line 4
			- No exemplo, vemos que na primeira slice (myslice1), os elementos reais não são especificados, tanto o comprimento quanto a capacidade da slice será zero. Na segunda slice (myslice2), os elementos são especificados, e tanto o comprimento quanto a capacidade são iguais ao número de elementos reais especificados
*/