package _1763_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-nice-substring/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "YazaAay",
			res:  "aAa",
		},
		{
			name: "2",
			s:    "Bb",
			res:  "Bb",
		},
		{
			name: "3",
			s:    "c",
			res:  "",
		},
		{
			name: "93",
			s:    "BebjJE",
			res:  "BebjJE",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, longestNiceSubstring(c.s))
		})
	}
}

func longestNiceSubstring(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			candidate := s[i : j+1]
			if checkIfNice(candidate) && len(candidate) > len(res) {
				res = candidate
			}
		}
	}

	return res
}

func checkIfNice(s string) bool {
	if len(s) <= 1 {
		return false
	}

	m := make(map[rune]struct{}, len(s))
	for _, r := range s {
		m[r] = struct{}{}
	}

	for r := range m {
		if unicode.IsUpper(r) {
			lower := unicode.ToLower(r)
			if _, ok := m[lower]; !ok {
				return false
			}
		}

		if unicode.IsLower(r) {
			upper := unicode.ToUpper(r)
			if _, ok := m[upper]; !ok {
				return false
			}
		}
	}

	return true
}
