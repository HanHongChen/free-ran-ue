package util

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func HexStringToBytes(hexStr string) ([]byte, error) {
	if len(hexStr)%2 != 0 {
		return nil, fmt.Errorf("hex string length must be even")
	}

	return hex.DecodeString(hexStr)
}

func HexStringToEscaped(hexStr string) (string, error) {
	if len(hexStr)%2 != 0 {
		return "", fmt.Errorf("hex string length must be even")
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	var escaped strings.Builder
	for _, b := range bytes {
		escaped.WriteString(fmt.Sprintf("\\x%02x", b))
	}

	return escaped.String(), nil
}
