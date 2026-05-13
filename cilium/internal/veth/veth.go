package veth

import (
	"github.com/vishvananda/netlink"
)

type config struct {
	host_ifname string
	peer_ifname string
	host_ip string
	peer_ip string
}

func veth(name string, vethConfig *config) error{
	LinkAdd
} 
