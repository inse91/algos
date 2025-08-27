package _3110_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/harshad-number

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "hello",
			res:  13,
		},
		{
			name: "2",
			s:    "zaz",
			res:  50,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, scoreOfString(c.s))
		})
	}
}

func scoreOfString(s string) int {
	var res int
	for i := 1; i < len(s); i++ {
		cur := s[i]
		prev := s[i-1]
		res += diff(cur, prev)
	}

	return res
}

func diff(a, b byte) int {
	if a > b {
		return int(a - b)
	}

	return int(b - a)
}
