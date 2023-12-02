package easy

import (
	"fmt"
	"strings"
)

func romanToInt(s string) (res int) {
	s = strings.Replace(s, "IV", "IIII", -1)
	s = strings.Replace(s, "IX", "VIIII", -1)
	s = strings.Replace(s, "XL", "XXXX", -1)
	s = strings.Replace(s, "XC", "LXXXX", -1)
	s = strings.Replace(s, "CD", "CCCC", -1)
	s = strings.Replace(s, "CM", "DCCCC", -1)
	for _, ch := range s {
		switch ch {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			res += 100
		case 'L':
			res += 50
		case 'X':
			res += 10
		case 'V':
			res += 5
		case 'I':
			res += 1
		default:
			return 0
		}
	}
	return
}

func romanToIntOld(s string) (res int) {
	for _, ch := range s {
		switch ch {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			res += 100
		case 'L':
			res += 50
		case 'X':
			res += 10
		case 'V':
			res += 5
		case 'I':
			res += 1
		default:
			res = 0
			return
		}
	}

	res -= (strings.Count(s, "IV")+strings.Count(s, "IX"))*2 +
		(strings.Count(s, "XL")+strings.Count(s, "XC"))*20 +
		(strings.Count(s, "CD")+strings.Count(s, "CM"))*200

	return
}

func RunRoman2int() {
	fmt.Println(romanToIntOld("MCMXCIV"))
}
