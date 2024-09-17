package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 5)

	go func() {
		for x := 1; x < 10; x++ {
			ch <- x
		}
		close(ch)
		fmt.Println("Channel has been closed")
	}()

	for {
		// The book I'm reading states that `ok` indicates
		// that the channel is still opened, which I suspect
		// is slightly incorrect.
		// My guess is that `ok` actually indicates that a
		// value was read from the channel. This is important
		// because a buffered channel might be closed before
		// all of its values have been read.
		x, ok := <-ch
		fmt.Printf("%d %t\n", x, ok)
		time.Sleep(time.Second / 2)
		if !ok {
			break
		}
	}
}
