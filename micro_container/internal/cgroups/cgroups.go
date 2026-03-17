package cgroups

import (
	"os"
	"strconv"
	"fmt"
)
func SetCgroup(containerID string, pid int) error {
	cgroupPath := "/sys/fs/cgroup/container-" + containerID
	
	// sudo 로 실행되기 때문에 차피 다 권한이 있음

	// 폴더 생성(cgroup가상 파일시스템 폴더 생성, 이거 생성되면 안에서 알아서
	// 다른 것들도 만들어짐)
	if err := os.Mkdir(cgroupPath, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	fmt.Println("로그: cgroup파일 만들어짐, 경로: %s", cgroupPath)
	// 메모리 제한 설정하기, 질리게 했던거임
	os.WriteFile(cgroupPath+"/memory.max", []byte("20000000"), 0644)

	// 스왑 제한하기
	os.WriteFile(cgroupPath + "/memory.swap.max", []byte("0"), 0644)

	// 프로세스 등록하기
	err := os.WriteFile(cgroupPath+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644)
	return err
}
 
// 파일 삭제를 위한 별도 함수
func RemoveCgroup(containerID string) error {
	cgroupPath := "/sys/fs/cgroup/container-" + containerID
	return os.Remove(cgroupPath)
}
