package main

import "os"

func main() {
	_, err := os.Open("nomeArquivoTeste") 

	if err != nil {
		panic(err)
	}
}