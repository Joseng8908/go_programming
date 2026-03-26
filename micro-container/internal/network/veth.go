package network

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func SetupVeth(pid int, bridgeName string) error {
	// 호스트쪽 랜선? 이름 정의
	hostVethName := fmt.Sprintf("veth%d", pid)
	// 컨테이너 쪽에서는 eth0으로 보이게 정의
	containerVethName := "eth0" 

	// veth 쌍 설정하기,,,,
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{Name: hostVethName},
		PeerName: containerVethName,
	}

	// 커널에 랜선 만들자고 요청하기
	if err := netlink.LinkAdd(veth); err != nil {
		return fmt.Errorf("veth 생성 실패: %v", err)
	}

	// 브릿지에 끼울 랜선 가져오기
	hostVeth, _ := netlink.LinkByName(hostVethName)
	// 브릿지 장치 가져오기
	br, _ := netlink.LinkByName(bridgeName)

	// 호스트쪽 랜선 활성화 (UP시키기)
	if err := netlink.LinkSetMaster(hostVeth, br); err != nil {
		return fmt.Errorf("브릿지 쪽 veth 활성화 실패: %v", err)
	}

	// 컨테이너에 끼울 랜선 가져오기
	peer, err := netlink.LinkByName(containerVethName)
	if err != nil {
		return fmt.Errorf("컨테이너 veth 찾기 실패: %v", err)
	}

	// 컨테이너의 네임스페이스로 이동
	if err := netlink.LinkSetNsPid(peer, pid); err != nil {
		return fmt.Errorf("네임스페이스로 이동 실패: %v", err)
	}

	return nil


}
