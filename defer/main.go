package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./arquivoTest.txt")
	defer file.Close()
	defer finalizandoArquivo()

	if err != nil {
		panic(err)
	}

	file.Write([]byte("teste"))
}

func finalizandoArquivo(){
	fmt.Println("Finalizando arquivo")
}
