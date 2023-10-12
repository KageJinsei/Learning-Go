package main
import ("fmt")

func main() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e, f = 6, "Hello"
	g, h := 7, "World"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

/*
	Declaração de múltiplas variáveis em Go:
		- Em Go, é possível declarar múltiplas variáveis na mesma linha (ex: line 5).
	
	Note: Se você usar a palavra-chave "type" (ex: int), só será possível declarar um tipo de variável por linha.

	Obs: Se a palavra-chave "type" não for especificada, você pode declarar diferentes tipos de variáveis na mesma linha (ex: line 12 and 13).
*/