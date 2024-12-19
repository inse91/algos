package _1614_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "(1+(2*3)+((8)/4))+1",
			res:  3,
		},
		{
			name: "2",
			s:    "(1+(2*3)+((8)/4))+1",
			res:  3,
		},
		{
			name: "3",
			s:    "()(())((()()))",
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxDepth(c.s))
		})
	}
}

func maxDepth(s string) int {
	var current int
	var mx int
	for _, r := range s {
		switch r {
		case '(':
			current++
		case ')':
			current--
		}

		mx = max(mx, current)
	}

	return mx
}
