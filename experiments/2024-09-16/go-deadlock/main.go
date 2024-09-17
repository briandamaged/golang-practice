package main

import "fmt"

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)

	go func() {
		ch1 <- "Hello from GoRoutine!" // Blocks until ch1 is read
		fmt.Println(<-ch2)
	}()

	ch2 <- "Hello from main!" // Blocks until ch2 is read
	fmt.Println(<-ch1)
}
