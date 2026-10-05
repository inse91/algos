package _3438_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-valid-pair-of-adjacent-digits-in-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "2523533",
			res:  "23",
		},
		{
			name: "2",
			s:    "221",
			res:  "21",
		},
		{
			name: "3",
			s:    "22",
			res:  "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findValidPair(c.s))
		})
	}
}
func findValidPair(s string) string {
	entries := make(map[rune]int, len(s)/3)

	for _, r := range s {
		entries[r]++
	}

	for i, r := range s {
		if i == len(s)-1 {
			continue
		}

		nextRune := rune(s[i+1])
		if r == nextRune {
			continue
		}

		dig := int(r - '0')
		nextDig := int(nextRune - '0')
		if entries[r] != dig || entries[nextRune] != nextDig {
			continue
		}

		return string(r) + string(nextRune)
	}

	return ""
}
