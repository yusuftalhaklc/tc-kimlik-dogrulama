package model

import "encoding/xml"

type TCKimlikNoDogrula struct {
	XMLName  xml.Name `xml:"soap12:Body"`
	TCKimlik TCKimlik `xml:"http://tckimlik.nvi.gov.tr/WS TCKimlikNoDogrula"`
}

type TCKimlikNoDogrulaResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		Result  struct {
			XMLName xml.Name `xml:"TCKimlikNoDogrulaResponse"`
			Value   bool     `xml:"TCKimlikNoDogrulaResult"`
		} `xml:"TCKimlikNoDogrulaResponse"`
	} `xml:"Body"`
}
