package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go ShowMessage(&wg)
	go CallApi(&wg)
	go CallDatabase(&wg)

	wg.Wait()
}

func ShowMessage(wg *sync.WaitGroup) {
	fmt.Println("ShowMessage")
	wg.Done()
}
func CallApi( wg *sync.WaitGroup) {
	fmt.Println("CallApi")
	wg.Done()
}
func CallDatabase( wg *sync.WaitGroup) {
	fmt.Println("CallDatabase")
	wg.Done()
}