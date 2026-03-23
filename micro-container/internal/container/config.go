package container

import (
	"context"
	"os/exec"
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
