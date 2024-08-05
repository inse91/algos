package _1221_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/split-a-string-in-balanced-strings/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "RLRRLLRLRL",
			res:  4,
		},
		{
			name: "2",
			s:    "RLRRRLLRLL",
			res:  2,
		},
		{
			name: "3",
			s:    "LLLLRRRR",
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, balancedStringSplit(c.s))
		})
	}
}

func balancedStringSplit(s string) int {
	var balance, res int
	for _, r := range s {
		if r == 'R' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			res++
		}
	}

	return res
}
