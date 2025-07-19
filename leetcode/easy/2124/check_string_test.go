package _2124_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "aaabbb",
			res:  true,
		},
		{
			name: "2",
			s:    "abab",
			res:  false,
		},
		{
			name: "3",
			s:    "bbb",
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkString(c.s))
		})
	}
}

func checkString(s string) bool {
	var gotB bool
	for _, r := range s {
		if r == 'b' {
			gotB = true
			continue
		}

		// r == 'a'
		if gotB {
			return false
		}

	}

	return true
}
