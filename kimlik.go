package tckimlik

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yusuftalhaklc/tc-kimlik-dogrulama/model"
	"github.com/yusuftalhaklc/tc-kimlik-dogrulama/util"
)

func Dogrula(tc *model.TCKimlik) bool {
	payload := model.Envelope{
		XSI:    util.XSI,
		XSD:    util.XSD,
		Soap12: util.Soap12,
		Body: model.TCKimlikNoDogrula{
			TCKimlik: *tc,
		},
	}
	xmlPayload, err := xml.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return false
	}
	return sendHTTPClient(xmlPayload)
}

func sendHTTPClient(xmlPayload []byte) bool {
	payload := bytes.NewReader(xmlPayload)
	client := &http.Client{}
	req, err := http.NewRequest(util.HTTP_METHOD, util.BASE_URL, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", util.CONTENT_TYPE)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	response := new(model.TCKimlikNoDogrulaResponse)
	if err := xml.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return response.Body.Result.Value
}
