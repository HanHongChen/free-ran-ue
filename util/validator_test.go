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
