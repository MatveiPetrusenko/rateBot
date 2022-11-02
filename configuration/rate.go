// Package configuration /*
package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// CurrencyData represents ...
type CurrencyData struct {
	Table      string    `json:"table"`
	Rates      Rates     `json:"rates"`
	Lastupdate time.Time `json:"lastupdate"`
}
type Rates struct {
	Aed       float64 `json:"AED"`
	Afn       float64 `json:"AFN"`
	All       float64 `json:"ALL"`
	Amd       float64 `json:"AMD"`
	Ang       float64 `json:"ANG"`
	Aoa       float64 `json:"AOA"`
	Ars       float64 `json:"ARS"`
	Aud       float64 `json:"AUD"`
	Awg       float64 `json:"AWG"`
	Azn       float64 `json:"AZN"`
	Bam       float64 `json:"BAM"`
	Bbd       int     `json:"BBD"`
	Bdt       float64 `json:"BDT"`
	Bgn       float64 `json:"BGN"`
	Bhd       float64 `json:"BHD"`
	Bif       float64 `json:"BIF"`
	Bmd       int     `json:"BMD"`
	Bnd       float64 `json:"BND"`
	Bob       float64 `json:"BOB"`
	Brl       float64 `json:"BRL"`
	Bsd       int     `json:"BSD"`
	Btc       float64 `json:"BTC"`
	Btn       float64 `json:"BTN"`
	Bts       float64 `json:"BTS"`
	Bwp       float64 `json:"BWP"`
	Byn       float64 `json:"BYN"`
	Bzd       float64 `json:"BZD"`
	Cad       float64 `json:"CAD"`
	Cdf       float64 `json:"CDF"`
	Chf       float64 `json:"CHF"`
	Clf       float64 `json:"CLF"`
	Clp       float64 `json:"CLP"`
	Cnh       float64 `json:"CNH"`
	Cny       float64 `json:"CNY"`
	Cop       float64 `json:"COP"`
	Crc       float64 `json:"CRC"`
	Cuc       int     `json:"CUC"`
	Cup       float64 `json:"CUP"`
	Cve       float64 `json:"CVE"`
	Czk       float64 `json:"CZK"`
	Dash      float64 `json:"DASH"`
	Djf       float64 `json:"DJF"`
	Dkk       float64 `json:"DKK"`
	Doge      float64 `json:"DOGE"`
	Dop       float64 `json:"DOP"`
	Dzd       float64 `json:"DZD"`
	Egp       float64 `json:"EGP"`
	Ern       int     `json:"ERN"`
	Etb       float64 `json:"ETB"`
	Eth       float64 `json:"ETH"`
	Eur       float64 `json:"EUR"`
	Fjd       float64 `json:"FJD"`
	Fkp       float64 `json:"FKP"`
	Gbp       float64 `json:"GBP"`
	Gel       float64 `json:"GEL"`
	Ggp       float64 `json:"GGP"`
	Ghs       float64 `json:"GHS"`
	Gip       float64 `json:"GIP"`
	Gmd       int     `json:"GMD"`
	Gnf       float64 `json:"GNF"`
	Gtq       float64 `json:"GTQ"`
	Gyd       float64 `json:"GYD"`
	Hkd       float64 `json:"HKD"`
	Hnl       float64 `json:"HNL"`
	Hrk       float64 `json:"HRK"`
	Htg       float64 `json:"HTG"`
	Huf       float64 `json:"HUF"`
	Idr       float64 `json:"IDR"`
	Ils       float64 `json:"ILS"`
	Imp       float64 `json:"IMP"`
	Inr       float64 `json:"INR"`
	Iqd       float64 `json:"IQD"`
	Irr       int     `json:"IRR"`
	Isk       float64 `json:"ISK"`
	Jep       float64 `json:"JEP"`
	Jmd       float64 `json:"JMD"`
	Jod       float64 `json:"JOD"`
	Jpy       float64 `json:"JPY"`
	Kes       float64 `json:"KES"`
	Kgs       float64 `json:"KGS"`
	Khr       float64 `json:"KHR"`
	Kmf       float64 `json:"KMF"`
	Kpw       int     `json:"KPW"`
	Krw       float64 `json:"KRW"`
	Kwd       float64 `json:"KWD"`
	Kyd       float64 `json:"KYD"`
	Kzt       float64 `json:"KZT"`
	Lak       float64 `json:"LAK"`
	Lbp       float64 `json:"LBP"`
	Ld        int     `json:"LD"`
	Lkr       float64 `json:"LKR"`
	Lrd       float64 `json:"LRD"`
	Lsl       float64 `json:"LSL"`
	Ltc       float64 `json:"LTC"`
	Lyd       float64 `json:"LYD"`
	Mad       float64 `json:"MAD"`
	Mdl       float64 `json:"MDL"`
	Mga       float64 `json:"MGA"`
	Mkd       float64 `json:"MKD"`
	Mmk       float64 `json:"MMK"`
	Mnt       float64 `json:"MNT"`
	Mop       float64 `json:"MOP"`
	Mro       float64 `json:"MRO"`
	Mru       float64 `json:"MRU"`
	Mur       float64 `json:"MUR"`
	Mvr       float64 `json:"MVR"`
	Mwk       float64 `json:"MWK"`
	Mxn       float64 `json:"MXN"`
	Myr       float64 `json:"MYR"`
	Mzn       float64 `json:"MZN"`
	Nad       float64 `json:"NAD"`
	Ngn       float64 `json:"NGN"`
	Nio       float64 `json:"NIO"`
	Nok       float64 `json:"NOK"`
	Npr       float64 `json:"NPR"`
	Nxt       float64 `json:"NXT"`
	Nzd       float64 `json:"NZD"`
	Omr       float64 `json:"OMR"`
	Pab       int     `json:"PAB"`
	Pen       float64 `json:"PEN"`
	Pgk       float64 `json:"PGK"`
	Php       float64 `json:"PHP"`
	Pkr       float64 `json:"PKR"`
	Pln       float64 `json:"PLN"`
	Pyg       float64 `json:"PYG"`
	Qar       float64 `json:"QAR"`
	Ron       float64 `json:"RON"`
	Rsd       float64 `json:"RSD"`
	Rub       float64 `json:"RUB"`
	Rwf       float64 `json:"RWF"`
	Sar       float64 `json:"SAR"`
	Sbd       float64 `json:"SBD"`
	Scr       float64 `json:"SCR"`
	Sdg       float64 `json:"SDG"`
	Sek       float64 `json:"SEK"`
	Sgd       float64 `json:"SGD"`
	Shp       float64 `json:"SHP"`
	Sll       int     `json:"SLL"`
	Sos       float64 `json:"SOS"`
	Srd       float64 `json:"SRD"`
	Ssp       float64 `json:"SSP"`
	Std       float64 `json:"STD"`
	Stn       float64 `json:"STN"`
	Str       float64 `json:"STR"`
	Svc       float64 `json:"SVC"`
	Syp       float64 `json:"SYP"`
	Szl       float64 `json:"SZL"`
	Thb       float64 `json:"THB"`
	Tjs       float64 `json:"TJS"`
	Tmt       float64 `json:"TMT"`
	Tnd       float64 `json:"TND"`
	Top       float64 `json:"TOP"`
	Try       float64 `json:"TRY"`
	Ttd       float64 `json:"TTD"`
	Twd       float64 `json:"TWD"`
	Tzs       float64 `json:"TZS"`
	Uah       float64 `json:"UAH"`
	Ugx       float64 `json:"UGX"`
	Usd       float64 `json:"USD"`
	Uyu       float64 `json:"UYU"`
	Uzs       float64 `json:"UZS"`
	VefBlkmkt float64 `json:"VEF_BLKMKT"`
	VefDicom  float64 `json:"VEF_DICOM"`
	VefDipro  float64 `json:"VEF_DIPRO"`
	Ves       float64 `json:"VES"`
	Vnd       float64 `json:"VND"`
	Vuv       float64 `json:"VUV"`
	Wst       float64 `json:"WST"`
	Xaf       float64 `json:"XAF"`
	Xag       float64 `json:"XAG"`
	Xau       float64 `json:"XAU"`
	Xcd       float64 `json:"XCD"`
	Xdr       float64 `json:"XDR"`
	Xmr       float64 `json:"XMR"`
	Xof       float64 `json:"XOF"`
	Xpd       float64 `json:"XPD"`
	Xpf       float64 `json:"XPF"`
	Xpt       float64 `json:"XPT"`
	Xrp       float64 `json:"XRP"`
	Yer       float64 `json:"YER"`
	Zar       float64 `json:"ZAR"`
	Zmw       float64 `json:"ZMW"`
	Zwl       int     `json:"ZWL"`
	Nmc       float64 `json:"NMC"`
	Ppc       float64 `json:"PPC"`
	Nvc       float64 `json:"NVC"`
	Xpm       float64 `json:"XPM"`
	Eac       float64 `json:"EAC"`
	Vtc       float64 `json:"VTC"`
	Emc       float64 `json:"EMC"`
	Fct       float64 `json:"FCT"`
}

// ParsingJSON ...
func ParsingJSON() string {
	url := "https://cdn.cur.su/api/latest.json"

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
	var usd = fmt.Sprintf("%.2f", mainData.Rates.Usd)
	var eur = fmt.Sprintf("%.2f", mainData.Rates.Eur)
	var gbp = fmt.Sprintf("%.2f", mainData.Rates.Gbp)
	var chf = fmt.Sprintf("%.2f", mainData.Rates.Chf)

	var rub = fmt.Sprintf("%.2f", mainData.Rates.Rub)
	var uah = fmt.Sprintf("%.2f", mainData.Rates.Uah)
	var byn = fmt.Sprintf("%.2f", mainData.Rates.Byn)
	var kzt = fmt.Sprintf("%.2f", mainData.Rates.Kzt)

	var arrCurrency = [8]string{usd, eur, gbp, chf, rub, uah, byn, kzt}

	dateReport := mainData.Lastupdate

	result := initializeMassage(arrCurrency, dateReport)

	return result
}
