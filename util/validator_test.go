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
