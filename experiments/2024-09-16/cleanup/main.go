package main

import (
	"context"
	"fmt"
	"time"
)

func MakeIntCounter(ctx context.Context) <-chan int {
	var ch = make(chan int)
	go func() {
		for x := 0; ; x++ {
			// Emit the current value of `x` if `ch` is ready
			select {
			case ch <- x:
			case <-ctx.Done():
				fmt.Println("Terminating the GoRoutine!")
				close(ch)
				return
			default:
			}

			fmt.Println("tick...")
			time.Sleep(time.Second)
		}
	}()

	return ch
}

func doStuff() {
	var ctx, cancel = context.WithCancel(context.Background())

	// Ensure `ctx` will get cancelled when `doStuff()` returns
	defer cancel()

	var ch = MakeIntCounter(ctx)
	for x := range ch {
		fmt.Println(x)
		if x >= 10 {
			fmt.Println("That's enough! I'm done!")
			return
		}
	}
}

func main() {
	doStuff()
	fmt.Println("Sleeping 5 seconds before terminating...")
	time.Sleep(5 * time.Second)
	fmt.Println("Terminating!")
}
