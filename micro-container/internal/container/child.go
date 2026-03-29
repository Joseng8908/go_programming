package container

import (
	"fmt"
	"micro-container/internal/network"
	"os"
	"syscall"
	"time"
)

func Child(args []string) error { 
	fmt.Printf("자식 프로세스(컨테이너) 시작 (PID: %d)", os.Getpid())

	containerID := "test-container"
	mergedDir := fmt.Sprintf("./tmp/%s/merged", containerID)
	pivotRoot(mergedDir)

	// 바뀐 root 안에 있는 proc폴더에 가상 파일시스템 올리기
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return fmt.Errorf("proc 마운트 실패: %v", err)
	}

	// 부모가 veth형성하고 UP시킬때까지 기다리기
	time.Sleep(100 * time.Millisecond)
	// 호스트 네임 설정하기
	if err := syscall.Sethostname([]byte("micro-container")); err != nil {
		return fmt.Errorf("호스트 네임 설정 실패: %v", err)
	}

	// 네트워크 설정하기
	if err := network.ConfigureContainerNetwork("eth0-temp"); err != nil {
		return fmt.Errorf("컨테이너 네트워크 설정 실패: %v", err)
	}

	fmt.Printf("log: 지금 실행하려는 파일: %s\n", args[0])
	// 프로그램 실행
	if err := syscall.Exec(args[0], args, os.Environ()); err != nil {
		return fmt.Errorf("프로그램 실행 실패: %v", err)
	}

	return nil
}

func pivotRoot(newRoot string) {
	// 해당 폴더로 이동하기
	syscall.Chdir(newRoot)

	// chroot 실행
	syscall.Chroot(".")

	// 다시 루트로 이돟하기
	syscall.Chdir("/")
}
