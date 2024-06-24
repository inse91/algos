package _1021_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://leetcode.com/problems/remove-outermost-parentheses/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "(()())(())",
			res:  "()()()",
		},
		{
			name: "2",
			s:    "(()())(())(()(()))",
			res:  "()()()()(())",
		},
		{
			name: "3",
			s:    "()()",
			res:  "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, removeOuterParentheses(c.s))
		})
	}
}

func removeOuterParentheses(s string) string {
	var res strings.Builder
	var c int
	for i, r := range s {
		_ = i
		if c == 1 && r == ')' {
			c--
			continue
		}
		if c == 0 {
			if r == '(' {
				c++
			}
			continue
		}
		res.WriteRune(r)
		if r == '(' {
			c++
		}
		if r == ')' {
			c--
		}
	}

	return res.String()
}
