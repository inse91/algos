package _1844_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/replace-all-digits-with-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "a1c1e1",
			res:  "abcdef",
		},
		{
			name: "2",
			s:    "a1b2c3d4e",
			res:  "abbdcfdhe",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, replaceDigits(c.s))
		})
	}
}

func replaceDigits(s string) string {
	var res strings.Builder
	res.Grow(len(s))
	for i, r := range s {
		if i%2 == 0 || !unicode.IsDigit(r) {
			res.WriteRune(r)

			continue
		}

		r += rune(s[i-1]) - 48
		res.WriteRune(r)
	}

	return res.String()
}
