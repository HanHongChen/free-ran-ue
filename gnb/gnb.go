package gnb

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Alonza0314/free-ran-ue/logger"
	"github.com/Alonza0314/free-ran-ue/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
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

	gnbId   []byte
	gnbName string

	plmnId ngapType.PLMNIdentity
	tai    ngapType.TAI
	snssai ngapType.SNSSAI

	logger *logger.Logger
}

func NewGnb(config *model.GnbConfig, logger *logger.Logger) *Gnb {
	gnbId, err := util.HexStringToBytes(config.Gnb.GnbId)
	if err != nil {
		logger.Error("GNB", fmt.Sprintf("Error converting gnbId to escaped: %v", err))
		return nil
	}
	plmnId := ngapConvert.PlmnIdToNgap(models.PlmnId{
		Mcc: config.Gnb.PlmnId.Mcc,
		Mnc: config.Gnb.PlmnId.Mnc,
	})
	tai := ngapConvert.TaiToNgap(models.Tai{
		Tac: config.Gnb.Tai.Tac,
		PlmnId: &models.PlmnId{
			Mcc: config.Gnb.Tai.BroadcastPlmnId.Mcc,
			Mnc: config.Gnb.Tai.BroadcastPlmnId.Mnc,
		},
	})
	sstInt, err := strconv.Atoi(config.Gnb.Snssai.Sst)
	if err != nil {
		logger.Error("GNB", fmt.Sprintf("Error converting sst to int: %v", err))
		return nil
	}
	snssai := ngapConvert.SNssaiToNgap(models.Snssai{
		Sst: int32(sstInt),
		Sd:  config.Gnb.Snssai.Sd,
	})

	return &Gnb{
		amfN2Ip: config.Gnb.AmfN2Ip,
		gnbN2Ip: config.Gnb.GnbN2Ip,

		amfN2Port: config.Gnb.AmfN2Port,
		gnbN2Port: config.Gnb.GnbN2Port,

		ngapPpid: config.Gnb.NgapPpid,

		gnbId:   gnbId,
		gnbName: config.Gnb.GnbName,

		plmnId: plmnId,
		tai:    tai,
		snssai: snssai,

		logger: logger,
	}
}

func (g *Gnb) Start() error {
	g.logger.Info("GNB", "Starting GNB")
	if err := g.connectToAmf(); err != nil {
		g.logger.Error("GNB", err.Error())
		return err
	}

	if err := g.setupN2(); err != nil {
		g.logger.Error("GNB", fmt.Sprintf("Error setting up N2: %v", err))
		return err
	}

	g.PrintBasicInfo()

	g.logger.Info("GNB", "GNB started")
	return nil
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

func (g *Gnb) setupN2() error {
	g.logger.Info("GNB", "Setting up N2")
	if err := ngapSetup(g.n2Conn, g.gnbId, g.gnbName, g.plmnId, g.tai, g.snssai); err != nil {
		g.logger.Error("GNB", fmt.Sprintf("Error setting up N2: %v", err))
		return err
	}
	g.logger.Info("GNB", "N2 setup complete")
	return nil
}

func (g *Gnb) PrintBasicInfo() {
	g.logger.Info("GNB", "========== gNB Basic Info ==========")

	gnbId := util.BytesToHexString(g.gnbId)
	g.logger.Info("GNB", fmt.Sprintf("gNB ID: %v, name: %s", gnbId, g.gnbName))

	plmnId := ngapConvert.PlmnIdToModels(g.plmnId)
	g.logger.Info("GNB", fmt.Sprintf("PLMN ID: %v", plmnId))

	tai := ngapConvert.TaiToModels(g.tai)
	g.logger.Info("GNB", fmt.Sprintf("TAC: %v, broadcast PLMN ID: %v", tai.Tac, tai.PlmnId))

	snssai := ngapConvert.SNssaiToModels(g.snssai)
	g.logger.Info("GNB", fmt.Sprintf("SST: %v, SD: %v", snssai.Sst, snssai.Sd))

	g.logger.Info("GNB", "====================================")
}
