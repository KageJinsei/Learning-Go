package main
import ("fmt")

func main() {
	var b1 bool = true // tipo declarado com valor inicial
	var b2 = true // tipo não declarado com valor inicial
	var b3 bool // tipo declarado sem valor inicial
	b4 := true // tipo não declarado com valor inicial

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}

/*
	Tipo de dados booleanos em Go:
		- Um tipo de dados booleano é declarado com a palavra chave "bool" e só pode tomar os valores "true" ou "false".
		- O valor padrão de um tipo de dados booleano é "false".
*/