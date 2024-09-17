package main

import "fmt"

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)

	go func() {
		select {
		case ch1 <- "Hello from GoRoutine!":
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}()

	select {
	case ch2 <- "Hello from main!":
	case msg := <-ch1:
		fmt.Println(msg)
	}
}
