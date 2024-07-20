package main

import (
	"fmt"
	"os"
)

func main() {
	ReadyFile()
	fmt.Println("fim")
}

func ReadyFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")

		}
	}()

	_, err := os.Open("./arquivoTest")
	if err != nil {
		panic("FileNotExists")
	}
}
