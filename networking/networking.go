package networking

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"time"

	httpCli "github.com/gojektech/heimdall"
)

//SendSoap message
func SendSoap(endpoint, message string) (*http.Response, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

//HTTPPostEx use heimdall for circle break、 timeout、retry
func HTTPPostEx(endpoint, message string, timeout int) (*http.Response, error) {
	client := httpCli.NewHTTPClient(time.Duration(timeout) * time.Second)
	headers := http.Header{}

	postBody := []byte(message)
	headers.Add("Content-Type", "text/xml; charset=utf-8")
	headers.Add("Content-Length", strconv.Itoa(len(postBody)))

	resp, err := client.Post(endpoint, bytes.NewBuffer(postBody), headers)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

//SendSoapEx to post a request
func SendSoapEx(connect ConnectParameter, message string) ([]byte, error) {
	var url string
	if len(connect.Address) > 0 {
		url = connect.Address
	} else {
		url = "http://" + connect.Host + connect.URL
	}

	resp, err := HTTPPostEx(url, message, connect.Timeout)

	if err != nil {
		//	fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	var buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}
