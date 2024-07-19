package main

import (
	"fmt"
	"goland/composicao"
	"goland/encapsulamento"
)

func main() {
	fmt.Println("struct")

	endereco := encapsulamento.Endereco {
		Rua: "Rua x",
		Numero: 20,
		Cidade: "Asgard",
	}
	
	pessoa := composicao.Pessoa {
		Nome: "Oscar",
		Endereco: endereco,
	}

	fmt.Println(pessoa)
}