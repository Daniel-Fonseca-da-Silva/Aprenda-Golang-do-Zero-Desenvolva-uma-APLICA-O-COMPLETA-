package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Mario", "Souza", 30, 175}
	fmt.Println(p1)

	e1 := estudante{p1, "medicina", "UBA"}
	fmt.Println(e1.altura)
}
