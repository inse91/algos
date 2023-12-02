package leetcode

import "strings"

func intToRoman(num int) string {

	romanMap := []struct {
		n      int
		symbol string
	}{
		{n: 1000, symbol: "M"},
		{n: 900, symbol: "CM"},
		{n: 500, symbol: "D"},
		{n: 400, symbol: "CD"},
		{n: 100, symbol: "C"},
		{n: 90, symbol: "XC"},
		{n: 50, symbol: "L"},
		{n: 40, symbol: "XL"},
		{n: 10, symbol: "X"},
		{n: 9, symbol: "IX"},
		{n: 5, symbol: "V"},
		{n: 4, symbol: "IV"},
		{n: 1, symbol: "I"},
	}

	sb := new(strings.Builder)

	for _, rm := range romanMap {
		for num >= rm.n {
			sb.WriteString(rm.symbol)
			num -= rm.n
		}
	}
	return sb.String()
}
