package container

import "fmt"

// 컨테이너 프로세스 실행 로직
func (c *Container) Start() error {
	if err := c.Cmd.Start(); err != nil {
		return err
	}

	// 확인용 print
	fmt.Printf("컨테이너 ID: %s, PID: %d 로 실행됨\n", c.ID, c.Cmd.Process.Pid)

	return nil
}
