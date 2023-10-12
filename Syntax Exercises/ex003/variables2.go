package main
import ("fmt")

func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
	Declarando variáveis em Go sem valor inicial:
		- Em Go, todas a variáveis são inicializadas. Portanto, se você declarar uma variável sem um valor inicial, o seu valor será definido como o valor padrão do seu tipo.
	
	No exemplo acima, existem 3 variáveis:
		- a (line 5)
		- b (line 6)
		- c (line 7)

	Ao executar o código, podemos ver que as variáveis já possuem um valor padrão:
		- a = ""
		- b = 0
		- c = false
*/