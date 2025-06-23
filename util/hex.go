package util

import (
	"encoding/hex"
	"fmt"
)

func HexStringToBytes(hexStr string) ([]byte, error) {
	if len(hexStr)%2 != 0 {
		return nil, fmt.Errorf("hex string length must be even")
	}

	return hex.DecodeString(hexStr)
}

func BytesToHexString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
