package _1903_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-odd-number-in-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "52",
			res:  "5",
		},
		{
			name: "2",
			s:    "4206",
			res:  "",
		},
		{
			name: "3",
			s:    "35427",
			res:  "35427",
		},
		{
			name: "4",
			s:    "35426",
			res:  "35",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestOddNumber(c.s))
		})
	}
}

func largestOddNumber(s string) string {
	for i := len(s) - 1; i > -1; i-- {
		b := s[i]
		if isOdd(b) {
			return s[:i+1]
		}
	}

	return ""
}

func isOdd(r byte) bool {
	switch r {
	case '1', '3', '5', '7', '9':
		return true
	}

	return false
}
