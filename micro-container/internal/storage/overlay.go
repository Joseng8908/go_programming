package storage

import (
	"fmt"
	"os"
	"syscall"
)

func MountOverlay(containerID string) (string, error) {
	// 필요한 폴더 경로 설정하기
	basePath := fmt.Sprint("/tmp/micro-docker/%s", containerID)
	lowerDir := "./rootfs" //원본 이미지(ReadOnly)
	uppderDir := basePath + "/upper"
	workDir := basePath + "/work"
	mergedDir := basePath + "merged"

	// 폴더 생성하기
	for _, dir := range []string{uppderDir, workDir, mergedDir} {
		if err := os.Mkdir(dir, 0755); err != nil {
			return "", err
		}
	}

	// 마운트 옵션 문자열 생서하기
	opts := fmt.Sprint("lowerdir=%s,upperdir%s,workdir=%s", lowerDir, uppderDir, workDir)
	
	// syscall.Moun호출하기
	// 형식: syscall.Mount(source, target, fstype, flags, data)
	err := syscall.Mount("overlay", mergedDir, "overlay", 0, opts)
	if err != nil {
		return "", fmt.Errorf("OverlayFS 마운트 실패: %v", err)
	}
	return mergedDir, nil
}

func UnmountOverlay(containerID string) error {
	mergedDir := fmt.Sprintf("/tmp/micro-docker/%s/merged", containerID)
	basePath := fmt.Sprint("/tmp/micro-docker/%s", containerID)

	// 마운트 해제하기
	if err := syscall.Unmount(mergedDir, 0); err != nil {
		return fmt.Errorf("마운트 해제 실패: %v", err)
	}

	// 임시 폴더 삭제(upper, work, merged가 담긴 부모 폴더를 그냥 삭제해버리기)
	return os.RemoveAll(basePath)
}
