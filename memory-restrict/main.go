package main

import (
	"fmt"
	"time"
)
//
//func main() {
//	// 100ㅡ만큼 할당하고, 죽지 않게 대기시키기
//	b := make([]byte, 100*1024*1024)
//	for i := range b {b[i] = 1}
//	time.Sleep(10*time.Minute)
//}

func main() {

	// 10Mb씩 계속 할당하면서 무한 루프 돌리기
	var memoryLeak [][]byte

	for i := 0; i < 20; i++ {
		b := make([]byte, 10*1024*1024) 
		for j := range b {b[j] = 1}
		memoryLeak = append(memoryLeak, b)
		fmt.Printf("현재 할당량: %d Mb\n", (i+1)*10)
		time.Sleep(1*time.Second)
	}
	select {}
}
