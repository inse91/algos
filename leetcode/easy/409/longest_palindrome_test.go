package _09__test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  int
	}{
		{
			name: "1",
			in:   "babad",
			res:  5,
		},
		{
			name: "2",
			in:   "cbbd",
			res:  3,
		},
		{
			name: "3",
			in:   "a",
			res:  1,
		},
		{
			name: "4",
			in:   "abccccdd",
			res:  7,
		},
		{
			name: "5",
			in:   "abaccccdd",
			res:  9,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := longestPalindrome(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func longestPalindrome(s string) int {
	var maxLenPal int

	m := make(map[byte]int, 32)
	var hasAtLeastOneSingleChar bool

	for i := range s {
		b := s[i]
		if _, ok := m[b]; !ok {
			m[b] = 1
			continue
		}
		m[b]++
	}

	for _, v := range m {
		if v%2 == 0 {
			maxLenPal += v
			continue
		}
		hasAtLeastOneSingleChar = true
		maxLenPal += v - 1
	}

	if hasAtLeastOneSingleChar {
		maxLenPal++
	}

	return maxLenPal
}
