package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario:\n%s \ncom idade %d", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 18}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Fernando", 30}
	fmt.Println(usuario2)
	usuario2.salvar()

	maiorDeIDade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIDade)

	usuario2.fazAniversario()
	fmt.Println(usuario2.idade)

}
