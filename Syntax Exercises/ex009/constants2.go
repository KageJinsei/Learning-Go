package main
import ("fmt")

const A int = 1
const B = 1

func main() {
	fmt.Println(A)
	fmt.Println(B)
}

/*
	Tipos de constantes em GO:
		- Existem dois tipos de constantes:
			1 - Constantes digitadas
			2 - Constantes não digitadas
	
	Constantes digitadas:
		- Constantes digitadas são declaradas com um tipo definido (ex: line 4).

	Constantes não digitadas:
		- Constantes não digitadas são declaradas sem um tipo definido (ex: line 5).
		Note: Nesse caso, o tipo de constante é inferido a partir do valor recebido (significa que o compilador decido o tipo da constante com bae no valor).
*/