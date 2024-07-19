package main

import (
	"fmt"
	"goland/composicao"
	"goland/encapsulamento"
	"goland/heranca"
	"time"
)

func main() {
	fmt.Println("struct")

	endereco := encapsulamento.Endereco{
		Rua:    "Rua x",
		Numero: 20,
		Cidade: "Asgard",
	}

	pessoa := composicao.Pessoa{
		Nome:             "Oscar",
		Endereco:         endereco,
		DataDeNacsimento: time.Date(1996, 06, 14, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(endereco)
	fmt.Println(pessoa)
	pessoa.CalculaIdade()

	fmt.Println("Idade: ", pessoa.Idade)

	autoMoto := heranca.Automovel{
		Ano:    2022,
		Placa:  "XPTO",
		Modelo: "XPTO",
	}

	moto := heranca.Moto{
		Automovel:   autoMoto,
		Cilindradas: 250,
	}

	fmt.Println(moto)

	var itens []string

	itens = append(itens, "Arroz")
	itens = append(itens, "feijão")
	itens = append(itens, "Carne")

	compra, err := heranca.NewCompra("Mercado A", time.Now(), itens)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println(compra)
	}
}
