// Package configuration  /*
package configuration

var mainValueCurrency = make(map[string]string)

var nameCurrency = []string{"US Dollar", "Euro", "British Pound", "Swiss Franc", "Russian Ruble", "Ukrainian Hryvnia", "Belarusian Ruble", "Kazakhstani Tenge"}
var symbolCurrency = []string{"$", "€", "£", "₣", "₽", "₴", "Rbl", "₸"}
var flagsCurrency = []string{"🇺🇸", "🇪🇺", "🇬🇧", "🇨🇭", "🇷🇺", "🇺🇦", "🇧🇾", "🇰🇿"}

var date string

// InitializeMassage initial each element with country flag
func InitializeMassage(arrCurrency [8]string, dateReport string) string {
	mainValueCurrency["🇺🇸"] = arrCurrency[0]
	mainValueCurrency["🇪🇺"] = arrCurrency[1]
	mainValueCurrency["🇬🇧"] = arrCurrency[2]
	mainValueCurrency["🇨🇭"] = arrCurrency[3]

	mainValueCurrency["🇷🇺"] = arrCurrency[4]
	mainValueCurrency["🇺🇦"] = arrCurrency[5]
	mainValueCurrency["🇧🇾"] = arrCurrency[6]
	mainValueCurrency["🇰🇿"] = arrCurrency[7]

	date = dateReport
	date += "\n"

	message := MakeMessage("🇺🇸")

	return message
}

// MakeMessage build message
func MakeMessage(flagStr string) string {
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

// SubMessage keep previous response and represent new message
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

/*В 4.5 раза быстрее среза, а расходует на 30% меньше памяти. И всего 6 аллокаций памяти!
Когда собираете строку из большого числа кусочков — используйте strings.Builder.*/
//text/template
