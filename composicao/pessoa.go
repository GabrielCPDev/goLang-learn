package composicao

import (
	"fmt"
	"goland/encapsulamento"
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         encapsulamento.Endereco
	DataDeNacsimento time.Time
	Idade            int
}

func (p *Pessoa) CalculaIdade()  {
	anoNascimento := p.DataDeNacsimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoNascimento
	fmt.Println("metodo calcula idade: ", p.Idade )
}