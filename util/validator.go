package util

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Alonza0314/free-ran-ue/model"
	loggergoUtil "github.com/Alonza0314/logger-go/v2/util"
)

/*
As before using validator, we have loaded the config from yaml file.
In the loaded function, it is ensured that the Data Type is correct.
All the validator functions need to do is to ensure the data value is valid.
*/

func ValidateLoggerIe(loggerIe *model.LoggerIE) error {
	switch loggergoUtil.LogLevelString(loggerIe.Level) {
	case loggergoUtil.LEVEL_STRING_ERROR:
		return nil
	case loggergoUtil.LEVEL_STRING_WARN:
		return nil
	case loggergoUtil.LEVEL_STRING_INFO:
		return nil
	case loggergoUtil.LEVEL_STRING_DEBUG:
		return nil
	case loggergoUtil.LEVEL_STRING_TRACE:
		return nil
	case loggergoUtil.LEVEL_STRING_TEST:
		return nil
	default:
		return fmt.Errorf("invalid logger level: %s", loggerIe.Level)
	}
}

func ValidateIp(ip string) error {
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return fmt.Errorf("invalid ip address: %s", ip)
	}
	return nil
}

func ValidatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port range: %d, range should be 1-65535", port)
	}
	return nil
}

func ValidateIntStringWithLength(intString string, length int) error {
	if _, err := strconv.Atoi(intString); err != nil {
		return fmt.Errorf("invalid int string: %s", intString)
	}
	if len(intString) != length {
		return fmt.Errorf("invalid int string: %s, length should be %d", intString, length)
	}
	return nil
}

func ValidatePlmnId(plmnId *model.PlmnIdIE) error {
	if len(plmnId.Mcc) != 3 {
		return fmt.Errorf("invalid mcc: %s, mcc should be 3 digits", plmnId.Mcc)
	}
	if err := ValidateIntStringWithLength(plmnId.Mcc, 3); err != nil {
		return err
	}
	if len(plmnId.Mnc) != 2 {
		return fmt.Errorf("invalid mnc: %s, mnc should be 2 digits", plmnId.Mnc)
	}
	if err := ValidateIntStringWithLength(plmnId.Mnc, 2); err != nil {
		return err
	}
	return nil
}

func ValidateUeIe(ueIe *model.UeIE) error {
	return nil
}

func ValidateUe(ue *model.UeConfig) error {
	if err := ValidateUeIe(&ue.Ue); err != nil {
		return err
	}
	if err := ValidateLoggerIe(&ue.Logger); err != nil {
		return err
	}
	return nil
}
