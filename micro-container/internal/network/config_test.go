package network

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
)

func TestConfigureContainerNetwork(t *testing.T) {
	// ⚠️ 중요: 네임스페이스 변경은 현재 스레드에만 적용되므로 스레드를 고정합니다.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	bridgeName := "test-br0"
	
	// 1. 브릿지 준비
	GetOrCreateBridge(bridgeName)
	defer cleanupLink(bridgeName)

	// 2. 격리된 프로세스 실행 (CLONE_NEWNET 필수!)
	cmd := exec.Command("sleep", "10")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET,
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("테스트 프로세스 실행 실패: %v", err)
	}
	pid := cmd.Process.Pid
	defer cmd.Process.Kill()

	// 3. veth 생성 및 컨테이너로 이동 (SetupVeth 활용)
	// 이전 코드에서 "eth0-temp"를 썼다면 여기서도 맞춰줘야 합니다.
	// 여기서는 편의상 SetupVeth가 peer 이름을 "eth0"로 바꿨다고 가정하거나 
	// 테스트용으로 수동 생성합니다.
	setupTestVeth(pid, bridgeName)
	defer cleanupLink(fmt.Sprintf("veth%d", pid))

	time.Sleep(100 * time.Millisecond)
	// 4. 테스트 대상 함수 실행
	t.Run("ConfigureIPAndUp", func(t *testing.T) {
		err := ConfigureContainerNetwork(pid)
		if err != nil {
			t.Fatalf("네트워크 설정 실패: %v", err)
		}

		// 5. 검증: 실제로 컨테이너 네임스페이스에 들어가서 확인
		err = verifyNetworkInNamespace(pid)
		if err != nil {
			t.Errorf("검증 실패: %v", err)
		}
	})
}

// 컨테이너 네임스페이스 내부를 들여다보는 검증 함수
func verifyNetworkInNamespace(pid int) error {
	// 잠시 컨테이너 안으로 들어감
	newns, _ := netns.GetFromPid(pid)
	origns, _ := netns.Get()
	defer netns.Set(origns)
	netns.Set(newns)

	// eth0 찾기
	link, err := netlink.LinkByName("eth0-temp")
	if err != nil {
		return fmt.Errorf("컨테이너 내 eth0-temp 없음: %v", err)
	}

	// IP 확인
	addrs, _ := netlink.AddrList(link, netlink.FAMILY_V4)
	if len(addrs) == 0 || addrs[0].IP.String() != "172.17.0.2" {
		return fmt.Errorf("IP 설정 오류: %v", addrs)
	}

	// 상태 확인
	if link.Attrs().Flags&net.FlagUp == 0 {
		return fmt.Errorf("eth0-temp가 여전히 DOWN 상태임")
	}

	return nil
}

// 헬퍼: 테스트용 veth 세팅 (Peer 이름을 "eth0"로 생성)
func setupTestVeth(pid int, bridgeName string) {
	la := netlink.NewLinkAttrs()
	la.Name = fmt.Sprintf("veth%d", pid)
	veth := &netlink.Veth{
		LinkAttrs: la,
		PeerName:  "eth0-temp", // ConfigureContainerNetwork가 찾는 이름
	}
	netlink.LinkAdd(veth)
	
	br, _ := netlink.LinkByName(bridgeName)
	netlink.LinkSetMaster(veth, br)
	netlink.LinkSetUp(veth)

	peer, _ := netlink.LinkByName("eth0-temp")
	netlink.LinkSetNsPid(peer, pid) // Peer가 자동으로 pid 네임스페이스로 이동
}
