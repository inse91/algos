package _1961_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array

func Test(t *testing.T) {
	cases := []struct {
		name  string
		s     string
		words []string
		res   bool
	}{
		{
			name:  "1",
			s:     "iloveleetcode",
			words: []string{"i", "love", "leetcode", "apples"},
			res:   true,
		},
		{
			name:  "2",
			s:     "iloveleetcode",
			words: []string{"apples", "i", "love", "leetcode"},
			res:   false,
		},
		{
			name:  "340",
			s:     "a",
			words: []string{"aa", "aaaa", "banana"},
			res:   false,
		},
		{
			name:  "347",
			s:     "ccccccccc",
			words: []string{"c", "cc"},
			res:   false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isPrefixString(c.s, c.words))
		})
	}
}

func isPrefixString(s string, words []string) bool {
	//return strings.HasPrefix(strings.Join(words, ""), s)
	var sb strings.Builder
	var gap int
	for _, w := range words {
		sb.WriteString(w)
		cand := sb.String()

		if len(s) < len(cand) {
			return false
		}

		gap += len(w)
		ss := s[:gap]
		if cand != ss {
			return false
		}

		if gap == len(s) {
			return true
		}
	}

	return false
}
