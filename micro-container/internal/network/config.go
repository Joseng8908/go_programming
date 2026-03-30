package network

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

// 여기서 pid는 container의 pid
func ConfigureContainerNetwork(ifaceName string) error {
	// 이미 네임스페이스 안에서 실행하므로 호스트에서 들어갈 열쇠 구하고,,,등등의 과정 필요 없음
	// 컨테이너 안에서 세팅하기
	var eth0 netlink.Link
    var err error

    eth0, err = netlink.LinkByName(ifaceName)

    if err != nil {
        return fmt.Errorf("결국 장치를 찾지 못했습니다: %v", err)
    }

	// 컨테이너에게 ip주소 부여하기(1712.17.0.2)
	addr, _ := netlink.ParseAddr("172.17.0.2/24")
	if err := netlink.AddrAdd(eth0, addr); err != nil {
		return err
	}

	// eth0 UP시키기
	if err := netlink.LinkSetUp(eth0); err != nil {
		return err
	}

	lo, _ := netlink.LinkByName("lo")
	netlink.LinkSetUp(lo)
	return nil
}
