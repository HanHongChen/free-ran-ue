package util_test

import (
	"fmt"
	"testing"

	"github.com/Alonza0314/free-ran-ue/model"
	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/stretchr/testify/assert"
)

var testValidateLoggerIeCases = []struct {
	name          string
	loggerIe      model.LoggerIE
	expectedError error
}{
	{
		name:          "testError",
		loggerIe:      model.LoggerIE{Level: "error"},
		expectedError: nil,
	},
	{
		name:          "testWarn",
		loggerIe:      model.LoggerIE{Level: "warn"},
		expectedError: nil,
	},
	{
		name:          "testInfo",
		loggerIe:      model.LoggerIE{Level: "info"},
		expectedError: nil,
	},
	{
		name:          "testDebug",
		loggerIe:      model.LoggerIE{Level: "debug"},
		expectedError: nil,
	},
	{
		name:          "testTrace",
		loggerIe:      model.LoggerIE{Level: "trace"},
		expectedError: nil,
	},
	{
		name:          "testTest",
		loggerIe:      model.LoggerIE{Level: "test"},
		expectedError: nil,
	},
	{
		name:          "testInvalid",
		loggerIe:      model.LoggerIE{Level: "invalid"},
		expectedError: fmt.Errorf("invalid logger level: invalid"),
	},
}

func TestValidateLoggerIe(t *testing.T) {
	for _, testCase := range testValidateLoggerIeCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidateLoggerIe(&testCase.loggerIe)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidateIpCases = []struct {
	name          string
	ip            string
	expectedError error
}{
	{
		name:          "testValidIp",
		ip:            "192.168.1.1",
		expectedError: nil,
	},
	{
		name:          "testInvalidRangeIp",
		ip:            "192.168.1.256",
		expectedError: fmt.Errorf("invalid ip address: 192.168.1.256"),
	},
	{
		name:          "testInvalidFormatIp",
		ip:            "192.168.1.1.1",
		expectedError: fmt.Errorf("invalid ip address: 192.168.1.1.1"),
	},
}

func TestValidateIp(t *testing.T) {
	for _, testCase := range testValidateIpCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidateIp(testCase.ip)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidatePortCases = []struct {
	name          string
	port          int
	expectedError error
}{
	{
		name:          "testValidPort",
		port:          8080,
		expectedError: nil,
	},
	{
		name:          "testInvalidPort",
		port:          0,
		expectedError: fmt.Errorf("invalid port range: 0, range should be 1-65535"),
	},
}

func TestValidatePort(t *testing.T) {
	for _, testCase := range testValidatePortCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidatePort(testCase.port)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidateIntStringWithLengthCases = []struct {
	name          string
	intString     string
	length        int
	expectedError error
}{
	{
		name:          "testValidIntString",
		intString:     "12345",
		length:        5,
		expectedError: nil,
	},
	{
		name:          "testInvalidIntString",
		intString:     "12345a",
		length:        5,
		expectedError: fmt.Errorf("invalid int string: 12345a"),
	},
	{
		name:          "testInvalidIntStringLength",
		intString:     "12345",
		length:        10,
		expectedError: fmt.Errorf("invalid int string: 12345, length should be 10"),
	},
}

func TestValidateIntString(t *testing.T) {
	for _, testCase := range testValidateIntStringWithLengthCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidateIntStringWithLength(testCase.intString, testCase.length)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidatePlmnIdCases = []struct {
	name          string
	plmnId        model.PlmnIdIE
	expectedError error
}{
	{
		name:          "testValidPlmnId",
		plmnId:        model.PlmnIdIE{Mcc: "208", Mnc: "93"},
		expectedError: nil,
	},
	{
		name:          "testInvalidPlmnId",
		plmnId:        model.PlmnIdIE{Mcc: "208", Mnc: "930"},
		expectedError: fmt.Errorf("invalid mnc: 930, mnc should be 2 digits"),
	},
	{
		name:          "testInvalidNonIntMcc",
		plmnId:        model.PlmnIdIE{Mcc: "20a", Mnc: "93"},
		expectedError: fmt.Errorf("invalid int string: 20a"),
	},
	{
		name:          "testInvalidNonIntMnc",
		plmnId:        model.PlmnIdIE{Mcc: "208", Mnc: "9a"},
		expectedError: fmt.Errorf("invalid int string: 9a"),
	},
}

func TestValidatePlmnId(t *testing.T) {
	for _, testCase := range testValidatePlmnIdCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidatePlmnId(&testCase.plmnId)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidateMsinCases = []struct {
	name          string
	msin          string
	expectedError error
}{
	{
		name:          "testValidMsin",
		msin:          "0000000001",
		expectedError: nil,
	},
	{
		name:          "testInvalidMsin",
		msin:          "00000000010",
		expectedError: fmt.Errorf("invalid msin: 00000000010, msin should be 10 digits"),
	},
	{
		name:          "testInvalidNonIntMsin",
		msin:          "000000000a",
		expectedError: fmt.Errorf("invalid int string: 000000000a"),
	},
}

func TestValidateMsin(t *testing.T) {
	for _, testCase := range testValidateMsinCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidateMsin(testCase.msin)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

var testValidateAuthenticationSubscriptionCases = []struct {
	name                       string
	authenticationSubscription model.AuthenticationSubscriptionIE
	expectedError              error
}{
	{
		name: "testValidAuthenticationSubscription",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "8000",
			SequenceNumber:                "000000000023",
		},
		expectedError: nil,
	},
	{
		name: "testInvalidNonHexEncPermanentKey",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "zzzzzzzzzzzz",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "8000",
			SequenceNumber:                "000000000023",
		},
		expectedError: fmt.Errorf("invalid enc permanent key, invalid hex string: zzzzzzzzzzzz"),
	},
	{
		name: "testInvalidNonHexEncOpcKey",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "zzzzzzzzzzzz",
			AuthenticationManagementField: "8000",
			SequenceNumber:                "000000000023",
		},
		expectedError: fmt.Errorf("invalid enc opc key, invalid hex string: zzzzzzzzzzzz"),
	},
	{
		name: "testInvalidNonIntAuthenticationManagementField",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "800a",
			SequenceNumber:                "000000000023",
		},
		expectedError: fmt.Errorf("invalid authentication management field, invalid int string: 800a"),
	},
	{
		name: "testInvalidIntLengthAuthenticationManagementField",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "80000",
			SequenceNumber:                "000000000023",
		},
		expectedError: fmt.Errorf("invalid authentication management field, invalid int string: 80000, length should be 4"),
	},
	{
		name: "testInvalidNonIntSequenceNumber",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "8000",
			SequenceNumber:                "00000000002a",
		},
		expectedError: fmt.Errorf("invalid sequence number, invalid int string: 00000000002a"),
	},
	{
		name: "testInvalidIntLengthSequenceNumber",
		authenticationSubscription: model.AuthenticationSubscriptionIE{
			EncPermanentKey:               "8baf473f2f8fd09487cccbd7097c6862",
			EncOpcKey:                     "8e27b6af0e692e750f32667a3b14605d",
			AuthenticationManagementField: "8000",
			SequenceNumber:                "0000000000230",
		},
		expectedError: fmt.Errorf("invalid sequence number, invalid int string: 0000000000230, length should be 12"),
	},
}

func TestValidateAuthenticationSubscription(t *testing.T) {
	for _, testCase := range testValidateAuthenticationSubscriptionCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := util.ValidateAuthenticationSubscription(&testCase.authenticationSubscription)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
