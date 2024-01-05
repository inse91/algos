package _290_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWordPattern(t *testing.T) {
	cases := []struct {
		name    string
		pattern string
		s       string
		res     bool
	}{
		{
			name:    "1",
			pattern: "abba",
			s:       "dog cat cat dog",
			res:     true,
		},
		{
			name:    "2",
			pattern: "abba",
			s:       "dog cat cat fish",
			res:     false,
		},
		{
			name:    "3",
			pattern: "aaaa",
			s:       "dog cat cat dog",
			res:     false,
		},
		{
			name:    "4",
			pattern: "abba",
			s:       "dog dog dog dog",
			res:     false,
		},
		{
			name:    "5",
			pattern: "aaa",
			s:       "dog dog dog dog dog",
			res:     false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			res := wordPattern(c.pattern, c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

func wordPattern(pattern string, s string) bool {
	m := make(map[rune]string, 16)
	sl := strings.Split(s, " ")
	if len(pattern) != len(sl) {
		return false
	}

	wordSet := make(map[string]struct{}, 16)

	var val string
	var ok bool
	for i, r := range pattern {
		word := sl[i]
		if val, ok = m[r]; !ok {
			if _, ok = wordSet[word]; ok {
				return false
			}
			m[r] = word
			wordSet[word] = struct{}{}
			continue
		}

		if val == word {
			continue
		}

		return false
	}

	return true
}
