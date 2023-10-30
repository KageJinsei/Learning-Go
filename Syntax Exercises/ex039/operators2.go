package main
import ("fmt")

func main() {
  var x = 10
  x += 5
  fmt.Println(x)
}

/*
	Operadores de atribuição em Go:
		- Operadores de atribuição são usados para atribuir valores a variáveis. No exemplo em line 5, usamos o operador de atribuição "=" para atribuir o valor 10 na variável "x".

		- O operador de atribuição de adição (+=) adiciona um valor a uma variável (ex: line 6).

		- Uma lista com todos os operadores de atribuição:

			Operator		Example			Same As	
			=				x = 5			x = 5	
			+=				x += 3			x = x + 3	
			-=				x -= 3			x = x - 3	
			*=				x *= 3			x = x * 3	
			/=				x /= 3			x = x / 3	
			%=				x %= 3			x = x % 3	
			&=				x &= 3			x = x & 3	
			|=				x |= 3			x = x | 3	
			^=				x ^= 3			x = x ^ 3	
			>>=				x >>= 3			x = x >> 3	
			<<=				x <<= 3			x = x << 3
*/