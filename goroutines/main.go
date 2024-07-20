package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i<10 ; i++ {
		go ShowMessage(strconv.Itoa(i))
	}

}

func ShowMessage(message string) {
	fmt.Println(message)
}