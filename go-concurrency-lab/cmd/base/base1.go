package main

import (
	"fmt"
)

func main() {
	go sayHello()

	fmt.Println("main finished")
}

func sayHello() {
	fmt.Println("hello from goroutine")
}
