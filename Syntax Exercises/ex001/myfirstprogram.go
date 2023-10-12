package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}

/*
	Um script em Go consiste nas seguintes partes:
		- Declaração do pacote (line 1);
		- Importação de pacotes (line 2);
		- Funções (line 4);
		- Declaração de expressões (line 5).

	Observações importantes:
		- Line 1: Na linguagem Go, cada programa faz parte de uma pacote (Package). Definimos isso utilizando a palavra-chave "package". Neste caso, o programa pertence ao pacote "main".
		- Line 2: O comando "import ("fmt")" nos permite importar os arquivos que estão incluídos no pacote "fmt".
		- Line 4: O comando "func main() {}" é uma função. Qualquer código que está dentro das suas chaves ({}) será executado.
		- Line 5: O comando "fmt.Println()" é uma função disponibilizada a partir do pacote "fmt". É utilizada para imprimir texto na tela.

		Note: Em Go, qualquer código executável pertence ao pacote "Main".

	Declarações em Go:
		- O comando "fmt.Println("Hello World!")" é uma declaração.
		- Em Go, as declarações são separadas após pular uma linha ou por um ponto e vírgula ";" no final.
*/ 