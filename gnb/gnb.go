package gnb

import (
	"errors"
	"fmt"

	"github.com/Alonza0314/free-ran-ue/logger"
	"github.com/Alonza0314/free-ran-ue/model"
	"github.com/free5gc/sctp"
)

type Gnb struct {
	amfN2Ip string
	gnbN2Ip string

	amfN2Port int
	gnbN2Port int

	gnbULTeid string
	gnbDLTeid string

	n2Conn *sctp.SCTPConn

	ngapPpid uint32

	logger *logger.Logger
}

func NewGnb(config *model.GnbConfig, logger *logger.Logger) *Gnb {
	return &Gnb{
		amfN2Ip: config.Gnb.AmfN2Ip,
		gnbN2Ip: config.Gnb.GnbN2Ip,

		amfN2Port: config.Gnb.AmfN2Port,
		gnbN2Port: config.Gnb.GnbN2Port,

		ngapPpid: config.Gnb.NgapPpid,

		logger: logger,
	}
}

func (g *Gnb) Start() {
	g.logger.Info("GNB", "Starting GNB")
	if err := g.connectToAmf(); err != nil {
		g.logger.Error("GNB", err.Error())
		return
	}

	g.logger.Info("GNB", "GNB started")
}

func (g *Gnb) Stop() {
	g.logger.Info("GNB", "Stopping GNB")
	if err := g.n2Conn.Close(); err != nil {
		g.logger.Error("GNB", fmt.Sprintf("Error stopping GNB: %v", err))
		return
	}
	g.logger.Info("GNB", "GNB stopped")
}

func (g *Gnb) connectToAmf() error {
	g.logger.Info("GNB", "Connecting to AMF")

	amfAddr, gnbAddr, err := getAmfAndGnbSctpN2Addr(g.amfN2Ip, g.gnbN2Ip, g.amfN2Port, g.gnbN2Port)
	if err != nil {
		return err
	}
	g.logger.Debug("GNB", fmt.Sprintf("AMF N2 Address: %v", amfAddr.String()))
	g.logger.Debug("GNB", fmt.Sprintf("GNB N2 Address: %v", gnbAddr.String()))

	conn, err := sctp.DialSCTP("sctp", gnbAddr, amfAddr)
	if err != nil {
		return errors.New(fmt.Sprintf("Error connecting to AMF: %v", err))
	}

	info, err := conn.GetDefaultSentParam()
	if err != nil {
		return err
	}
	info.PPID = g.ngapPpid
	if err := conn.SetDefaultSentParam(info); err != nil {
		return errors.New(fmt.Sprintf("Error setting default sent param: %v", err))
	}

	g.n2Conn = conn

	g.logger.Info("GNB", fmt.Sprintf("Connected to AMF: %v", amfAddr.String()))
	return nil
}
