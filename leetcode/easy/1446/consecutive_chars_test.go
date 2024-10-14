package _1446_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/consecutive-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "leetcode",
			res:  2,
		},
		{
			name: "2",
			s:    "abbcccddddeeeeedcba",
			res:  5,
		},
		{
			name: "3",
			s:    "aa",
			res:  2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxPower(c.s))
		})
	}
}

func maxPower(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var res int
	acc := 1
	for i := range s {
		if i == 0 {
			continue
		}

		if s[i] == s[i-1] {
			acc++
			continue
		}

		res = max(res, acc)
		acc = 1
	}

	res = max(res, acc)

	return res
}
