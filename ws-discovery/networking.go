package wsdiscovery

/*******************************************************
 * Copyright (C) 2018 Palanjyan Zhorzhik
 *
 * This file is part of ws-discovery project.
 *
 * ws-discovery can be copied and/or distributed without the express
 * permission of Palanjyan Zhorzhik
 *******************************************************/
/*
<?xml version="1.0" encoding="UTF-8"?>
<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope"
    xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing">
    <Header>
        <a:Action mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</a:Action>
        <a:MessageID>uuid:78a2ed98-bc1f-4b08-9668-094fcba81e35</a:MessageID>
        <a:ReplyTo>
            <a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
        </a:ReplyTo>
        <a:To mustUnderstand="1">urn:schemas-xmlsoap-org:ws:2005:04:discovery</a:To>
    </Header>
    <Body>
        <Probe xmlns="http://schemas.xmlsoap.org/ws/2005/04/discovery">
            <d:Types xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery"
                xmlns:dp0="http://www.onvif.org/ver10/network/wsdl">dp0:NetworkVideoTransmitter</d:Types>
        </Probe>
    </Body>
</Envelope>
*/

import (
	"log"
	"net"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/net/ipv4"
)

var (
	bufSize      = 8192            //buffer bytes
	multicastTTL = 2               // 2 sec
	readDuration = time.Second * 1 //1 sec
)

//SetSearchContext Set context for search,set 0 to use default
func SetSearchContext(bufferSize int, multicastTimeToLive int, udpReadDuration time.Duration) {
	if bufferSize > 0 {
		bufSize = bufferSize
	}
	if multicastTimeToLive > 0 {
		multicastTTL = multicastTimeToLive
	}
	if readDuration > 0 {
		readDuration = udpReadDuration
	}
}

//SendProbe SendProbe
func SendProbe(interfaceName string, scopes, types []string, namespaces map[string]string) ([]string, error) {
	// Creating UUID Version 4
	uuidV4 := uuid.Must(uuid.NewV4())
	//fmt.Printf("probeSOAP: %s\n", probeSOAP)
	probeSOAP := buildProbeMessage(uuidV4.String(), scopes, types, namespaces)
	//fmt.Printf("probeSOAP: %s\n", probeSOAP)
	return sendUDPMulticast(probeSOAP.String(), interfaceName)

}

//SendProbeByInterface SendProbe by ip address
func SendProbeByInterface(uuidv4RequestString string, netInterface *net.Interface, scopes, types []string, namespaces map[string]string) ([]string, error) {

	if uuidv4RequestString == "" {
		uuidv4RequestString = uuid.Must(uuid.NewV4()).String()
	}
	if namespaces == nil {
		namespaces = map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"}
	}

	probeSOAP := buildProbeMessage(uuidv4RequestString, scopes, types, namespaces)
	//fmt.Printf("probeSOAP: %s\n", probeSOAP)
	return sendUDPMulticastByInterface(probeSOAP.String(), netInterface, "")

}

func sendUDPMulticast(msg string, interfaceName string) ([]string, error) {
	var result []string
	data := []byte(msg)
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return nil, err
	}
	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		return nil, err
	}
	defer c.Close()

	group := net.IPv4(239, 255, 255, 250)
	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		return nil, err
	}

	if err := p.SetMulticastInterface(iface); err != nil {
		log.Println("SetMulticastInterface: ", err)
	}
	p.SetMulticastTTL(multicastTTL)

	dst := &net.UDPAddr{IP: group, Port: 3702}
	if _, err := p.WriteTo(data, nil, dst); err != nil {
		log.Println("SetMulticastTTL: ", err)
	}
	if err := p.SetReadDeadline(time.Now().Add(readDuration)); err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, bufSize)
		n, _, _, err := p.ReadFrom(b)
		if err != nil {
			log.Println("ReadFrom UDP Buffer: ", err)
			break
		}
		result = append(result, string(b[0:n]))
	}
	return result, nil
}

func sendUDPMulticastByInterface(msg string, iface *net.Interface, port string) ([]string, error) {
	var result []string
	data := []byte(msg)

	if port == "" {
		port = "1024"
	}
	c, err := net.ListenPacket("udp4", "0.0.0.0:"+port)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	group := net.IPv4(239, 255, 255, 250)
	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		return nil, err
	}

	if err := p.SetMulticastInterface(iface); err != nil {
		log.Println("SetMulticastInterface: ", err)
	}
	if err := p.SetMulticastTTL(multicastTTL); err != nil {
		log.Println("SetMulticastTTL: ", err)
	}

	dst := &net.UDPAddr{IP: group, Port: 3702}
	if _, err := p.WriteTo(data, nil, dst); err != nil {
		log.Println("WriteTo : ", err)
	}
	if err := p.SetReadDeadline(time.Now().Add(readDuration)); err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, bufSize)
		n, _, _, err := p.ReadFrom(b)
		if err != nil {
			log.Println("ReadFrom UDP Buffer: ", err)
			break
		}
		result = append(result, string(b[0:n]))
	}
	return result, nil
}
