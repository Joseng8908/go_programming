package main 

import(
	"fmt"
	"context"
	"sync"
) 

func main() {

}

func WorkerPool(ctx context.Context, jobs <- chan int) {
	var wg sync.WaitGroup

	// 1. 워커 3개 가동 (Fan - out: 한 단계에서 여러 워커를 띄워 병렬로 처리하는 것)
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <- ctx.Done():
					fmt.Printf("워커 %d 퇴근 \n", id)
					return
				case job, ok := <- jobs:
					if !ok {return}
					fmt.Printf("워커 %d 가 일하는 중: %d\n", id, job)
				}
			}
		}(i)
	}
	wg.Wait()
}
