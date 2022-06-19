package net

import (
	"errors"
	"net"
)

func PublicIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var interface_ip string
	for _, iface := range ifaces {
		addr, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		if iface.Flags&net.FlagMulticast == net.FlagMulticast &&
			iface.Flags&net.FlagUp == net.FlagUp {
			if len(addr) < 1 {
				return "", errors.New("could not determine IP")
			}
			a, _, err := net.ParseCIDR(addr[0].String())
			if err != nil {
				return "", errors.New("could not extract IP")
			}
			interface_ip = a.String()
		}
	}
	return interface_ip, nil
}
