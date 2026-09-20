package _3120_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-the-number-of-special-characters-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "aaAbcBC",
			res:  3,
		},
		{
			name: "2",
			s:    "abc",
			res:  0,
		},
		{
			name: "3",
			s:    "abBCab",
			res:  1,
		},
		{
			name: "627",
			s:    "aAA",
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfSpecialChars(c.s))
		})
	}
}

func numberOfSpecialChars(word string) int {
	var (
		set = make(map[string]struct{ l, u, ok bool }, len(word))
		res int
	)

	for _, ch := range word {
		isLc := unicode.IsLower(ch)
		key := string(unicode.ToLower(ch))
		v, ok := set[key]
		if !ok {
			set[key] = struct{ l, u, ok bool }{l: isLc, u: !isLc}
			continue
		}

		if v.ok {
			continue
		}

		if isLc {
			v.l = true
		} else {
			v.u = true
		}

		if !v.l || !v.u {
			set[key] = v
			continue
		}

		res++
		v.ok = true
		set[key] = v
	}

	return res
}
