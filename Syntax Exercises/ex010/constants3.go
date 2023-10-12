package main
import ("fmt")

const(
	A int = 1
	B = 3.14
	C = "Hello!"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

/*
	Declaração de múltiplas constantes em Go:
		- Várias constantes podem ser agrupadas em um bloco para melhorar a legibilidade (assim como nas variáveis).
*/