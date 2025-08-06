package _2716_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-pairs-that-form-a-complete-day-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "aaabc",
			res:  3,
		},
		{
			name: "2",
			s:    "cbbd",
			res:  3,
		},
		{
			name: "3",
			s:    "baadccab",
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimizedStringLength(c.s))
		})
	}
}

func minimizedStringLength(s string) int {
	m := make(map[rune]struct{}, 26)
	for _, r := range s {
		m[r] = struct{}{}
	}

	return len(m)
}
