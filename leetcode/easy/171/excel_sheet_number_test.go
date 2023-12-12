package _171

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func TestExcelSheetNumber(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		res  int
	}{
		{
			name: "1",
			in:   "A",
			res:  1,
		},
		{
			name: "2",
			in:   "AB",
			res:  28,
		},
		{
			name: "3",
			in:   "ZY",
			res:  701,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := titleToNumber(tc.in)
			assert.Equal(t, tc.res, res)
		})
	}
}

func titleToNumber(columnTitle string) int {
	m := letterMap()
	sum := 0
	//columnTitle = stringsReverse(columnTitle)
	//for i, r := range columnTitle {
	c := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		letter := string(columnTitle[i])
		v := m[letter]
		sum += v * int(math.Pow(26, float64(c)))
		c++
	}
	return sum
}

func letterMap() map[string]int {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[string]int, len(letters))

	// Split the letters into a slice of strings
	letterSlice := strings.Split(letters, "")

	// Add entries to the map
	for i, letter := range letterSlice {
		m[letter] = i + 1
	}
	return m
}
