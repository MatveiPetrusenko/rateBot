// Package cbr_api /*
package cbr_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CurrencyData represents ...
type CurrencyData struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Currency     struct {
		Usd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"USD"`
		Eur struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"EUR"`
	} `json:"Valute"`
}

//ParsingJSON ...
func ParsingJSON() string {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var mainData CurrencyData

	if err := json.Unmarshal(data, &mainData); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	//prepare new message
	var usd = strconv.FormatFloat(mainData.Currency.Usd.Value, 'f', 2, 64)
	var eur = strconv.FormatFloat(mainData.Currency.Eur.Value, 'f', 2, 64)

	dateReport := mainData.Date

	moneyStr := "Курсы валют на: " + dateReport.Format(time.RFC1123) + "\n" +
		"\xF0\x9F\x92\xB5" + "US Dollar" + "  " + usd + "$" + "\n" +
		"\xF0\x9F\x92\xB6" + "Euro" + "  " + eur + "€"

	return moneyStr
}
