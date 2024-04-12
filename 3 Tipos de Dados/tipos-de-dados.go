package main

import "fmt"

func main() {
	var numero int64 = 100000000
	fmt.Println(numero)

	var numero2 uint8 = 100
	fmt.Println(numero2)

	// Alias
	// int32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// Byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.5
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 15200.45
	fmt.Println(numeroReal2)

	var str string = "texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)
}
