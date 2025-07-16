package gnb

import (
	"reflect"
	"testing"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

var testBuildNgapSetupRequestCases = []struct {
	name    string
	gnbId   []byte
	gnbName string
	plmnId  ngapType.PLMNIdentity
	tai     ngapType.TAI
	snssai  ngapType.SNSSAI
}{
	{
		name:    "testBuildNgapSetupRequest",
		gnbId:   []byte("\x00\x03\x14"),
		gnbName: "gNB",
		plmnId: ngapType.PLMNIdentity{
			Value: aper.OctetString("\x02\xF8\x39"),
		},
		tai: ngapType.TAI{
			TAC: ngapType.TAC{
				Value: aper.OctetString("\x00\x00\x01"),
			},
			PLMNIdentity: ngapType.PLMNIdentity{
				Value: aper.OctetString("\x02\xF8\x39"),
			},
		},
		snssai: ngapType.SNSSAI{
			SST: ngapType.SST{
				Value: aper.OctetString("\x01"),
			},
			SD: &ngapType.SD{
				Value: aper.OctetString("\x01\x02\x03"),
			},
		},
	},
}

func TestBuildNgapSetupRequest(t *testing.T) {
	for _, testCase := range testBuildNgapSetupRequestCases {
		t.Run(testCase.name, func(t *testing.T) {
			pdu := buildNgapSetupRequest(testCase.gnbId, testCase.gnbName, testCase.plmnId, testCase.tai, testCase.snssai)
			encodeData, err := ngap.Encoder(pdu)
			if err != nil {
				t.Fatalf("Failed to encode NGAP setup request: %v", err)
			} else {
				decodeData, err := ngap.Decoder(encodeData)
				if err != nil {
					t.Fatalf("Failed to decode NGAP setup request: %v", err)
				} else if !reflect.DeepEqual(pdu, *decodeData) {
					t.Fatalf("NGAP setup request mismatch")
				}
			}
		})
	}
}

var testBuildIntialUeMessageCases = []struct {
	name                  string
	ranUeNgapId           int64
	ueRegistrationRequest []byte
	plmnId                ngapType.PLMNIdentity
	tai                   ngapType.TAI
}{
	{
		name:                  "testBuildIntialUeMessage",
		ranUeNgapId:           1,
		ueRegistrationRequest: []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
		plmnId: ngapType.PLMNIdentity{
			Value: aper.OctetString("\x02\xF8\x39"),
		},
		tai: ngapType.TAI{
			TAC: ngapType.TAC{
				Value: aper.OctetString("\x00\x00\x01"),
			},
			PLMNIdentity: ngapType.PLMNIdentity{
				Value: aper.OctetString("\x02\xF8\x39"),
			},
		},
	},
}

func TestBuildIntialUeMessage(t *testing.T) {
	for _, testCase := range testBuildIntialUeMessageCases {
		t.Run(testCase.name, func(t *testing.T) {
			pdu := buildInitialUeMessage(testCase.ranUeNgapId, testCase.ueRegistrationRequest, testCase.plmnId, testCase.tai)
			encodeData, err := ngap.Encoder(pdu)
			if err != nil {
				t.Fatalf("Failed to encode NGAP initial ue message: %v", err)
			} else {
				decodeData, err := ngap.Decoder(encodeData)
				if err != nil {
					t.Fatalf("Failed to decode NGAP initial ue message: %v", err)
				} else if !reflect.DeepEqual(pdu, *decodeData) {
					t.Fatalf("NGAP initial ue message mismatch")
				}
			}
		})
	}
}

var testBuildUplinkNasTransportCases = []struct {
	name        string
	amfUeNgapId int64
	ranUeNgapId int64
	plmnId      ngapType.PLMNIdentity
	tai         ngapType.TAI
	nasPdu      []byte
}{
	{
		name:        "testBuildUplinkNasTransport",
		amfUeNgapId: 1,
		ranUeNgapId: 1,
		plmnId: ngapType.PLMNIdentity{
			Value: aper.OctetString("\x02\xF8\x39"),
		},
		tai: ngapType.TAI{
			TAC: ngapType.TAC{
				Value: aper.OctetString("\x00\x00\x01"),
			},
			PLMNIdentity: ngapType.PLMNIdentity{
				Value: aper.OctetString("\x02\xF8\x39"),
			},
		},
		nasPdu: []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
	},
}

func TestBuildUplinkNasTransport(t *testing.T) {
	for _, testCase := range testBuildUplinkNasTransportCases {
		t.Run(testCase.name, func(t *testing.T) {
			pdu := buildUplinkNasTransport(testCase.amfUeNgapId, testCase.ranUeNgapId, testCase.plmnId, testCase.tai, testCase.nasPdu)
			encodeData, err := ngap.Encoder(pdu)
			if err != nil {
				t.Fatalf("Failed to encode NGAP uplink nas transport: %v", err)
			} else {
				decodeData, err := ngap.Decoder(encodeData)
				if err != nil {
					t.Fatalf("Failed to decode NGAP uplink nas transport: %v", err)
				} else if !reflect.DeepEqual(pdu, *decodeData) {
					t.Fatalf("NGAP uplink nas transport mismatch")
				}
			}
		})
	}
}

var testBuildNgapInitialContextSetupResponseCases = []struct {
	name        string
	amfUeNgapId int64
	ranUeNgapId int64
}{
	{
		name:        "testBuildNgapInitialContextSetupResponse",
		amfUeNgapId: 1,
		ranUeNgapId: 1,
	},
}

func TestBuildNgapInitialContextSetupResponse(t *testing.T) {
	for _, testCase := range testBuildNgapInitialContextSetupResponseCases {
		t.Run(testCase.name, func(t *testing.T) {
			pdu := buildNgapInitialContextSetupResponse(testCase.amfUeNgapId, testCase.ranUeNgapId)
			encodeData, err := ngap.Encoder(pdu)
			if err != nil {
				t.Fatalf("Failed to encode NGAP initial context setup response: %v", err)
			} else {
				decodeData, err := ngap.Decoder(encodeData)
				if err != nil {
					t.Fatalf("Failed to decode NGAP initial context setup response: %v", err)
				} else if !reflect.DeepEqual(pdu, *decodeData) {
					t.Fatalf("NGAP initial context setup response mismatch")
				}
			}
		})
	}
}

var testBuildPduSessionResourceSetupResponseTransferMessageCases = []struct {
	name    string
	dlTeid  []byte
	ranN3Ip string
	qosId   int64
}{
	{
		name:    "testBuildPduSessionResourceSetupResponseTransferMessage",
		dlTeid:  []byte("\x00\x00\x00\x01"),
		ranN3Ip: "127.0.0.1",
		qosId:   1,
	},
}

func TestBuildPduSessionResourceSetupResponseTransferMessage(t *testing.T) {
	for _, testCase := range testBuildPduSessionResourceSetupResponseTransferMessageCases {
		t.Run(testCase.name, func(t *testing.T) {
			transferMessage := buildPduSessionResourceSetupResponseTransfer(testCase.dlTeid, testCase.ranN3Ip, testCase.qosId)
			encodeTransferMessage, err := aper.MarshalWithParams(transferMessage, "valueExt")
			if err != nil {
				t.Fatalf("Failed to marshal pdu session resource setup response transfer message: %v", err)
			} else {
				decodeTransferMessage := &ngapType.PDUSessionResourceSetupResponseTransfer{}
				if err := aper.UnmarshalWithParams(encodeTransferMessage, decodeTransferMessage, "valueExt"); err != nil {
					t.Fatalf("Failed to unmarshal pdu session resource setup response transfer message: %v", err)
				} else if !reflect.DeepEqual(transferMessage, *decodeTransferMessage) {
					t.Fatalf("PDU session resource setup response transfer message mismatch")
				}
			}
		})
	}
}

var testBuildPduSessionResourceSetupResponseCases = []struct {
	name            string
	amfUeNgapId     int64
	ranUeNgapId     int64
	pduSessionId    int64
	transferMessage []byte
}{
	{
		name:            "testBuildPduSessionResourceSetupResponse",
		amfUeNgapId:     1,
		ranUeNgapId:     1,
		pduSessionId:    1,
		transferMessage: []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
	},
}

func TestBuildPduSessionResourceSetupResponse(t *testing.T) {
	for _, testCase := range testBuildPduSessionResourceSetupResponseCases {
		t.Run(testCase.name, func(t *testing.T) {
			pdu := buildPduSessionResourceSetupResponse(testCase.amfUeNgapId, testCase.ranUeNgapId, testCase.pduSessionId, testCase.transferMessage)
			encodeData, err := ngap.Encoder(pdu)
			if err != nil {
				t.Fatalf("Failed to encode NGAP pdu session resource setup response: %v", err)
			} else {
				decodeData, err := ngap.Decoder(encodeData)
				if err != nil {
					t.Fatalf("Failed to decode NGAP pdu session resource setup response: %v", err)
				} else if !reflect.DeepEqual(pdu, *decodeData) {
					t.Fatalf("NGAP pdu session resource setup response mismatch")
				}
			}
		})
	}
}
