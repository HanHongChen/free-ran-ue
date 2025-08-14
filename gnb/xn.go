package gnb

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
)

type XnPdu struct {
	ImsiLength uint16
	Imsi       string
	Data       []byte
}

func NewXnPdu(imsi string, data []byte) *XnPdu {
	return &XnPdu{
		ImsiLength: 0,
		Imsi:       imsi,
		Data:       data,
	}
}

func (x *XnPdu) Marshal() ([]byte, error) {
	imsiBytes := []byte(x.Imsi)

	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, uint16(len(imsiBytes)))

	buffer = append(buffer, imsiBytes...)
	buffer = append(buffer, x.Data...)

	return buffer, nil
}

func (x *XnPdu) Unmarshal(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("data too short")
	}

	x.ImsiLength = binary.BigEndian.Uint16(data[:2])
	data = data[2:]

	if len(data) < int(x.ImsiLength) {
		return fmt.Errorf("data too short")
	}

	x.Imsi = string(data[:x.ImsiLength])
	data = data[x.ImsiLength:]

	x.Data = data

	return nil
}

func xnInterfaceProcessor(conn net.Conn, g *Gnb) {
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		g.XnLog.Warnf("Error reading XN packet: %v", err)
		return
	}
	g.XnLog.Tracef("Received %d bytes of XN packet: %+v", n, buffer[:n])
	g.XnLog.Debugln("Receive XN packet")

	xnPdu := XnPdu{}
	if err := xnPdu.Unmarshal(buffer[:n]); err != nil {
		g.XnLog.Errorf("Error unmarshal xn pdu: %v", err)
		return
	}
	g.XnLog.Tracef("Received XN PDU: %+v", xnPdu)
	g.XnLog.Debugln("Receive XN PDU")

	ngapPduSessionResourceSetupRequest, err := ngap.Decoder(xnPdu.Data)
	if err != nil {
		g.XnLog.Warnf("Error decoding NGAP PDU Session Resource Setup Request: %v", err)
		return
	}

	if ngapPduSessionResourceSetupRequest.Present != ngapType.NGAPPDUPresentInitiatingMessage {
		g.XnLog.Warnf("Error NGAP PDU Session Resource Setup Request Present: %v, expected %v", ngapPduSessionResourceSetupRequest.Present, ngapType.NGAPPDUPresentInitiatingMessage)
		return
	}

	switch ngapPduSessionResourceSetupRequest.InitiatingMessage.ProcedureCode.Value {
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		g.XnLog.Debugln("Processing NGAP PDU Session Resource Setup Request")
		xnPduSessionResourceSetupRequestProcessor(g, conn, xnPdu.Imsi, ngapPduSessionResourceSetupRequest)
	default:
		g.XnLog.Warnf("Unknown NGAP PDU Session Resource Setup Request Procedure Code: %v", ngapPduSessionResourceSetupRequest.InitiatingMessage.ProcedureCode.Value)
		return
	}
}

