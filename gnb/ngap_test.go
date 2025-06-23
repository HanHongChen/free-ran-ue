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
