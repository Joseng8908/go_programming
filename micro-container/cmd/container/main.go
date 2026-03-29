package main

import (
	"fmt"
	"log"
	"micro-container/internal/container"
	"micro-container/internal/storage"
	"os"
	"os/exec"
	"syscall"
)


func main() {
	// 사용법
	if len(os.Args) < 2 {
		log.Fatal("사용법: ./micro-container <commands>, [run | child] [args...]")
		return
	}

	switch os.Args[1] {
	case "run": run()
	
	case "child": child()

	default: log.Fatal("알 수 없는 명령어")
	
	}
}

func run() {
	fmt.Printf("부모 프로세스 시작(PID: %d)\n", os.Getgid())

	// 자기 자신을 child 인자에 붙이고 실행할 준비
	// os.Args[2:]는 사용자가 입력한 명령어 (ex) /bin/sh)ㅏ
	args := append([]string{"child"}, os.Args[2:]...)
	cmd := exec.Command("/proc/self/exe", args...)

	// 격리된 네임스페이스 설정하기
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	c := &container.Container{
		ID: "test-container",
		Cmd: cmd,
	}

	// 마운트 하기
	_, err := storage.MountOverlay(c.ID)
	if err != nil {
		fmt.Printf("마운트 에러: %v\n", err)
	}
	fmt.Println("마운트 완료")
	defer storage.UnmountOverlay(c.ID)

	if err := c.Run(); err != nil {
		fmt.Printf("부모 에러: %v\n", err)
		os.Exit(1)
	}
}

func child() {
	if err := container.Child(os.Args[2:]); err != nil {
		fmt.Printf("자식 에러: %v\n", err)
		os.Exit(1)
	}
}
