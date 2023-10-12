package main
import ("fmt")

const PI = 3.14

func main() {
	fmt.Println(PI)
}

/*
	Constantes em Go:
		- Se uma variável deve ter um valor fixo e não será alterado, você pode utilizar a palavra-chave "const".
		- A palavra-chave "const" declara uma variável como "constante", o que significa que é imutável.

		Syntax de uma constante: const CONSTNAME type = value

		Note: O valor de uma constante deve ser atribuído quando você a declara.
	
	Declarando uma constante:
		- exemplo em: line 4

	Regras de uma constante:
		- Mesmas regras de nomeação que as variáveis;
		- Nomes das constantes são geralmente escritas em letras maiúsculas (para facilitar a identificação e diferenciação das variáveis);
		- As constantes podem ser declaradas dentro e fora de uma função.
*/