package util

import (
	"fmt"

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
