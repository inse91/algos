package _392_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//leeeeetcode

func TestIsSubseq(t *testing.T) {
	cases := []struct {
		name string
		s    string
		t    string
		res  bool
	}{
		{
			name: "1",
			s:    "abc",
			t:    "ahbgdc",
			res:  true,
		},
		{
			name: "2",
			s:    "axc",
			t:    "ahbgdc",
			res:  false,
		},
		{
			name: "3",
			s:    "ace",
			t:    "abcde",
			res:  true,
		},
		{
			name: "4",
			s:    "aec",
			t:    "abcde",
			res:  false,
		},
		{
			name: "5",
			s:    "ab",
			t:    "baab",
			res:  true,
		},
		{
			name: "6",
			s:    "abcd",
			t:    "abcd",
			res:  true,
		},
		{
			name: "7",
			s:    "leeeeetcode",
			t:    "yyyyyylyyyyyyyeyyyyyeyyeyyeyyeeeeyyyytyyyyyyyyycyyyyyyyyyyoyyyyyyyyyyydyyyyyyyyeyyyyyy",
			res:  true,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := isSubsequence(c.s, c.t)
			assert.Equal(t, c.res, res)
		})
	}
}
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j < len(t) && i < len(s) {
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == len(s)
}
