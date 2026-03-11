# Base1: what is goroutine?

Before, out program just executed by one line
~~~go
func main() {
    fmt.Println("hello")
}
~~~
This is synchronous execution.

But, goroutine execute the function simultaneously.
~~~go
go someFunction()
~~~
This is goroutine context.

## Experiment
~~~go
package main

import (
	"fmt"
)

func main() {
	go sayHello() // goroutine
    
    // time.Sleep(time.Second) <- can ensure print sayHello
	fmt.Println("main finished")
}

func sayHello() {
	fmt.Println("hello from goroutine")
}
~~~

## Observation
- "main finished" is printed
- "hello from goroutine" is not printed.
- But if we wait seconds, it can be printed -> main func doesn't wait for goroutine 

## Key Insight
1. goroutine is not OS Thread
2. go runtime manage goroutine
3. we don't know scheduling method

But use time.Sleep is not good choice
