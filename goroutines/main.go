package main

import "fmt"

func main() {

	channel := make(chan int)
	go setList(channel)

	for valor := range channel {
		fmt.Println(valor)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
