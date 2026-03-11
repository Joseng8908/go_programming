package main

import "fmt"

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("hello from goroutine")
		done <- true
	}()

	<-done
	fmt.Println("main finished")
}
