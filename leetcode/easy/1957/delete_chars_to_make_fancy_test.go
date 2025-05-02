package _1957_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/delete-characters-to-make-fancy-string/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "leeetcode",
			res:  "leetcode",
		},
		{
			name: "2",
			s:    "aaabaaaa",
			res:  "aabaa",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, makeFancyString(c.s))
		})
	}
}

func makeFancyString(s string) string {
	var sb strings.Builder
	for _, c := range s {
		cur := sb.String()
		if len(cur) < 2 || rune(cur[len(cur)-1]) != c || rune(cur[len(cur)-2]) != c {
			sb.WriteRune(c)
		}

	}

	return sb.String()
}
