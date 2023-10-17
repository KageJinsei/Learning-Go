package main
import ("fmt")

func main() {
	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
}

/*
	Verbos de formatação de inteiros em Go:
		- Os seguintes verbos podem ser usados com o tipo de dados inteiro (int):
			Verb		Description
			%b 	   -	Base 2 (converte em binário)
			%d 	   -	Base 10 (base decimal, sistema numérico padrão)
			%+d    -	Base 10 e sempre mostra o sinal (base decimal mostrando o sinal a frente)
			%o	   - 	Base 8 (sistema octal, representa números usando os dígitos de 0 a 7)
			%O	   - 	Base 8, com "Oo" inicial (sistema octal, ela representa números usando os dígitos de 0 a 7 com "Oo" no início)
			%x	   - 	Base 16, lowercase (sistema hexadecimal, é usada para representar números usando os dígitos de 0 a 9 e as letras de A a F (ou a-f) e em letra minúscula)
			%X	   - 	Base 16, uppercase (sistema hexadecimal, é usada para representar números usando os dígitos de 0 a 9 e as letras de A a F (ou a-f) e em letra maiúscula)
			%#x	   - 	Base 16, com "ox" inicial (sistema hexadecimal, é usada para representar números usando os dígitos de 0 a 9 e as letras de A a F (ou a-f) com "0x" no início)
			%4d	   - 	Pad com espaços (4 de largura e ajustando o valor na direita)
			%-4d   - 	Pad com espaços (4 de largura e ajustando o valor na esquerda)
			%04d   - 	Pad com zeros (ocupando 4 de largura)
*/