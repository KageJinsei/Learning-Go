package main
import ("fmt")

func main() {
	var a = 15 + 25
	fmt.Println(a)

	var (
		sum1 = 100 + 50 // 150 (100 + 50)
		sum2 = sum1 + 250 // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)

	fmt.Println(sum3)
}

/*
	Operadores em Go:
		- Operadores são usados para executar operações em variáveis e valores. O operador "+" soma dois valores, como no exemplo em line 5.

		- Embora o operador "+" seja frequentemente utilizado para somar dois valores, ele também pode ser usado para somar uma variável e um valor, ou uma variável e outra variável (ex: line 9, 10 and 11).

		Go Divide os operadores nos seguintes grupos:
			- Operadores aritméticos
			- Operadores de atribuição
			- Operadores de comparação
			- Operadores lógicos
			- Operadores bitwise

	Operadores Aritméticos em Go:
		- Operadores aritméticos são usados para executar operações matemáticas comuns.
		- Operadores aritiméticos utilizados em Go:

			Operator		Name					Description										Example

			+				Addition				Adds together two values						x + y	
			-				Subtraction				Subtracts one value from another				x - y	
			*				Multiplication			Multiplies two values							x * y	
			/				Division				Divides one value by another					x / y	
			%				Modulus					Returns the division remainder					x % y	
			++				Increment				Increases the value of a variable by 1			x++	
			--				Decrement				Decreases the value of a variable by 1			x--
*/