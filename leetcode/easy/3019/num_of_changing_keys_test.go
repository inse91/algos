package _3019_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-changing-keys

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "aAbBcC",
			res:  2,
		},
		{
			name: "2",
			s:    "AaAaAaaA",
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countKeyChanges(c.s))
		})
	}
}

func countKeyChanges(s string) int {
	var (
		c    int  = -1
		prev rune = ' '
	)
	for _, char := range s {
		char = unicode.ToLower(char)
		if char == prev {
			continue
		}

		c++
		prev = char
	}
	return c
}
