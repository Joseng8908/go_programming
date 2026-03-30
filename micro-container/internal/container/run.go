package container

import (
	"fmt"
	"micro-container/internal/cgroups"
	"micro-container/internal/network"
	"micro-container/internal/storage"
)

// 컨테이너 프로세스 실행 로직
func (c *Container) Run() error {
	// 마운트 하기
	if _, err := storage.MountOverlay(c.ID); err != nil {
		fmt.Printf("마운트 에러: %v\n", err)
	}
	fmt.Println("마운트 완료")
	defer storage.UnmountOverlay(c.ID)

	// 브릿지 이름 생성하고, 브릿지 만들기
	bridgeName := "mybridge"
	network.GetOrCreateBridge(bridgeName)
	
	if err := c.Cmd.Start(); err != nil {
		return fmt.Errorf("자식 프로세스 시작 실패: %v", err)
	}
	pid := c.Cmd.Process.Pid

	// cgroup관련 파일 만들고, cgroup지우는거 defer로 확정하기
	cgroups.SetCgroup(c.ID, pid)
	defer cgroups.RemoveCgroup(c.ID)

	// Veth세팅하고, 호스트(브릿지 쪽 올리기)
	// 아직 컨테이너 쪽은 모름
	if err := network.SetupVeth(c.Cmd.Process.Pid, bridgeName); err != nil {
		return fmt.Errorf("Veth 설정 실패: %v", err)
	}

	// 부모 설정 끝, 자식에게 신호 보내기
	fmt.Println("부모: 모든 설정 완료. 자식에게 신호를 보냅니다.")
	c.SyncPipe.Write([]byte("go"))
	c.SyncPipe.Close()

	// 컨테이너 프로세스가 끝날때까지 대기하기
	err := c.Cmd.Wait()
	if err != nil {
		fmt.Printf("컨테이너 에러: %v\n", err)
	} else {
		fmt.Println("컨테이너 종료")
	}
	return nil
}
