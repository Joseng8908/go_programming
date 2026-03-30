package storage

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func MountOverlay(containerID string) (string, error) {
	// 필요한 폴더 경로 설정하기
	basePath := fmt.Sprintf("./tmp/%s", containerID)
	lowerDir := "./rootfs" //원본 이미지(ReadOnly)
	uppderDir := basePath + "/upper"
	workDir := basePath + "/work"
	mergedDir := basePath + "/merged"

	// 폴더 생성하기
	for _, dir := range []string{uppderDir, workDir, mergedDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}
	fmt.Print("폴더 생성 완료")

	// 마운트 옵션 문자열 생서하기
	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, uppderDir, workDir)
	
	// syscall.Moun호출하기
	// 형식: syscall.Mount(source, target, fstype, flags, data)
	err := syscall.Mount("overlay", mergedDir, "overlay", 0, opts)
	if err != nil {
		return "", fmt.Errorf("OverlayFS 마운트 실패: %v", err)
	}
	return mergedDir, nil
}

func UnmountOverlay(containerID string) error {
	mergedDir := fmt.Sprintf("./tmp/%s/merged", containerID)
	basePath := fmt.Sprintf("./tmp/%s", containerID)

	// 마운트 해제하기
	if err := syscall.Unmount(mergedDir, 0); err != nil {
		return fmt.Errorf("마운트 해제 실패: %v", err)
	}

	// 언마운트 시간 벌기
	time.Sleep(100 * time.Millisecond)

	// 임시 폴더 삭제(upper, work, merged가 담긴 부모 폴더를 그냥 삭제해버리기)
	return os.RemoveAll(basePath)
}
