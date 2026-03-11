# Base2 - Channel 
 
## Experiment
~~~go
package main

import "fmt"

func main() {
	done := make(chan bool) // channel

	go func() {
		fmt.Println("hello from goroutine")
		done <- true // channel sender
	}() 

	<- done // channel receiver
	fmt.Println("main finished")
}
~~~

## Observation
- "hello from goroutine" and "main finished" are printed

## Why?
- the channel "done" is unbuffered channel. So sender and receiver are both ready for send or receive.
- when the sender(goroutine) is done, the receiver(main goroutine) is start.

## Deeper
- if we change 
~~~go
make(chan bool) -> make(chan bool, 1)
~~~
This is buffered goroutine. The sender can send until the buffer number is filled.
## Key Insights
1. A goroutine does not automatically synchronize with main.
2. when main exits. the entire program terminates -- even if other goroutines are still running
3. An unbuffered channel createes a strong synchronization point between sender and receiver.
4. A buffered channel weakens synchronization because sends may proceed without an immediate receiver.
5. Sending to a channel may block if:
    - The channel is unbuffered and no reciver is ready.
    - The channel buffer is fulled.
6. Sending to a closed channel causes panic.
7. Using time.Sleep is not synchronization

## Mental Model
### Goroutine
A goroutine is a function executing concurrently, but it does not keep the program alive

main contols the lifetime of the process

If main returns -> everything stops.

### Unbuffered Channel
Think of it as a handshake.
 
Sender and receiver must meet at the same time.
send <-> receive
Both pause until the other side is ready.
Strong synchronization

### Buffered Channel
Think of it as a mailbox.

Sender can drop a message and leave (if mailbox not full).
Receiver can check it later.
Weaker synchronization

### Corret Waiting
Do NOT wait by time.

no: Sleep-based waiting
yes: Event-based waiting(channel / sync primitives)
