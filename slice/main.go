package main

import "fmt"

// interesting behaviour of nil as a slice of length 0
func main() {
	slice := []int{}
	var nilSlice []int
	fmt.Println(len(slice))
	fmt.Println(slice)
	fmt.Println(nilSlice == nil)
}
