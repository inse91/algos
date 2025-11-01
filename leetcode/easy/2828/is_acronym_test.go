package _2678_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		s     string
		res   bool
	}{
		{
			name:  "1",
			words: []string{"alice", "bob", "charlie"},
			s:     "abc",
			res:   true,
		},
		{
			name:  "2",
			words: []string{"an", "apple"},
			s:     "a",
			res:   false,
		},
		{
			name:  "3",
			words: []string{"never", "gonna", "give", "up", "on", "you"},
			s:     "ngguoy",
			res:   true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isAcronym(c.words, c.s))
		})
	}
}

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}

	return true
}
