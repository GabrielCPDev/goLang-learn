package main

import (
	"fmt"
	"goland/composicao"
	"goland/encapsulamento"
	"time"
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
		DataDeNacsimento: time.Date(1996, 06, 14, 0, 0, 0, 0, time.Local),
	}


	fmt.Println(endereco)
	fmt.Println(pessoa)

	idade := composicao.CalculaIdade(pessoa)
	
	idade = pessoa.IdadeAtual()
	fmt.Println(idade)
}