package main

import (
	"fmt"
)

func sendDataToChannel(ch chan string, s string) {
	ch <- s
	close(ch)
}

func main() {
	ch := make(chan string)

	go sendDataToChannel(ch, "hello")

	value, ok := <-ch

	if !ok {
		  fmt.Println("channel closed")
	}

	fmt.Println(value)
}