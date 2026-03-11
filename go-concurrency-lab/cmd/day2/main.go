package main

import (
	"context"
	"fmt"
)

func main() {
	// 1. Unbuffered 채널 생성
	dataChan := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	// 2. 워커 고루틴 실행
	go func() {
		// 이 defer는 과연 실행될 수 있을까?
		defer fmt.Println("Worker finished")
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-dataChan:
				fmt.Println("Received:", d)
			}
		}
	}()

	dataChan <- 1
	// cancel()을 호출하지 않고, dataChan도 닫지 않은 채 메인이 끝난다면?
	// 혹은 dataChan을 닫았는데 워커가 모르고 보내려고 한다면?
}
