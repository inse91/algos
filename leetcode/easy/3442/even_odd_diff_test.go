package _3442_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "aaaaabbc",
			res:  3,
		},
		{
			name: "2",
			s:    "abcabcab",
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxDifference(c.s))
		})
	}
}

func maxDifference(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	var (
		mx = 0
		mn = len(s)
	)
	for _, freq := range m {
		switch freq % 2 {
		case 0: // even
			mn = min(mn, freq)
		case 1: // odd
			mx = max(mx, freq)
		}
	}

	return mx - mn
}
