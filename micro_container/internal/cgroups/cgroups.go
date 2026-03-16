package cgroups

import (
	"os"
	"strconv"
)
func setCgroup(containerID string, pid int) error {
	cgroupPath := "/sys/fs/cgroup/container-" + containerID
	
	// sudo 로 실행되기 때문에 차피 다 권한이 있음

	// 폴더 생성(cgroup가상 파일시스템 폴더 생성, 이거 생성되면 안에서 알아서
	// 다른 것들도 만들어짐)
	os.Mkdir(cgroupPath, 0755)

	// 메모리 제한 설정하기, 질리게 했던거임
	os.WriteFile(cgroupPath+"/memory.max", []byte("20000000"), 0644)

	// 프로세스 등록하기
	err := os.WriteFile(cgroupPath+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644)
	return err
}
