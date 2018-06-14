package helper

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/beevik/etree"
)

//Key Field
const (
	SOAPENVEnveploe = "SOAP-ENV:Envelope"
	SOAPENVBody     = "SOAP-ENV:Body"
	SOAPENVFault    = "SOAP-ENV:Fault"
)

//GetBody from http resp
func GetBody(messageBody []byte, tdsName string) ([]byte, error) {

	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(messageBody); err != nil {
		return nil, err
	}

	body := doc.FindElement("./Envelope/Body")

	if body == nil {
		return messageBody, errors.New("bay response body")
	}

	fault := body.FindElement(SOAPENVFault)
	if fault != nil {
		doc.SetRoot(fault)
		buf, err := doc.WriteToBytes()
		if err != nil {
			return nil, err
		}
		//	log.Printf("%s\n", buf)
		return nil, errors.New(string(buf))
	}

	content := body.SelectElement(tdsName)
	if content == nil {
		return nil, errors.New(fmt.Sprint("element <", tdsName, "> not found in response soap body"))
	}

	doc.SetRoot(content)
	buf, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	// log.Printf("%s\n", buf)
	return buf, nil
}

func readMessageInbody(msgBody io.Reader) []byte {
	b, err := ioutil.ReadAll(msgBody)
	if err != nil {
		panic(err)
	}
	return b
}

//Unmarshal fuction from body
func Unmarshal(msgInBody io.Reader, tdsName string, v interface{}) error {

	messageContentBytes := readMessageInbody(msgInBody)

	buf, err := GetBody(messageContentBytes, tdsName)
	if err != nil {
		return err
	}
	//	log.Printf("%s\n", buf)
	return xml.Unmarshal(buf, v)
}
