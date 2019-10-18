package networking

import (
	"bytes"
	"net/http"
	"time"
)

// SendSoap send soap message
func SendSoap(endpoint string, message []byte) (*http.Response, error) {
	return SendSoapWithTimeout(endpoint, message, time.Second*3)
}

// SendSoapWithTimeout send soap message with timeOut
func SendSoapWithTimeout(endpoint string, message []byte, timeout time.Duration) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: timeout,
	}

	return httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewReader(message))
}
