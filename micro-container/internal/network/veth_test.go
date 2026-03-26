package network

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"testing"

	"github.com/vishvananda/netlink"
)

func TestSetupVeth(t *testing.T) {
	// OS 스레드를 고정하여 네임스페이스 혼선을 방지
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	bridgeName := "test-br0"
	
	// 1. 사전 준비: 테스트용 브릿지 생성
	err := GetOrCreateBridge(bridgeName)
	if err != nil {
		t.Fatalf("테스트용 브릿지 생성 실패: %v", err)
	}
	defer cleanupLink(bridgeName)

	// 2. 테스트용 프로세스 실행 (잠시 대기하는 프로세스)
	// 이 프로세스의 PID를 컨테이너의 PID인 것처럼 사용합니다.
	cmd := exec.Command("sleep", "10")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET,
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("테스트 프로세스 실행 실패: %v", err)
	}
	pid := cmd.Process.Pid
	defer cmd.Process.Kill()

	// veth 이름 정의 (함수 내부 로직과 동일하게)
	hostVethName := fmt.Sprintf("veth%d", pid)
	defer cleanupLink(hostVethName)

	// 3. SetupVeth 실행
	t.Run("SetupVethConnectivity", func(t *testing.T) {
		err := SetupVeth(pid, bridgeName)
		if err != nil {
			t.Fatalf("SetupVeth 실행 실패: %v", err)
		}

		// 검증 1: 호스트 쪽에 veth가 생성되었는지 확인
		hostVeth, err := netlink.LinkByName(hostVethName)
		if err != nil {
			t.Errorf("호스트 veth(%s)를 찾을 수 없음: %v", hostVethName, err)
		}

		// 검증 2: 호스트 veth가 브릿지에 제대로 연결(Master)되었는지 확인
		if hostVeth.Attrs().MasterIndex <= 0 {
			t.Error("호스트 veth가 브릿지에 연결되지 않았습니다.")
		}

		// 검증 3: 컨테이너 쪽 veth(eth0)가 호스트에서 사라졌는지 확인
		// (네임스페이스를 이동했다면 호스트의 기본 네임스페이스에서는 조회되지 않아야 함)
		_, err = netlink.LinkByName("eth0")
		if err == nil {
			t.Error("eth0가 여전히 호스트 네임스페이스에 존재합니다. 이동 실패!")
		}
	})
}

// 테스트용 헬퍼 함수: 장치 삭제
func cleanupLink(name string) {
	link, err := netlink.LinkByName(name)
	if err == nil {
		netlink.LinkDel(link)
	}
}

