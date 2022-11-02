package configuration

import (
	"time"
)

var mainValueCurrency = make(map[string]string)

var nameCurrency = []string{"US Dollar", "Euro", "British Pound", "Swiss Franc", "Russian Ruble", "Ukrainian Hryvnia", "Belarusian Ruble", "Kazakhstani Tenge"}
var symbolCurrency = []string{"$", "€", "£", "₣", "₽", "₴", "Rbl", "₸"}
var flagsCurrency = []string{"🇺🇸", "🇪🇺", "🇬🇧", "🇨🇭", "🇷🇺", "🇺🇦", "🇧🇾", "🇰🇿"}

var date string

func initializeMassage(arrCurrency [8]string, dateReport time.Time) string {
	mainValueCurrency["🇺🇸"] = arrCurrency[0]
	mainValueCurrency["🇪🇺"] = arrCurrency[1]
	mainValueCurrency["🇬🇧"] = arrCurrency[2]
	mainValueCurrency["🇨🇭"] = arrCurrency[3]

	mainValueCurrency["🇷🇺"] = arrCurrency[4]
	mainValueCurrency["🇺🇦"] = arrCurrency[5]
	mainValueCurrency["🇧🇾"] = arrCurrency[6]
	mainValueCurrency["🇰🇿"] = arrCurrency[7]

	date = dateReport.Format(time.RFC850)
	date += "\n"

	message := makeMessage("🇺🇸")

	return message
}

func makeMessage(flagStr string) string {
	var message, headStr string

	for k, val := range flagsCurrency {
		if val == flagStr {
			headStr += val + nameCurrency[k] + " " + mainValueCurrency[val] + symbolCurrency[k] + "\n" + "\n"
			break
		}

	}

	for k, val := range flagsCurrency {
		if val == flagStr {
			continue
		}
		message += val + nameCurrency[k] + " " + mainValueCurrency[val] + symbolCurrency[k] + "\n"
	}

	return date + headStr + message
}

func SubMessage(flagStr string) string {
	subValueCurrency := Recalculation(flagStr, mainValueCurrency)

	var replyMessage, headStr string

	for k, val := range flagsCurrency {
		if val == flagStr {
			headStr += val + nameCurrency[k] + " " + subValueCurrency[val] + symbolCurrency[k] + "\n" + "\n"
			break
		}

	}

	for k, val := range flagsCurrency {
		if val == flagStr {
			continue
		}
		replyMessage += val + nameCurrency[k] + " " + subValueCurrency[val] + symbolCurrency[k] + "\n"
	}

	return date + headStr + replyMessage
}
