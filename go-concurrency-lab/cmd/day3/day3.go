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
