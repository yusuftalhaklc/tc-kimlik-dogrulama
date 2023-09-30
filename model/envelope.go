package model

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	XSI     string   `xml:"xmlns:xsi,attr"`
	XSD     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    TCKimlikNoDogrula
}
