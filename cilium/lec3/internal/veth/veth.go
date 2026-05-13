package veth

import (
	"os"
	"runtime"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

type config struct {
	host_ifname string
	peer_ifname string
	host_ip string
	peer_ip string
}

func veth(name string, vethConfig *config) error{
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name: vethConfig.host_ifname,
		},	
		PeerName: vethConfig.peer_ifname,
	}
	// veth pair 생성
	err := netlink.LinkAdd(veth)
	if err != nil {
		return err 
	}

	fd, err := os.Open("/var/run/netns/" + name)
	if err != nil {
		return err
	}
	defer fd.Close()

	// host ip addr세팅
	hostAddr, err := netlink.ParseAddr(vethConfig.host_ip)
	if err != nil {
		return err
	}
	netlink.AddrAdd(veth, hostAddr)

	// host쪽 up
	netlink.LinkSetUp(veth)



	// peer로 ns이동
	peer, _:= netlink.LinkByName(vethConfig.peer_ifname)
	netlink.LinkSetNsFd(peer, int(fd.Fd()))
	configPeerNs(name, vethConfig.peer_ifname, vethConfig.peer_ip)
	
	return nil
}   


func configPeerNs(name string, peerIfname string, peerIP string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	originNs, err := os.Open("/proc/self/ns/net")
	if err != nil {
		return err
	}
	defer originNs.Close()

	nsFd, err := os.Open("/var/run/netns/" + name)
	if err != nil {
		return err
	}
	defer nsFd.Close()

	unix.Setns(int(nsFd.Fd()), unix.CLONE_NEWNET)
	defer unix.Setns(int(originNs.Fd()), unix.CLONE_NEWNET)

	// peer ip 주소 할당
	peer, err := netlink.LinkByName(peerIfname)
	if err != nil {
		return err
	}
	peerAddr, err := netlink.ParseAddr(peerIP)
	if err != nil {
		return err
	}
	err = netlink.AddrAdd(peer, peerAddr)
	if err != nil {
		return err
	}
	// peer up
	netlink.LinkSetUp(peer)

	lo, _ := netlink.LinkByName("lo")
	netlink.LinkSetUp(lo)

	return nil
}
