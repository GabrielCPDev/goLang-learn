package main

import (
	"fmt"
	"goland/encapsulamento"
)

func main() {
	fmt.Println("struct")

	endereco := encapsulamento.Endereco {
		Rua: "Rua x",
		Numero: 20,
		Cidade: "Asgard",
	}

	fmt.Println(endereco)
}