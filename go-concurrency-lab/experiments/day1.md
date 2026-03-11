# Day 1 - Channel Close and Deadlock
 
## Experiment

Created an unbuffered channel and ranged over it without closing the channel.

~~~go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i ++ {
        ch <- i
    }
}()

for v := range ch{
    fmt.Println(v)
}
~~~

## Observation
- Values 0~4 are printed.
- Program does not exit.
- Runtime reports.
~~~bash
fatal error: all goroutines are asleep - deadlock!
~~~

## Why?
- range over a channel only exits when the channel is closed.
- The sender finished sending but did not close the channel.
- The receiver is blocked waiting for more values.
- Since no goroutine can proceed, the runtime detects a deadlock.

This is not a race condition

There is no shared memory access.
This is a synchronization / flow control issue.

## Key Insights
1. The sender is reponsible for closing the channel.
2. range depends on close.
3. Unbuffered channels act as synchronization points.
4. Deadlock != race condition.

## Mental Model
Send -> Receive -> Repeat
When sending ends -> Close channel -> Receiver exits loop
if not closed -> Receiver waits forever
