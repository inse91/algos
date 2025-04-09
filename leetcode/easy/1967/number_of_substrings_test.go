package _1967_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/description/

func Test(t *testing.T) {
	cases := []struct {
		name     string
		patterns []string
		word     string
		res      int
	}{
		{
			name:     "1",
			word:     "abc",
			patterns: []string{"a", "abc", "bc", "d"},
			res:      3,
		},
		{
			name:     "2",
			word:     "aaaaabbbbb",
			patterns: []string{"a", "b", "c"},
			res:      2,
		},
		{
			name:     "3",
			word:     "ab",
			patterns: []string{"a", "a", "a"},
			res:      3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numOfStrings(c.patterns, c.word))
		})
	}
}

func numOfStrings(patterns []string, word string) int {
	var res int
	for _, p := range patterns {
		if strings.Contains(word, p) {
			res++
		}
	}

	return res
}
