package main
import ("fmt")

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var arr3 = [...]int{9, 9, 9}
	arr4 := [...]int{0, 0, 0, 0, 0}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}

/*
	Arrays (Matriz) em Go:
		- Arrays são utilizados para armazenar vários valores do mesmo tipo em uma única variável, isso ajuda na otimização do código pois não será necessário declarar variáveis separadas para cada valor.

	Declarando Array em Go:
		- Em Go, existem duas maneiras de declarar um Array:
			com "var":
				- array_name = [length]datatype{values} // aqui, o comprimento é definido (onde está length)
				- var array_name = [...]datatype{values} // aqui, o comprimento é inferido(deduzido pelo compilador)
			Obs: Exemplo de como declarar um array utilizando "var" em: line 5
			Obs: Exemplo de como declarar um array com comprimento inferido utilizando "var" em: line 8

			com o sinal ":=":
				- array_name :=  [length]datatype{values} // aqui, o comprimento é definido (onde está length)
				- array_name := [...]datatype{values} // aqui, o comprimento é inferido(deduzido pelo compilador)
			Obs: Exemplo de como declarar um array utilizando ":=" em: line 6
			Obs: Exemplo de como declarar um array com comprimento inferido utilizando ":=" em: line 9
			

	Nota: O comprimento especifica o número de elementos que serão armazenados no array. Em Go, arrays tem um comprimento fixo. O comprimeiro do array é definido por um número ou é inferido (significa que o compilador decide o comprimento do array com base no número de valores inseridos).
*/