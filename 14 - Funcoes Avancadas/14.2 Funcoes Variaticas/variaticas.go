package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 5, 6, 100, 200, 100, 50)
	fmt.Println(totalDaSoma)
	escrever("Ola GO!", 20, 25, 35, 25, 40)
}
