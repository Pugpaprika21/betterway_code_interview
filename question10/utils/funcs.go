package utils

import (
	"strconv"
	"strings"
)

func MoneyFormat(s float64, useDecimal bool) string {
	formatted := strconv.FormatFloat(s, 'f', 2, 64)
	parts := strings.Split(formatted, ".")
	integral := parts[0]

	decimal := ""
	if useDecimal {
		decimal = "." + parts[1]
	}

	integralWithCommas := ""
	for i, rune := range integral {
		if i != 0 && (len(integral)-i)%3 == 0 {
			integralWithCommas += ","
		}
		integralWithCommas += string(rune)
	}

	return integralWithCommas + decimal
}
