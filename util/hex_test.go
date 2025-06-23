package util_test

import (
	"errors"
	"testing"

	"github.com/Alonza0314/free-ran-ue/util"
	"github.com/go-playground/assert/v2"
)

var testHexToBytesCases = []struct {
	name          string
	hexStr        string
	expectedBytes []byte
	expectedError error
}{
	{
		name:          "testHexToBytes",
		hexStr:        "000102",
		expectedBytes: []byte{0x00, 0x01, 0x02},
		expectedError: nil,
	},
	{
		name:          "testHexToBytesOddLength",
		hexStr:        "0003040",
		expectedBytes: nil,
		expectedError: errors.New("hex string length must be even"),
	},
}

var testHexToEscapedCases = []struct {
	name            string
	hexStr          string
	expectedEscaped string
	expectedError   error
}{
	{
		name:            "testHexToEscaped",
		hexStr:          "000102",
		expectedEscaped: "\\x00\\x01\\x02",
		expectedError:   nil,
	},
	{
		name:            "testHexToEscapedOddLength",
		hexStr:          "0003040",
		expectedEscaped: "",
		expectedError:   errors.New("hex string length must be even"),
	},
}

func TestHexToBytes(t *testing.T) {
	for _, testCase := range testHexToBytesCases {
		t.Run(testCase.name, func(t *testing.T) {
			bytes, err := util.HexStringToBytes(testCase.hexStr)
			assert.Equal(t, testCase.expectedBytes, bytes)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestHexToEscaped(t *testing.T) {
	for _, testCase := range testHexToEscapedCases {
		t.Run(testCase.name, func(t *testing.T) {
			escaped, err := util.HexStringToEscaped(testCase.hexStr)
			assert.Equal(t, testCase.expectedEscaped, escaped)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
