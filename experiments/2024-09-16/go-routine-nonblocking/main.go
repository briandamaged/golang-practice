package main

import (
	"fmt"
	"time"
)

func main() {
	var xChannel = MakeIntCounter()
	var yChannel = MakeIntCounter()
	for {
		select {
		case x := <-xChannel: // Executes when xChannel is ready
			fmt.Printf("x: %d\n", x)
		case y := <-yChannel: // Executes when yChannel is ready
			fmt.Printf("y: %d\n", y)
		default: // Executes when all channels are blocked
			fmt.Println("No channels are ready. Waiting 1 second...")
			time.Sleep(time.Second)
		}
	}
}

func MakeIntCounter() <-chan int {
	var ch = make(chan int)
	go func() {
		for x := 0; ; x++ {
			ch <- x
			time.Sleep(3 * time.Second)
		}
	}()

	return ch
}
