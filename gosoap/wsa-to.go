package gosoap

/*************************
	To Type in Header
*************************/

//To type
type To struct {
	XMLName struct{} `xml:"wsa5:To"`
	Value   string   `xml:",chardata"`
}

/*
   <wsa:To>http://160.10.64.10/Subscription?Idx=0</wsa:To>
*/

//NewTo get a new Head Action Section
func NewTo(toString string) To {

	return To{
		Value: toString,
		//	Created: time.Now().UTC().Format(time.RFC3339Nano),
	}

}
