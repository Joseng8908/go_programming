package container

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// container 정의
// export로
type Container struct {
	ID string
	Cmd *exec.Cmd
	Status string
	Ctx context.Context
	Cancel context.CancelFunc
}

// 새로운 container instance를 생성하는 생성자
func NewContainer(id string, command string, args ...string) *Container {
	// 여기서 커맨드가 실행 되겠죠?, 즉 여기서 컨테이너가 만들어지는 거임
	cmd := exec.Command(command, args...)

	// 만들어진 cmd객체의 표준 입출력을, 터미널과 연결
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 그럼 만들어졌다고 여기서 context 보내기
	ctx, cancel := context.WithCancel(context.Background())
	// 격리 설정..., 보통 exec를 통해 새로운 프로세스를 만들면
	// 기존 권한이나 파일 시스템 등등을 복사해서 만듦
	// 하지만 cmd.SysProcAttr을 사용한다면, 복사해서 만드는 것이 아닌
	// 자신만의 방을 만들어서 쓰겠다는 것임
	// CLONE_<flag>, 이건 클론 플래그라고 무엇을 격리할 것인지 쓰는 것임
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // 호스트 명 격리
					syscall.CLONE_NEWPID | // PID 격리(내부에서 PID 1번이 됨)
					syscall.CLONE_NEWNS, // 마운트 격리
		Chroot: "./roofs", // 이게 컨테이너(격리된 프로세스의 입구)
	}
	cmd.Dir = "/" // 프로세스가 시작될 위치를 컨테이너 내부의 루트로 고정하기

	// container에 대한 정보 객체 반환 
	return &Container{
		ID: id,
		Cmd: cmd,
		Status: "Created",
		Ctx: ctx, 
		Cancel: cancel,
	}
}

// 컨테이너 프로세스 실행 로직
func (c *Container) Start() error {
	if err := c.Cmd.Start(); err != nil {
		return err
	}

	// 확인용 print
	fmt.Printf("컨테이너 ID: %s, PID: %d 로 실행됨\n", c.ID, c.Cmd.Process.Pid)

	return nil
}

// 파일 시스템 변경 method
func (c *Container) setupRootFS() error {
	syscall.Chroot("./rootfs")
	os.Chdir("/")
	return nil
}
