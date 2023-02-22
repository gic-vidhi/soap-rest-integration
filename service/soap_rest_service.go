package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CallGetSoap(country string) (string, error) {
	url := "https://health-api.com/api/v1/covid-19/" + country
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Cookie", "sid=58e28418-b277-11ed-b751-58acd784c574")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	return string(body), nil
}

func CallPostSoap(country string) (string, error) {
	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"
	method := "POST"
	countryStr := "<sCountryISOCode>" + country + "</sCountryISOCode>"
	payloadString := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
		  <sCountryISOCode>US</sCountryISOCode>
		</CapitalCity>
	  </soap:Body>
	</soap:Envelope>`

	payloadString = strings.ReplaceAll(payloadString, "<sCountryISOCode>country</sCountryISOCode>", countryStr)
	fmt.Println(string(payloadString))
	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))

	return string(body), nil
}
