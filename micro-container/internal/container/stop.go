package container

import (
	"fmt"
	"micro-container/internal/cgroups"
	"micro-container/internal/storage"
	"os"
	"path/filepath"
	"syscall"
)

func StopContainer(containerID string) error {
	// 컨테이너 정보 읽고 PID도출하기
	info, err := ReadContainerInfo(containerID)
	if err != nil {
		return fmt.Errorf("로그: 컨테이너 정보를 읽을 수 없습니다: %v", err)
	}

	// 프로세스 종료시키기(SIGTERM으로)
	fmt.Printf("로그: 컨테이너 프로세스(PID: %d) 종료 중...\n", info.Pid)
	if err := syscall.Kill(info.Pid, syscall.SIGKILL); err != nil {
		fmt.Printf("경고: 프로세스 종료 실패(이미 죽었을 수 있음): %v\n", err)
	}

	// 자원 회수
	fmt.Println("로그: 자원 회수 시작...")

	cgroups.RemoveCgroup(containerID)
	storage.UnmountOverlay(containerID)
	dirPath := filepath.Join(InfoLocation, containerID)
	os.RemoveAll(dirPath)

	fmt.Printf("로그: 컨테이너 %s가 성공적으로 제거되었습니다.\n", containerID)
	return nil

}
