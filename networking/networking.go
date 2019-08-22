package networking

import (
	"bytes"
	"net/http"
)

// SendSoap send soap message
func SendSoap(endpoint string, message []byte) (*http.Response, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewReader(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}
