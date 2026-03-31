package container

import (
	"fmt"
	"log"
	"micro-container/internal/cgroups"
	"micro-container/internal/network"
	"micro-container/internal/storage"
	"time"
)

// 컨테이너 프로세스 실행 로직
func (c *Container) Run(detach bool) error {
	// 마운트 하기
	if _, err := storage.MountOverlay(c.ID); err != nil {
		fmt.Printf("마운트 에러: %v\n", err)
	}
	fmt.Println("마운트 완료")

	// 브릿지 이름 생성하고, 브릿지 만들기
	bridgeName := "mybridge"
	network.GetOrCreateBridge(bridgeName)
	
	if err := c.Cmd.Start(); err != nil {
		return fmt.Errorf("자식 프로세스 시작 실패: %v", err)
	}
	pid := c.Cmd.Process.Pid

	// cgroup관련 파일 만들고, cgroup지우는거 defer로 확정하기
	cgroups.SetCgroup(c.ID, pid)

	// Veth세팅하고, 호스트(브릿지 쪽 올리기)
	// 아직 컨테이너 쪽은 모름
	if err := network.SetupVeth(c.Cmd.Process.Pid, bridgeName); err != nil {
		return fmt.Errorf("Veth 설정 실패: %v", err)
	}

	info := ContainerInfo{
		ID: c.ID,
		Pid: c.Cmd.Process.Pid,
		Command: "/bin/sh",
		CreateTime: time.Now(),
		Status: "running",
	}

	if err := WriteContainerInfo(&info); err != nil {
		log.Printf("로그: 컨테이너 정보 기록 실패: %v", err)
	}


	// 부모 설정 끝, 자식에게 신호 보내기
	fmt.Println("부모: 모든 설정 완료. 자식에게 신호를 보냅니다.")
	c.SyncPipe.Write([]byte("go"))
	c.SyncPipe.Close()

	// 백그라운드 설정이 on이면 이렇게 실행
	if detach {
		fmt.Printf("로그: 컨테이너 %s를 백그라운드로 전환합니다.\n", c.ID)
		return nil
	}

	defer storage.UnmountOverlay(c.ID)
	defer cgroups.RemoveCgroup(c.ID)
	
	// 컨테이너 프로세스가 끝날때까지 대기하기
	err := c.Cmd.Wait()
	if err != nil {
		fmt.Printf("컨테이너 에러: %v\n", err)
	} else {
		fmt.Println("컨테이너 종료")
	}
	return nil
}
