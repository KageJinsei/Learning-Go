package main
import ("fmt")

func main() {
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
}

/*
	Verbos de formatação de String em Go:
		- Os verbos a seguir podem ser usados com o tipo de dados string:
			Verb		Description
			%s 	   -	Imprime o valor com string simples
			%q 	   -	Imprime o valor com uma string entre aspas duplas
			%8s    -	Imprime o valor com string simples (com uma largura de 8 e ajustando à direita)
			%-8s   - 	Imprime o valor com string simples (com uma largura de 8 e ajustando à esquerda)
			%x	   - 	Imprime o valor como despejo hexadecimal de valores de bytes
			% x	   - 	Imprime o valor como dump hexadecimal com espaços
*/