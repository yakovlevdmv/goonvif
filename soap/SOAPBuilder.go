package soap

import (
	"encoding/xml"
	"log"

	"github.com/beevik/etree"
)

// SoapMessage  SOAP Message
type SoapMessage string

// NewEmptySOAP new empty SOAP
func NewEmptySOAP() SoapMessage {
	doc := buildSoapRoot()
	//doc.IndentTabs()

	res, _ := doc.WriteToString()

	return SoapMessage(res)
}

// NewSOAP new SOAP
func NewSOAP(headContent []*etree.Element, bodyContent []*etree.Element, namespaces map[string]string) SoapMessage {
	doc := buildSoapRoot()
	//doc.IndentTabs()

	res, _ := doc.WriteToString()

	return SoapMessage(res)
}

// String to String
func (msg SoapMessage) String() string {
	return string(msg)
}

// Bytes []byte
func (msg SoapMessage) Bytes() []byte {
	return []byte(msg)
}

// StringIndent string indent
func (msg SoapMessage) StringIndent() string {
	doc := etree.NewDocument()

	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	doc.IndentTabs()
	res, _ := doc.WriteToString()

	return res
}

// Body body
func (msg SoapMessage) Body() string {

	doc := etree.NewDocument()

	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}
	bodyTag := doc.Root().SelectElement("Body").ChildElements()[0]
	doc.SetRoot(bodyTag)
	doc.IndentTabs()

	res, _ := doc.WriteToString()

	return res
}

// AddStringBodyContent AddStringBodyContent
func (msg *SoapMessage) AddStringBodyContent(data string) {
	doc := etree.NewDocument()

	if err := doc.ReadFromString(data); err != nil {
		log.Println(err.Error())
	}

	element := doc.Root()

	doc = etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}
	//doc.FindElement("./Envelope/Body").AddChild(element)
	bodyTag := doc.Root().SelectElement("Body")
	bodyTag.AddChild(element)

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddBodyContent AddBodyContent
func (msg *SoapMessage) AddBodyContent(element *etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}
	//doc.FindElement("./Envelope/Body").AddChild(element)
	bodyTag := doc.Root().SelectElement("Body")
	bodyTag.AddChild(element)

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddBodyContents AddBodyContents
func (msg *SoapMessage) AddBodyContents(elements []*etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	bodyTag := doc.Root().SelectElement("Body")

	if len(elements) != 0 {
		for _, j := range elements {
			bodyTag.AddChild(j)
		}
	}

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddStringHeaderContent AddStringHeaderContent
func (msg *SoapMessage) AddStringHeaderContent(data string) error {
	doc := etree.NewDocument()

	if err := doc.ReadFromString(data); err != nil {
		//log.Println(err.Error())
		return err
	}

	element := doc.Root()

	doc = etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		//log.Println(err.Error())
		return err
	}

	bodyTag := doc.Root().SelectElement("Header")
	bodyTag.AddChild(element)

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)

	return nil
}

// AddHeaderContent AddHeaderContent
func (msg *SoapMessage) AddHeaderContent(element *etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	bodyTag := doc.Root().SelectElement("Header")
	bodyTag.AddChild(element)

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddHeaderContents AddHeaderContents
func (msg *SoapMessage) AddHeaderContents(elements []*etree.Element) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	headerTag := doc.Root().SelectElement("Header")

	if len(elements) != 0 {
		for _, j := range elements {
			headerTag.AddChild(j)
		}
	}

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddRootNamespace AddRootNamespace
func (msg *SoapMessage) AddRootNamespace(key, value string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg.String()); err != nil {
		log.Println(err.Error())
	}

	doc.Root().CreateAttr("xmlns:"+key, value)

	//doc.IndentTabs()
	res, _ := doc.WriteToString()

	*msg = SoapMessage(res)
}

// AddRootNamespaces AddRootNamespaces
func (msg *SoapMessage) AddRootNamespaces(namespaces map[string]string) {
	for key, value := range namespaces {
		msg.AddRootNamespace(key, value)
	}

	/*
		doc := etree.NewDocument()
		if err := doc.ReadFromString(msg.String()); err != nil {
			//log.Println(err.Error())
			return err
		}

		for key, value := range namespaces {
			doc.Root().CreateAttr("xmlns:" + key, value)
		}

		doc.IndentTabs()
		res, _ := doc.WriteToString()

		*msg = SoapMessage(res)*/
}

func buildSoapRoot() *etree.Document {
	doc := etree.NewDocument()

	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	env := doc.CreateElement("Envelope")
	env.CreateElement("Header")
	env.CreateElement("Body")

	env.CreateAttr("xmlns", "http://www.w3.org/2003/05/soap-envelope")

	return doc
}

// AddWSSecurity AddWSSecurity
func (msg *SoapMessage) AddWSSecurity(username, password string) {
	//doc := etree.NewDocument()
	//if err := doc.ReadFromString(msg.String()); err != nil {
	//	log.Println(err.Error())
	//}

	/*
		Getting an WS-Security struct representation
	*/
	auth := NewSecurity(username, password)

	/*
		Adding WS-Security namespaces to root element of SOAP message
	*/
	//msg.AddRootNamespace("wsse", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext1.0.xsd")
	//msg.AddRootNamespace("wsu", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility1.0.xsd")

	soapReq, err := xml.MarshalIndent(auth, "", "  ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		panic(err)
	}

	/*
		Adding WS-Security struct to SOAP header
	*/
	msg.AddStringHeaderContent(string(soapReq))

	//doc.IndentTabs()
	//res, _ := doc.WriteToString()
	//
	//*msg = SoapMessage(res)
}
