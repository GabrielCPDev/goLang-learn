package main

import "fmt"

type endereco struct {
	rua string
	numero int
	cidade string
}

func main() {
	fmt.Println("struct")

	endereco := endereco {
		rua: "Rua x",
		numero: 20,
		cidade: "Asgard",
	}

	fmt.Println(endereco)
}