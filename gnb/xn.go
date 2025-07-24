package gnb

import "net"

func xnInterfaceProcessor(conn net.Conn, g *Gnb) {
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		g.XnLog.Errorf("Error reading XN packet: %v", err)
		return
	}
	g.XnLog.Tracef("Received %d bytes of XN packet: %+v", n, buffer[:n])
	g.XnLog.Tracef("Received %d bytes of XN packet", n)
}
