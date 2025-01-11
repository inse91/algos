package _1704_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/determine-if-string-halves-are-alike/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  bool
	}{
		{
			name: "1",
			s:    "book",
			res:  true,
		},
		{
			name: "2",
			s:    "textbook",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, halvesAreAlike(c.s))
		})
	}
}

func halvesAreAlike(s string) bool {
	center := len(s) / 2

	var c1 int
	for _, r := range s[:center] {
		if isVowel(byte(r)) {
			c1++
		}
	}

	var c2 int
	for _, r := range s[center:] {
		if isVowel(byte(r)) {
			c2++
		}
	}

	return c1 == c2
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
