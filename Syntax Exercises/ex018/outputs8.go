package main
import ("fmt")

func main() {
	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)
}

/*
	Verbos de formatação de float em Go:
		- Os seguintes verbos podem ser usados com o tipo de dados float:
		Verb		Description
			%e 	   -	Notação científica com 'e' como expoente
			%f 	   -	Ponto decimal, sem expoente
			%.2f   -	Largura padrão, duas casas decimais
			%6.2f  -	6 de largura, duas casas decimais
			%g 	   -	Expoente conforme necessário, apenas dígitos necessários

*/