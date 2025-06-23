package gnb

import (
	"errors"
	"fmt"
	"net"

	"github.com/free5gc/sctp"
)

func getAmfAndGnbSctpN2Addr(amfN2Ip, gnbN2Ip string, amfN2Port, gnbN2Port int) (*sctp.SCTPAddr, *sctp.SCTPAddr, error) {
	amfIps := make([]net.IPAddr, 0)
	gnbIps := make([]net.IPAddr, 0)

	if ip, err := net.ResolveIPAddr("ip", amfN2Ip); err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Error resolving AMF N2 IP address '%s': '%v'", amfN2Ip, err))
	} else {
		amfIps = append(amfIps, *ip)
	}
	amfAddr := &sctp.SCTPAddr{
		IPAddrs: amfIps,
		Port:    amfN2Port,
	}

	if ip, err := net.ResolveIPAddr("ip", gnbN2Ip); err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Error resolving GNB N2 IP address '%s': '%v'", gnbN2Ip, err))
	} else {
		gnbIps = append(gnbIps, *ip)
	}
	gnbAddr := &sctp.SCTPAddr{
		IPAddrs: gnbIps,
		Port:    gnbN2Port,
	}

	return amfAddr, gnbAddr, nil
}
