package _2309_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "lEeTcOdE",
			res:  "E",
		},
		{
			name: "2",
			s:    "arRAzFif",
			res:  "R",
		},
		{
			name: "3",
			s:    "AbCdEfGhIjK",
			res:  "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, greatestLetter(c.s))
		})
	}
}

func greatestLetter(s string) string {
	uc := make(map[rune]struct{})
	lc := make(map[rune]struct{})
	var res rune
	for _, r := range s {
		if unicode.IsUpper(r) {
			uc[r] = struct{}{}
			_, ok := lc[r]
			if ok && r > res {
				res = r
			}

			continue
		}

		// is lower
		r = unicode.ToUpper(r)
		lc[r] = struct{}{}
		_, ok := uc[r]
		if ok && r > res {
			res = r
		}
	}

	if res == 0 {
		return ""
	}

	return string(res)
}
