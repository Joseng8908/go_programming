package network

import (
	"strconv"
	"github.com/vishvananda/netlink"
)

func SetupNetwork(pid int) error {
	// veth pair 이름 정의하기
	hostVethName := "veth-host-" + strconv.Itoa(pid)
	containerVethName := "eth0" // 컨테이너 안에서는 이더넷처럼 행동하므로 이더넷을 붙임
	
	// veth pair 생성하기, link add 하는거임
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{Name: hostVethName},
		PeerName: containerVethName,
	}
	netlink.LinkAdd(veth)

	// 컨테이너 쪽 랜선을 네임스페이스로 이동시키기, link set하는거임
	peer, _ := netlink.LinkByName(containerVethName)
	netlink.LinkSetNsPid(peer, pid)

	// 호스토 쪽 설정하기, ip부여하고 up시키기
	hostVeth, _ := netlink.LinkByName(hostVethName)
	addr, _ := netlink.ParseAddr("172.17.0.1/24") //이러면 안되는거 아닌가?
	netlink.AddrAdd(hostVeth, addr)
	netlink.LinkSetUp(hostVeth)

	return nil
}
