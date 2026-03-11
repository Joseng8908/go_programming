package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending:", i)
			ch <- i
		}
		//close (ch) // <- Day 1: intentionally not closing
	}()

	for v := range ch {
		fmt.Println("received:", v)
	}

	fmt.Println("main exited")
}
