package _2062_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-vowel-substrings-of-a-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		w    string
		res  int
	}{
		{
			name: "1",
			w:    "aeiouu",
			res:  2,
		},
		{
			name: "2",
			w:    "unicornarihan",
			res:  0,
		},
		{
			name: "3",
			w:    "cuaieuouac",
			res:  7,
		},
		{
			name: "44",
			w:    "aeiouufffaeiouu",
			res:  4,
		},
		{
			name: "57",
			w:    "aeiouufffaeiouuae",
			res:  9,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countVowelSubstrings(c.w))
		})
	}
}

func countVowelSubstrings(word string) int {
	var (
		res int
		sb  strings.Builder
	)

	for i, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			sb.WriteRune(r)
			if i != len(word)-1 {
				continue
			}
		default:
		}

		cand := sb.String()
		sb.Reset()

		valid := checkIfStringContainsAllVowels(cand)
		if !valid {
			continue
		}

		res += count(cand)
	}

	return res
}

func count(s string) int {
	var a, o, i, u, e, res int
	for j := range s {
		var gotAll bool
		for k := j; k < len(s); k++ {
			c := s[k]
			switch c {
			case 'a':
				a++
			case 'o':
				o++
			case 'u':
				u++
			case 'e':
				e++
			case 'i':
				i++
			}

			if !gotAll && a > 0 && o > 0 && e > 0 && i > 0 && u > 0 {
				gotAll = true
			}

			if gotAll {
				res++
			}
		}

		a = 0
		o = 0
		i = 0
		e = 0
		u = 0

		gotAll = false
	}

	return res
}

func checkIfStringContainsAllVowels(s string) bool {
	if len(s) < 5 {
		return false
	}

	var (
		a, e, i, o, u bool
		c             int
	)
	for _, r := range s {
		if !a && r == 'a' {
			a = true
			c++
		}
		if !e && r == 'e' {
			e = true
			c++
		}
		if !i && r == 'i' {
			i = true
			c++
		}
		if !o && r == 'o' {
			o = true
			c++
		}
		if !u && r == 'u' {
			u = true
			c++
		}
	}

	return c == 5
}
