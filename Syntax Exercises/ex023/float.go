package main
import ("fmt")

func main() {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	var z float64 = 1.7e+308

	fmt.Printf("Type: '%T', value: '%v'", x, x)
	fmt.Printf("\nType: '%T', value: '%v'", y, y)
	fmt.Printf("\nType: '%T', value: '%v'", z, z)
}

/*
	Tipos de dados Float em Go:
		- Os tipos de dados "float" são usados para armazenar números positivos e negativos com um ponto decimal, como 35.3, -2.34 ou 3597.34987.
		- O tipo de dados "float" tem duas palavras-chave:
			
			Type	  Size		Range
			float32 | 32 bits | -3.4e+38 to 3.4e+38
			float64 | 64 bits | -1.7e+308 to +1.7e+308

		Obs: o tipo padrão para "float" é "float64". Se você não especificar o tamanho (32 ou 64), será "float64".

	Utilizando float32:
		- O exemplo (line 5 and 6) mostra como declarar algumas variáveis do tipo "float32".

	Utilizando float64:
		- O tipo de dados "float64" pode armazenar um conjunto maior de números do que "float32". O exemplo (line 7) mostra como declarar uma variável do tipo "float64".

	Qual tipo de float usar?
		- O tipo de float a escolher depende do valor que a variável terá para armazenar. Se float32 recebe um número maior do que o limite que ele suporta, causará um erro e, nesse caso, é correto usar float64.
*/