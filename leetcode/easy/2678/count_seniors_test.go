package _2678_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/description/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		details []string
		res     int
	}{
		{
			name:    "1",
			details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			res:     2,
		},
		{
			name:    "2",
			details: []string{"1313579440F2036", "2921522980M5644"},
			res:     0,
		},
		{
			name:    "175",
			details: []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"},
			res:     2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countSeniors(c.details))
		})
	}
}

func countSeniors(details []string) int {
	var res int
	for _, d := range details {
		if d[11] >= '7' {
			res++
			continue
		}

		if d[11] != '6' {
			continue
		}

		if d[12] != '0' {
			res++
		}

	}

	return res
}
