package main

import (
	"fmt"
	"os"
	"micro_container/container"
)

func main() {
	// 사용법 설명: go run main.go <command> <args...>
	// ex) go run main.go /bin/sh
	if len(os.Args) < 2 {
		fmt.Println("사용법: go run main.go <command>")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	fmt.Printf("새로운 컨테이너 생성중,,,%s\n", cmd)

	// 컨테이너 생성
	c := container.NewContainer("1", cmd, args...)

	// 컨테이너 실행
	if err := c.Start(); err != nil {
		fmt.Printf("컨테이너 실행 실패: %v\n", err)
		return
	}

	// 프로세스 대기
	// 커테이너 프로세스각 끝나야 끝남
	err := c.Cmd.Wait()
	if err != nil {
		fmt.Printf("컨테이너가 에러남: %v\n", err)
	} else {
		fmt.Println("컨테이너 종료")
	}
}
