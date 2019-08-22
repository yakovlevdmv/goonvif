package soap

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*************************
	WS-Security types
*************************/
const (
	passwordType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
	encodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
)

//XMLName xml.Name `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
type security struct {
	//XMLName xml.Name  `xml:"wsse:Security"`
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
	Auth    wsAuth
}

type password struct {
	//XMLName xml.Name `xml:"wsse:Password"`
	Type     string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type nonce struct {
	//XMLName xml.Name `xml:"wsse:Nonce"`
	Type  string `xml:"EncodingType,attr"`
	Nonce string `xml:",chardata"`
}

type wsAuth struct {
	XMLName  xml.Name `xml:"UsernameToken"`
	Username string   `xml:"Username"`
	Password password `xml:"Password"`
	Nonce    nonce    `xml:"Nonce"`
	Created  string   `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Created"`
}

/*
   <Security s:mustUnderstand="1" xmlns="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
       <UsernameToken>
           <Username>admin</Username>
           <Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest">edBuG+qVavQKLoWuGWQdPab4IBE=</Password>
           <Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">S7wO1ZFTh0KXv2CR7bd2ZXkLAAAAAA==</Nonce>
           <Created xmlns="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">2018-04-10T18:04:25.836Z</Created>
       </UsernameToken>
   </Security>
*/

func NewSecurity(username, passwd string) security {
	/** Generating Nonce sequence **/
	//charsToGenerate := 32
	//charSet := gostrgen.Lower | gostrgen.Digit

	crtTime := time.Now().UTC()
	crtTimeStr := crtTime.Format(time.RFC3339Nano)

	nonceStr, nonceBase64 := genNonce()

	auth := security{
		Auth: wsAuth{
			Username: username,
			Password: password{
				Type:     passwordType,
				Password: generateToken(nonceStr, crtTimeStr, passwd),
			},
			Nonce: nonce{
				Type:  encodingType,
				Nonce: nonceBase64,
			},
			Created: crtTimeStr,
		},
	}

	return auth
}

func genNonce() (string, string) {
	nonceObj, _ := uuid.NewV4()
	nonce := hex.EncodeToString(nonceObj.Bytes())
	nonceBase64 := base64.StdEncoding.EncodeToString([]byte(nonce))
	return nonce, nonceBase64
}

//Digest = B64ENCODE( SHA1( B64DECODE( Nonce ) + Date + Password ) )
func generateToken(Nonce, Created, Password string) string {

	hasher := sha1.New()
	//hasher.Write([]byte((base64.StdEncoding.EncodeToString([]byte(Nonce)) + Created.Format(time.RFC3339) + Password)))
	hasher.Write([]byte(Nonce + Created + Password))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
