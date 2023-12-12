package _168

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestExcelTitle(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		res  string
	}{
		{
			name: "1",
			in:   1,
			res:  "A",
		},
		{
			name: "2.1",
			in:   28,
			res:  "AB",
		},
		{
			name: "2.2",
			in:   29,
			res:  "AC",
		},
		{
			name: "2.3",
			in:   53,
			res:  "BA",
		},
		{
			name: "2.4",
			in:   82,
			res:  "CD",
		},
		{
			name: "3",
			in:   701,
			res:  "ZY",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := convertToTitle(tc.in)
			assert.Equal(t, tc.res, res)
		})

	}
}

func convertToTitle(columnNumber int) string {
	m := letterMap()

	sb := strings.Builder{}
	for x := columnNumber; x > 0; x = x / 26 {
		r := x % 26
		if r == 0 {
			x--
		}
		sb.WriteString(m[r])
	}
	s := sb.String()
	return stringsReverse(s)
}

func letterMap() map[int]string {
	letters := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	m := make(map[int]string, len(letters))

	// Split the letters into a slice of strings
	letterSlice := strings.Split(letters, "")

	// Add entries to the map
	for i, letter := range letterSlice {
		m[i] = letter
	}
	return m
}

func stringsReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
