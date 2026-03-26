package network

import (
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
)

// 여기서 pid는 container의 pid
func ConfigureContainerNetwork(pid int) error {
	// 현재는 호스트에 존재
	// 호스트의 주소를 미리 저장해두고
	// 컨테이너에 갔다가 다시 돌아올때 이 주소를 사용
	origns, _ := netns.Get()
	defer origns.Close()

	// 들어갈 컨테이너의 열쇠? 얻기
	targetns, err := netns.GetFromPid(pid)
	if err != nil {
		return err
	}
	defer targetns.Close()

	// 환경 세팅하기 위해 컨테이너로 들어가기
	if err := netns.Set(targetns); err != nil {
		return err
	}

	// 세팅 다 하고 다시 호스트로 돌아오도록 defer
	defer netns.Set(origns)

	// 컨테이너 안에서 세팅하기
	eth0, err := netlink.LinkByName("eth0-temp")
	if err != nil {
		return err
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

	return nil
}
