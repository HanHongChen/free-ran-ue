package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func SendHttpRequest(uri string, method string, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if string(body) == "" {
		return nil, fmt.Errorf("received empty response from server")
	}

	return body, nil
}
