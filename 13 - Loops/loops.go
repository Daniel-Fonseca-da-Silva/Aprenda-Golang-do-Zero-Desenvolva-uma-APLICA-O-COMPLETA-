package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Printf("Incrementando o valor de i: %d\n", i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Printf("Incrementando o valor de j: %d\n", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"Marcos", "Alexandre", "Annibal"}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Fulano",
		"sobrenome": "De Tal",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Ciclano", "De Tal"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }

	for {
		fmt.Println("Loop infinito!")
		time.Sleep(time.Second)
	}

}
