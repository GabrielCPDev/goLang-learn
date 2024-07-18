package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 20
	fmt.Println("funcao main")
	fmt.Println(x, *y)
	fmt.Println(&x, y)


	ImprimirValores(&x, y)
}

func ImprimirValores(x *int, y *int) {
	fmt.Println("funcao imprime")

	fmt.Println(x, y)
}
