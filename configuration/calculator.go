package configuration

import (
	"fmt"
	"strconv"
)

func Recalculation(flagStr string, m map[string]string) map[string]string {
	var divisor, err = strconv.ParseFloat(m[flagStr], 64)
	if err == nil {
		fmt.Println(err)
	}

	subMap := make(map[string]string)

	for k, val := range m {
		dividend, err := strconv.ParseFloat(val, 64)
		if err == nil {
			fmt.Println(err)
		}

		subMap[k] = fmt.Sprintf("%.2f", dividend/divisor)
	}

	return subMap
}
