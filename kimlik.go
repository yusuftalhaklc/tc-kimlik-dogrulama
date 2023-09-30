package tckimlik

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/yusuftalhaklc/tc-kimlik-dogrulama/model"
	"github.com/yusuftalhaklc/tc-kimlik-dogrulama/util"
)

func Dogrula(tc *model.TCKimlik) (bool, error) {
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
		return false, err
	}
	return sendHTTPClient(xmlPayload)
}

func sendHTTPClient(xmlPayload []byte) (bool, error) {
	payload := bytes.NewReader(xmlPayload)
	client := &http.Client{}

	req, err := http.NewRequest(util.HTTP_METHOD, util.BASE_URL, payload)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", util.CONTENT_TYPE)

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	response := new(model.TCKimlikNoDogrulaResponse)
	if err := xml.Unmarshal([]byte(body), &response); err != nil {
		return false, err
	}

	return response.Body.Result.Value, nil
}