func xnPduSessionResourceSetupRequestProcessor(g *Gnb, conn net.Conn, imsi string, ngapPduSessionResourceSetupRequest *ngapType.NGAPPDU) {
	var pduSessionResourceSetupRequestTransfer ngapType.PDUSessionResourceSetupRequestTransfer

	for _, ie := range ngapPduSessionResourceSetupRequest.InitiatingMessage.Value.PDUSessionResourceSetupRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
		case ngapType.ProtocolIEIDRANUENGAPID:
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq:
			for _, pduSessionResourceSetupItem := range ie.Value.PDUSessionResourceSetupListSUReq.List {
				if err := aper.UnmarshalWithParams(pduSessionResourceSetupItem.PDUSessionResourceSetupRequestTransfer, &pduSessionResourceSetupRequestTransfer, "valueExt"); err != nil {
					g.XnLog.Warnf("Error unmarshal pdu session resource setup request transfer: %v", err)
					return
				}
				g.XnLog.Tracef("Get PDUSessionResourceSetupRequestTransfer: %+v", pduSessionResourceSetupRequestTransfer)
			}
		case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
		}
	}

	xnUe := NewXnUe(g.teidGenerator.AllocateTeid(), conn)
	g.XnLog.Debugf("Allocated DLTEID for XnUe: %s", hex.EncodeToString(xnUe.GetDlTeid()))

	for _, ie := range pduSessionResourceSetupRequestTransfer.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDPDUSessionAggregateMaximumBitRate:
		case ngapType.ProtocolIEIDULNGUUPTNLInformation:
		case ngapType.ProtocolIEIDAdditionalULNGUUPTNLInformation:
			xnUe.SetUlTeid(ie.Value.AdditionalULNGUUPTNLInformation.List[0].NGUUPTNLInformation.GTPTunnel.GTPTEID.Value)
		case ngapType.ProtocolIEIDPDUSessionType:
		case ngapType.ProtocolIEIDQosFlowSetupRequestList:
		}
	}

	// DC QoS Flow per TNL Information
	dcQosFlowPerTNLInformationItem := ngapType.QosFlowPerTNLInformationItem{}
	dcQosFlowPerTNLInformationItem.QosFlowPerTNLInformation.UPTransportLayerInformation.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel

	// DC Transport Layer Information in QoS Flow per TNL Information
	dcUpTransportLayerInformation := &dcQosFlowPerTNLInformationItem.QosFlowPerTNLInformation.UPTransportLayerInformation
	dcUpTransportLayerInformation.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	dcUpTransportLayerInformation.GTPTunnel = new(ngapType.GTPTunnel)
	dcUpTransportLayerInformation.GTPTunnel.GTPTEID.Value = xnUe.GetDlTeid()
	dcUpTransportLayerInformation.GTPTunnel.TransportLayerAddress = ngapConvert.IPAddressToNgap(g.ranN3Ip, "")

	// DC Associated QoS Flow List in QoS Flow per TNL Information
	dcAssociatedQosFlowList := &dcQosFlowPerTNLInformationItem.QosFlowPerTNLInformation.AssociatedQosFlowList
	dcAssociatedQosFlowItem := ngapType.AssociatedQosFlowItem{}
	dcAssociatedQosFlowItem.QosFlowIdentifier.Value = 1
	dcAssociatedQosFlowList.List = append(dcAssociatedQosFlowList.List, dcAssociatedQosFlowItem)

	dcQosFlowPerTNLInformationMarshal, err := aper.MarshalWithParams(dcQosFlowPerTNLInformationItem, "valueExt")
	if err != nil {
		g.XnLog.Warnf("Error marshal dc qos flow per tnl information: %v", err)
		return
	}

	xnPdu := NewXnPdu(imsi, dcQosFlowPerTNLInformationMarshal)
	xnPduBytes, err := xnPdu.Marshal()
	if err != nil {
		g.XnLog.Warnf("Error marshal xn pdu: %v", err)
		return
	}

	n, err := conn.Write(xnPduBytes)
	if err != nil {
		g.XnLog.Warnf("Error write dc qos flow per tnl information: %v", err)
		return
	}
	g.XnLog.Tracef("Sent %d bytes of DC QoS Flow per TNL Information to XN", n)
	g.XnLog.Debugln("Send DC QoS Flow per TNL Information to XN")

	ueDataPlaneConn, err := (*g.ranDataPlaneListener).Accept()
	if err != nil {
		g.XnLog.Warnf("Error accept ue data plane connection: %v", err)
		return
	}
	xnUe.SetDataPlaneConn(ueDataPlaneConn)
	g.XnLog.Infof("Accepted UE data plane connection from: %v", ueDataPlaneConn.RemoteAddr())
	g.teidToConn.Store(hex.EncodeToString(xnUe.GetDlTeid()), xnUe.GetDataPlaneConn())
	g.XnLog.Debugf("Stored UE data plane connection with teid %s to teidToConn", hex.EncodeToString(xnUe.GetDlTeid()))

	go g.startUeDataPlaneProcessor(ueDataPlaneConn, xnUe.GetUlTeid(), xnUe.GetDlTeid())
}
