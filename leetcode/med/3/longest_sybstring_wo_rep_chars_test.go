package _3_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstrWithoutRepChars(t *testing.T) {
	cases := []struct {
		name string
		in   string
		res  int
	}{
		{
			name: "1",
			in:   "abcabcbb",
			res:  3,
		},
		{
			name: "2",
			in:   "bbbbb",
			res:  1,
		},
		{
			name: "3",
			in:   "pwwkew",
			res:  3,
		},
		{
			name: "4",
			in:   "dvdf",
			res:  3,
		},
		{
			name: "5",
			in:   " ",
			res:  1,
		},
		{
			name: "6",
			in:   "a",
			res:  1,
		},
		{
			name: "7",
			in:   "aab",
			res:  2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := lengthOfLongestSubstring(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	charMap := make(map[byte]int, 16)
	var subLenMax int
	var subLen int
	var done bool

	var start int
outer:
	for {
		for i := start; i < len(s); i++ {
			b := s[i]
			pos, ok := charMap[b]
			if !ok {
				charMap[b] = i
				subLen++
				continue
			}

			done = true
			start = pos + 1
			charMap = make(map[byte]int, 16)
			subLenMax = max(subLenMax, subLen)
			subLen = 0
			if i == len(s)-1 {
				break outer
			}
			break
		}
	}
	if !done {
		subLenMax = subLen
	}

	return subLenMax
}
