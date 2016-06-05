package irc

import (
	"errors"
	"log"
	"net"
)

func hasLocalAddr(inp string) (net.Addr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		_, netw, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil, err
		}

		log.Println(netw.String())
		log.Println(netw.Network())

		if netw.Contains(net.ParseIP(inp)) {
			_, err := net.ResolveIPAddr("ip", inp)
			if err != nil {
				return nil, err
			}
			return &net.TCPAddr{
				IP: net.ParseIP(inp),
			}, nil
		}
	}

	return nil, errors.New("irc: Requested bindhost not found")
}
