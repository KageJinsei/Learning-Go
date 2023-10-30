package main
import ("fmt")

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // Original slice

	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Criando uma cópia apenas com os números necessários
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

/*
	Eficiência de Memória em Go:
		- Ao usar slices, o Go carrega todos os elementos subjacentes na memória. Se a matriz for grande e você precisa de apenas alguns elementos, é melhor copiar esses elementos usando a função "copy()".
		- A função "copy()" cria uma nova matriz subjacente com apenas os elementos necessários para a slice. Isso reduzirá a memória usada para o programa.

		Syntax: copy(dest, src)

		- A função "copy()" recebe duas slices dest e src, e copia os dados de src para dest. Ela retorna o número de elementos copiados.

		Exemplo:
			- Em line 11 mostra como usar a função "copy()".

		- A capacidade da nova slice será menor do que a capacidade da slice original porque a nova matriz subjacente é menor.
*/