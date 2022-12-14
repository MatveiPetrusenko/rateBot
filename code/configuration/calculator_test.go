package configuration_test

import (
	"fmt"
	"github.com/matthew/rateBot/configuration"
	"testing"
)

func TestRecalculation(t *testing.T) {
	var testData = []struct {
		flagStr   string
		wantFlag  string
		wantValue string
	}{
		{"🇺🇸", "🇺🇸", "1.00"},
		{"🇪🇺", "🇺🇸", "0.97"},
		{"🇬🇧", "🇺🇸", "1.12"},
		{"🇨🇭", "🇺🇸", "0.99"},

		{"🇷🇺", "🇺🇸", "0.02"},
		{"🇺🇦", "🇺🇸", "0.03"},
		{"🇧🇾", "🇺🇸", "0.39"},
		{"🇰🇿", "🇺🇸", "0.00"},
	}

	mainValueCurrency := map[string]string{
		"🇺🇸": "1.00",
		"🇪🇺": "1.03",
		"🇬🇧": "0.89",
		"🇨🇭": "1.01",

		"🇷🇺": "62.03",
		"🇺🇦": "36.90",
		"🇧🇾": "2.54",
		"🇰🇿": "464.23",
	}

	for _, test := range testData {
		name := fmt.Sprintf("key%q:rate(%s)", test.flagStr, mainValueCurrency[test.flagStr])

		t.Run(name, func(t *testing.T) {
			got := configuration.Recalculation(test.flagStr, mainValueCurrency)
			if got[test.wantFlag] != test.wantValue {
				t.Errorf("got %q%s; want %q%s", test.wantFlag, got[test.wantFlag], test.wantFlag, test.wantValue)
			}
		})
	}
}
