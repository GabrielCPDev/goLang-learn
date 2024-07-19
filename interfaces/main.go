package main

import (
	"errors"
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	radius float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {

	fmt.Println("===== Interface=====")

	circulo := circulo{radius: 2.5}
	retangulo := retangulo{largura: 14, altura: 4}

	exibirArea(circulo)
	exibirArea(retangulo)

	fmt.Println("===== Interface - mais exemplos =====")

	ExibeErro(errors.New("teste"))

	p := ProblemaDeNetwork{ rede: true, hardware: false}

	ExibeErro(p)

}

func exibirArea(g geometria) {
	fmt.Println("Area: ", g.area())
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	} else {
		return "Outro problema"
	}
}

func ExibeErro(err error) {
	fmt.Println(err.Error())
}
