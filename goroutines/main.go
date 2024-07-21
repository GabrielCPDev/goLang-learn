package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 100)
	go setList(channel)

	for valor := range channel {
		fmt.Println("recebendo: ", valor)
		time.Sleep(time.Second)
	}
}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)

	}
	close(channel)
}
