package main
import ("fmt")

func main() {
	var student1 string = "Kage" // tipo definido (string)
	var student2 = "Gwenhwyfar"  // tipo inferido (o compilador deduz)
	x := 2                       // tipo inferido (o compilador deduz)

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

/*
	Tipos de variáveis em Go:
		int: Armazena números inteiros (ex: 123 ou -123);
		float32: Armazena números com ponto flutuante (ex: decimais, 50,00 ou -50,00);
		string: Armazena texto (ex: "Hello, World!"). As strings são definidas com aspas duplas;
		bool: Armazena valores com dois estados: verdadeiro ou falso.
	
	Declarando variáveis em Go:
		- Para declarar uma variável em Go, é necessário utilizar o comando "var", inserir o nome e o tipo da variável.
		Syntax de uma variável: var nomeVariável tipo = valor
		
		Note: Você sempre precisa especificar o tipo ou o valor (ou ambos) ao criar uma variável.

		- Outra forma de declarar uma variável é utilizando o sinal ":=", seguido do valor da variável.
		Syntax de uma variável com esta forma: nomeVariável := valor

		Note 1: Neste caso, o tipo da variável é deduzido com base no valor (o compilador decide o tipo da variável com base no valor recebido).
		Note 2: Não é possível declarar uma variável utilizando o ":=" sem atribuir um valor a ela.

	Declaração de Variável com valor inicial:
		- Se você sabe qual é o valor a ser definido, você pode declarar uma variável e atribuir o valor a ela na mesma linha (exemplo em line 5, 6 and 7).
		Note: Os tipos das variáveis "student2" e "x" é inferido de seus valores.	
*/