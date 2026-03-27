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

	// /proc 마운트
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
		return fmt.Errorf("proc 마운트 실패: %v", err)
	}

	// 프로그램 실행
	if err := syscall.Exec(args[0], args, os.Environ()); err != nil {
		return fmt.Errorf("프로그램 실행 실패: %v", err)
	}

	return nil
}
