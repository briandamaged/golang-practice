package main

import "fmt"

type Oops any

func main() {
	var x Oops = []int{1, 2, 3}
	var y Oops = []int{1, 2, 3}

	// PANIC!!!
	fmt.Println(x == y)
}
