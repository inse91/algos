package _8_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
	"unicode"
)

func TestAtoi(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1_1",
			s:    "42",
			res:  42,
		},
		{
			name: "1_2",
			s:    "f42",
			res:  0,
		},
		{
			name: "1_3",
			s:    "+42",
			res:  42,
		},
		{
			name: "2",
			s:    "   -42",
			res:  -42,
		},
		{
			name: "3",
			s:    "4193 with words",
			res:  4193,
		},
		{
			name: "4",
			s:    "words and 987",
			res:  0,
		},
		{
			name: "5_1",
			s:    "+-12",
			res:  0,
		},
		{
			name: "5_2",
			s:    "-+12",
			res:  0,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := myAtoi(c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

const minus = rune('-')
const plus = rune('+')

func isMinus(r rune) bool {
	return r == minus
}

func isPlus(r rune) bool {
	return r == plus
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var isNegative bool
	var hasMetAnyLeadingDigit bool

	var start int
	var end int

	for i, r := range s {
		curS := string(r)
		_ = curS
		if !hasMetAnyLeadingDigit {
			if isMinus(r) {
				if i < len(s)-1 && unicode.IsDigit(rune(s[i+1])) {
					isNegative = true
					continue
				}
				return 0
			}
			if !unicode.IsDigit(r) {
				if unicode.IsSpace(r) {
					continue
				}
				if isPlus(r) && i < len(s)-1 && unicode.IsDigit(rune(s[i+1])) {
					continue
				}
				return 0
			}
			hasMetAnyLeadingDigit = true
			start = i
			continue
		}

		if unicode.IsDigit(r) {
			continue
		}
		end = i
		break
	}
	if end == 0 {
		end = len(s)
	}
	s = s[start:end]
	num, _ := strconv.Atoi(s)
	if isNegative {
		return max(-num, math.MinInt32)
	}
	return min(num, math.MaxInt32)
}
