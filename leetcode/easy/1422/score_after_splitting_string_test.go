package _1422_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-score-after-splitting-a-string/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "011101",
			res:  5,
		},
		{
			name: "2",
			s:    "00111",
			res:  5,
		},
		{
			name: "3",
			s:    "1111",
			res:  3,
		},
		{
			name: "96",
			s:    "00",
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxScore(c.s))
		})
	}
}

func maxScore(s string) int {
	var zeroes int
	ones := strings.Count(s, "1")
	maxx := -1

	for _, r := range s[:len(s)-1] {
		if r == '0' {
			zeroes++
		} else { // r == '1'
			ones--
		}

		maxx = max(maxx, zeroes+ones)
	}

	return maxx
}
