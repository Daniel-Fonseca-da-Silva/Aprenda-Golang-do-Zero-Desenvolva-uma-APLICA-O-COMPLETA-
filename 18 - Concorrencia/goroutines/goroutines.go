package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Ola mundo!")
	escrever("Programando no go :D")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
