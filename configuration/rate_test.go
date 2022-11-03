package configuration_test

import (
	"encoding/json"
	"github.com/matthew/rateBot/configuration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyData_UnmarshallJSON(t *testing.T) {
	sourse := `{
    "Date": "2022-02-26T11:30:00+03:00",
    "PreviousDate": "2022-02-25T11:30:00+03:00",
    "PreviousURL": "\/\/www.cbr-xml-daily.ru\/archive\/2022\/02\/25\/daily_json.js",
    "Timestamp": "2022-02-26T23:00:00+03:00",
    "Valute": {
        "AUD": {
            "ID": "R01010",
            "NumCode": "036",
            "CharCode": "AUD",
            "Nominal": 1,
            "Name": "Австралийский доллар",
            "Value": 60.0964,
            "Previous": 62.4844
        },
        "AZN": {
            "ID": "R01020A",
            "NumCode": "944",
            "CharCode": "AZN",
            "Nominal": 1,
            "Name": "Азербайджанский манат",
            "Value": 49.1751,
            "Previous": 51.1647
        },
        "GBP": {
            "ID": "R01035",
            "NumCode": "826",
            "CharCode": "GBP",
            "Nominal": 1,
            "Name": "Фунт стерлингов Соединенного королевства",
            "Value": 111.8213,
            "Previous": 117.1539
        },
        "AMD": {
            "ID": "R01060",
            "NumCode": "051",
            "CharCode": "AMD",
            "Nominal": 100,
            "Name": "Армянских драмов",
            "Value": 17.4059,
            "Previous": 18.1102
        },
        "BYN": {
            "ID": "R01090B",
            "NumCode": "933",
            "CharCode": "BYN",
            "Nominal": 1,
            "Name": "Белорусский рубль",
            "Value": 30.3261,
            "Previous": 31.0349
        },
        "BGN": {
            "ID": "R01100",
            "NumCode": "975",
            "CharCode": "BGN",
            "Nominal": 1,
            "Name": "Болгарский лев",
            "Value": 47.7475,
            "Previous": 50.0137
        },
        "BRL": {
            "ID": "R01115",
            "NumCode": "986",
            "CharCode": "BRL",
            "Nominal": 1,
            "Name": "Бразильский реал",
            "Value": 16.3066,
            "Previous": 17.3504
        },
        "HUF": {
            "ID": "R01135",
            "NumCode": "348",
            "CharCode": "HUF",
            "Nominal": 100,
            "Name": "Венгерских форинтов",
            "Value": 25.3947,
            "Previous": 26.9505
        },
        "HKD": {
            "ID": "R01200",
            "NumCode": "344",
            "CharCode": "HKD",
            "Nominal": 1,
            "Name": "Гонконгский доллар",
            "Value": 10.7008,
            "Previous": 11.1339
        },
        "DKK": {
            "ID": "R01215",
            "NumCode": "208",
            "CharCode": "DKK",
            "Nominal": 1,
            "Name": "Датская крона",
            "Value": 12.5507,
            "Previous": 13.1467
        },
        "USD": {
            "ID": "R01235",
            "NumCode": "840",
            "CharCode": "USD",
            "Nominal": 1,
            "Name": "Доллар США",
            "Value": 83.5485,
            "Previous": 86.9288
        },
        "EUR": {
            "ID": "R01239",
            "NumCode": "978",
            "CharCode": "EUR",
            "Nominal": 1,
            "Name": "Евро",
            "Value": 93.5994,
            "Previous": 97.7688
        },
        "INR": {
            "ID": "R01270",
            "NumCode": "356",
            "CharCode": "INR",
            "Nominal": 10,
            "Name": "Индийских рупий",
            "Value": 11.0833,
            "Previous": 11.5478
        },
        "KZT": {
            "ID": "R01335",
            "NumCode": "398",
            "CharCode": "KZT",
            "Nominal": 100,
            "Name": "Казахстанских тенге",
            "Value": 18.0019,
            "Previous": 18.5306
        },
        "CAD": {
            "ID": "R01350",
            "NumCode": "124",
            "CharCode": "CAD",
            "Nominal": 1,
            "Name": "Канадский доллар",
            "Value": 65.2621,
            "Previous": 67.9344
        },
        "KGS": {
            "ID": "R01370",
            "NumCode": "417",
            "CharCode": "KGS",
            "Nominal": 100,
            "Name": "Киргизских сомов",
            "Value": 98.4926,
            "Previous": 10.2478
        },
        "CNY": {
            "ID": "R01375",
            "NumCode": "156",
            "CharCode": "CNY",
            "Nominal": 1,
            "Name": "Китайский юань",
            "Value": 13.2325,
            "Previous": 13.7485
        },
        "MDL": {
            "ID": "R01500",
            "NumCode": "498",
            "CharCode": "MDL",
            "Nominal": 10,
            "Name": "Молдавских леев",
            "Value": 45.6549,
            "Previous": 48.16
        },
        "NOK": {
            "ID": "R01535",
            "NumCode": "578",
            "CharCode": "NOK",
            "Nominal": 10,
            "Name": "Норвежских крон",
            "Value": 92.9628,
            "Previous": 96.8976
        },
        "PLN": {
            "ID": "R01565",
            "NumCode": "985",
            "CharCode": "PLN",
            "Nominal": 1,
            "Name": "Польский злотый",
            "Value": 20.0505,
            "Previous": 21.1434
        },
        "RON": {
            "ID": "R01585F",
            "NumCode": "946",
            "CharCode": "RON",
            "Nominal": 1,
            "Name": "Румынский лей",
            "Value": 18.8759,
            "Previous": 19.7637
        },
        "XDR": {
            "ID": "R01589",
            "NumCode": "960",
            "CharCode": "XDR",
            "Nominal": 1,
            "Name": "СДР (специальные права заимствования)",
            "Value": 116.4816,
            "Previous": 121.9107
        },
        "SGD": {
            "ID": "R01625",
            "NumCode": "702",
            "CharCode": "SGD",
            "Nominal": 1,
            "Name": "Сингапурский доллар",
            "Value": 61.7277,
            "Previous": 64.2917
        },
        "TJS": {
            "ID": "R01670",
            "NumCode": "972",
            "CharCode": "TJS",
            "Nominal": 10,
            "Name": "Таджикских сомони",
            "Value": 73.94,
            "Previous": 77.0304
        },
        "TRY": {
            "ID": "R01700J",
            "NumCode": "949",
            "CharCode": "TRY",
            "Nominal": 10,
            "Name": "Турецких лир",
            "Value": 59.403,
            "Previous": 61.553
        },
        "TMT": {
            "ID": "R01710A",
            "NumCode": "934",
            "CharCode": "TMT",
            "Nominal": 1,
            "Name": "Новый туркменский манат",
            "Value": 23.9052,
            "Previous": 24.8723
        },
        "UZS": {
            "ID": "R01717",
            "NumCode": "860",
            "CharCode": "UZS",
            "Nominal": 10000,
            "Name": "Узбекских сумов",
            "Value": 77.0723,
            "Previous": 80.1905
        },
        "UAH": {
            "ID": "R01720",
            "NumCode": "980",
            "CharCode": "UAH",
            "Nominal": 10,
            "Name": "Украинских гривен",
            "Value": 27.8032,
            "Previous": 29.075
        },
        "CZK": {
            "ID": "R01760",
            "NumCode": "203",
            "CharCode": "CZK",
            "Nominal": 10,
            "Name": "Чешских крон",
            "Value": 37.4741,
            "Previous": 39.2561
        },
        "SEK": {
            "ID": "R01770",
            "NumCode": "752",
            "CharCode": "SEK",
            "Nominal": 10,
            "Name": "Шведских крон",
            "Value": 87.6083,
            "Previous": 91.4588
        },
        "CHF": {
            "ID": "R01775",
            "NumCode": "756",
            "CharCode": "CHF",
            "Nominal": 1,
            "Name": "Швейцарский франк",
            "Value": 90.3325,
            "Previous": 94.4673
        },
        "ZAR": {
            "ID": "R01810",
            "NumCode": "710",
            "CharCode": "ZAR",
            "Nominal": 10,
            "Name": "Южноафриканских рэндов",
            "Value": 54.434,
            "Previous": 56.9685
        },
        "KRW": {
            "ID": "R01815",
            "NumCode": "410",
            "CharCode": "KRW",
            "Nominal": 1000,
            "Name": "Вон Республики Корея",
            "Value": 69.368,
            "Previous": 72.2627
        },
        "JPY": {
            "ID": "R01820",
            "NumCode": "392",
            "CharCode": "JPY",
            "Nominal": 100,
            "Name": "Японских иен",
            "Value": 72.4776,
            "Previous": 75.7582
        }
    }
}`
	var result configuration.CurrencyData
	err := json.Unmarshal([]byte(sourse), &result)
	assert.NoError(t, err)
	assert.Equal(t, "R01235", result.Currency.Usd.ID)
}
