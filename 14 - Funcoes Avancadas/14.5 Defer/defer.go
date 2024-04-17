package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entra na funcao e verifica aprovacao")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}
	
	return false
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
}
