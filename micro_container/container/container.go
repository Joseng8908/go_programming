package container

import "os/exec"

// container 정의
// export로
type Container struct {
	ID string
	Cmd *exec.Cmd
	Status string
}

// 새로운 container instance를 생성하는 생성자
func NewContainer(id string, command string, args ...string) *Container {
	cmd := exec.Command(command, args...)
	// TODO: linux Namespace 설정 넣기
	return &Container{
		ID: id,
		Cmd: cmd,
		Status: "Created",
	}
}
