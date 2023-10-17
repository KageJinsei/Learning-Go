package main
import ("fmt")

func main() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: '%T', value: '%v'\n", txt1, txt1)
	fmt.Printf("Type: '%T', value: '%v'\n", txt2, txt2)
	fmt.Printf("Type: '%T', value: '%v'\n", txt3, txt3)
}

/*
	Tipo de dados String em Go:
		- O tipo de dados "string" é usado para armazenar uma sequência de caracteres (texto). Os valores das strings devem ser cercados por aspas duplas (ex: line 5, 7, 9, 10 and 11).
*/