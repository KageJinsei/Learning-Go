package main
import ("fmt")

func main() {
	var x int = 500
	var y int = -4500

	fmt.Printf("Type: '%T', value: '%v'", x, x)
	fmt.Printf("\nType: '%T', value: '%v'", y, y)
}

/*
	Signed Integers (Inteiros Assinados) em Go:
		- Números inteiros assinados, declarados com uma das palavras chave "int", podem armazenar valores positivos e negativos (ex: line 5 and 6).
		- Go tem cinco palavras-chave de tipos de inteiros assinados:

			Type	Size		     Range
			int   | Depende da plataforma: 32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits (int32 e int64 mostram o range)
			int8  | 8 bits/1 byte  | -128 to 127
			int16 |	16 bits/2 byte | -32768 to 32767
			int32 |	32 bits/4 byte | -2147483648 to 2147483647
			int64 | 64 bits/8 byte | -9223372036854775808 to 9223372036854775807
*/