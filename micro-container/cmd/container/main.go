package main

import (
	"fmt"
	"log"
	"micro-container/internal/container"
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

	case "ps": ListContainers()

	case "exec": myExec(os.Args)

	case "nsenter":	container.NsEnter()
	
	case "stop" : stop()

	default: log.Fatal("알 수 없는 명령어")
	
	}
}

func run() {
	detach := false
	commandIdx := 2

	if len(os.Args) > 2 && os.Args[2] == "-d" {
		detach = true
		commandIdx = 3
	}
	fmt.Printf("로그: 부모 프로세스 시작(PID: %d), (Detach: %v)\n", os.Getgid(), detach)

	// 자기 자신을 child 인자에 붙이고 실행할 준비
	// os.Args[2:]는 사용자가 입력한 명령어 (ex) /bin/sh)ㅏ
	args := append([]string{"child"}, os.Args[commandIdx:]...)
	cmd := exec.Command("/proc/self/exe", args...)

	fmt.Println("로그: 파이프라인 구축...")
	r, w, _ := os.Pipe()
	cmd.ExtraFiles = []*os.File{r}
	
	// 격리된 네임스페이스 설정하기
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS,
	}
	if !detach {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	c := &container.Container{
		ID: "test-container",
		Cmd: cmd,
		SyncPipe: w,
	}

	if err := c.Run(detach); err != nil {
		fmt.Printf("로그: 부모 에러: %v\n", err)
		os.Exit(1)
	}
}

func child() {
	pipe := os.NewFile(3, "pipe")
	if pipe == nil {
		os.Exit(1)
	}

	msg := make([]byte, 2)
	pipe.Read(msg)
	pipe.Close()

	if err := container.Child(os.Args[2:]); err != nil {
		fmt.Printf("로그: 자식 에러: %v\n", err)
		os.Exit(1)
	}
}

func ListContainers() {
	// 컨테이너 정보 가져오기
	containers, err := container.GetInfoList()
	if err != nil {
		fmt.Printf("로그: 목록을 가져오는데 실패했습니다: %v\n", err)
		return
	}

	// 헤더 추력
	fmt.Printf("ID\t\tPID\t\tSTATUS\t\tCOMMAND\t\tCREATED\n")

	for _, item := range containers {
		fmt.Printf("%s\t%d\t\t%s\t\t%s\t\t%s\n",
		item.ID,
		item.Pid,
		item.Status,
		item.Command,
		item.CreateTime.Format("2026-03-31 16:47:05"),
		)
	}
}

func myExec(args []string) {
	if len(args) < 4 {
			fmt.Println("사용법: exec [컨테이너ID] [명령어]")
			return
		}
	containerID := args[2]
	command := args[3]

	// 저장된 정보 가져오기
	info, err := container.ReadContainerInfo(containerID)
	if err != nil {
		log.Fatal(err)
	}

	// nsenter기능하는 자식 프로세스 실행
	cmd := exec.Command("/proc/self/exe", "nsenter")
	cmd.Env = append(os.Environ(), fmt.Sprintf("target_pid=%d", info.Pid),
	fmt.Sprintf("target_cmd=%s", command))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("로그: 컨테이너 접속 실패: %v\n", err)
	}
} 

func stop() {
	if len(os.Args) < 3 {
		log.Fatal("사용법: stop [컨테이너ID]")
	}
	containerID := os.Args[2]
	if err := container.StopContainer(containerID); err != nil {
		log.Fatalf("로그: 멈추기 실패: %v", err)
	}
}
