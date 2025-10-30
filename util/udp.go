package util

import (
	"fmt"
	"net"
)

func UdpDialWithOptionalLocalAddress(remoteAddress string, remotePort int, localAddress string) (net.Conn, error) {
	if localAddress == "" {
		return net.Dial("udp", fmt.Sprintf("%s:%d", remoteAddress, remotePort))
	}

	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", localAddress, 0))
	if err != nil {
		return nil, fmt.Errorf("error resolving local address: %v", err)
	}

	dialer := &net.Dialer{
		LocalAddr: localAddr,
	}
	return dialer.Dial("udp", fmt.Sprintf("%s:%d", remoteAddress, remotePort))
}
