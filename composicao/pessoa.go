package composicao

import (
	"goland/encapsulamento"
	"time"
)

type Pessoa struct {
	Nome     string
	Endereco encapsulamento.Endereco
	DataDeNacsimento time.Time
}

func (p Pessoa) IdadeAtual() int {
	anoNascimento := p.DataDeNacsimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}

func CalculaIdade(p Pessoa) int {
	anoNascimento := p.DataDeNacsimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoNascimento
}