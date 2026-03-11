# Day 3 - Race Condition & Mutex
 
## Experiment

~~~go
package main
import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	mu sync.Mutex
	count int
}

// 1. Race condition 실험 (no mutex)
func raceCondition() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// 원자적이지 않은 연산
				// 즉 하나가 깨져도 그냥 돌아감
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("[Race Condition] Expected: 10000, Actual: %d\n", count)
}

// 2. Mutex 사용(정상 동작)
func safeCounter() {
	var mu sync.Mutex
	var count int
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() 
			for j := 0; j < 1000; j++ {
				mu.Lock() 
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait() 
	fmt.Printf("[Safe Counter] Expected: 10000, Actual: %d\n", count)
}

// 3. Mutex 복사 실험(Deadlock 유발)
func passByValue(c Container) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("In function: Locked copy sucessfully")
}

func main() {
	raceCondition()
	safeCounter()

	fmt.Println("\n--- Mutex Copy Experiment ---")
	c := Container{}
	c.mu.Lock()

	fmt.Println("Main: Locked original")
	go passByValue(c)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: Finished (If 'In function' didn't print, it's a deadlocok)")
}
~~~

## Observation
- raceCondition() 에서는 Mutex가 없기 때문에 10,000이 나오지 않음
- Mutex를 복사해서 넘기면 특정 상황에서 고루틴이 멈춰버림

## Why?
- count++는 Read - Add - Write의 3단계라 원자적이지 않음 -> 즉 Read를 한 상태인데 다른 고루틴에서
또 Read를 하려고 시도하면 이 작업이 1개가 되어서 2개가 올라야 하는데, 1개만 오르는 오류가 생김
- Mutex는 복사될 때 '잠금 상태', 비트까지 복사되므로, 복사본은 주인을 잃은채 평~생 잠겨버림.

## Key Insights
- RWMutex: 읽기 작업이 많을 때 성능을 비약적으로 높임
- Lock Ordering: 자물쇠를 잠그는 순서(A -> B)를 통일하는 것이 데드락 예방의 기본.
- Pass by Pointer: Mutex를 포함한 모든 것은 포인터로 전달

## Mental Model
"자물쇠는 물건이 아니라 권한이다!"
자물쇠를 복사하는 건 권한을 복사하는게 아님 -> 열쇠 구멍이 없는 자물쇠를 하나 더 만드는 것임
권한은 오직 한개!, 순서는 엄격하게

## Supplement 
go run -race ~.go 를 이용해서 raceCondition 발생 여부를 확인할 있다!
