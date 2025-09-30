package util

import (
	"fmt"
	"net"
)

func TcpDialWithOptionalLocalAddress(remoteAddress string, remotePort int, localAddress string) (net.Conn, error) {
	if localAddress == "" {
		return net.Dial("tcp", fmt.Sprintf("%s:%d", remoteAddress, remotePort))
	}

	// port 0 means to use any available port
	localAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localAddress, 0))
	if err != nil {
		return nil, fmt.Errorf("error resolving local address: %v", err)
	}

	dialer := &net.Dialer{
		LocalAddr: localAddr,
	}
	return dialer.Dial("tcp", fmt.Sprintf("%s:%d", remoteAddress, remotePort))
}
