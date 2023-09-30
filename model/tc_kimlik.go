package model

type TCKimlik struct {
	TCKimlikNo int64  `xml:"TCKimlikNo"`
	Ad         string `xml:"Ad"`
	Soyad      string `xml:"Soyad"`
	DogumYili  int    `xml:"DogumYili"`
}
