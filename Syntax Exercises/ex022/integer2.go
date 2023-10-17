package main
import ("fmt")

func main() {
	var x uint = 500
	var y uint = 4500

	fmt.Printf("Type: '%T', value: '%v'", x, x)
	fmt.Printf("\nType: '%T', value: '%v'", y, y)
}

/*
	Unsigned Integers (Inteiros Não Assinados) em Go:
		- Números inteiros não assinados, declarados com uma das palavras-chave "uint", só podem armazenar valores não negativos (ex: line 5 and 6).
		- Go tem cinco palavras-chaves de tipos de inteiros não assinados:
			Type	 Size		      Range
			uint   | Depende da plataforma: 32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits (uint32 e uint64 mostram o range)
			uint8  | 8 bits/1 byte  | 0 to 255
			uint16 | 16 bits/2 byte | 0 to 65535
			uint32 | 32 bits/4 byte | 0 to 4294967295
			uint64 | 64 bits/8 byte | 0 to 18446744073709551615

		Obs: O tipo de número inteiro que você irá escolher, depende do valor aque a variável tem para armazenar (ex: "uint8" não consegue armazenar o valor "1000").
*/