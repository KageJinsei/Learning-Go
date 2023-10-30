package main
import ("fmt")

func main() {
	var x = 5
	var y = 3
	fmt.Println(x > y) // Retorna 1 (verdadeiro) porque 5 é maior que 3
}

/*
	Operadores de comparação em Go:
		- Operadores de comparação são usados para comparar dois valores.

		Note: O valor de retorno de uma comparação é verdadeiro (1) ou falso (0).

		- No exemplo (line 7), é utilizado o operador de "maior que" (>) para descobrir se 5 é maior que 3.

		- Uma lista de todos os operadores de comparação:

			Operator			Name								Example
			==					Equal to							x == y	
			!=					Not equal							x != y	
			>					Greater than						x > y	
			<					Less than							x < y	
			>=					Greater than or equal to			x >= y	
			<=					Less than or equal to				x <= y
*/