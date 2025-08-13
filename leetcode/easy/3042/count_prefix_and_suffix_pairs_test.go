package _3042_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   int
	}{
		{
			name:  "1",
			words: []string{"a", "aba", "ababa", "aa"},
			res:   4,
		},
		{
			name:  "2",
			words: []string{"pa", "papa", "ma", "mama"},
			res:   2,
		},
		{
			name:  "3",
			words: []string{"abab", "ab"},
			res:   0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPrefixSuffixPairs(c.words))
		})
	}
}

func countPrefixSuffixPairs(words []string) int {
	var res int
	for i, pfxOrSfx := range words {
		for _, word := range words[i+1:] {
			if !strings.HasPrefix(word, pfxOrSfx) ||
				!strings.HasSuffix(word, pfxOrSfx) {
				continue
			}

			res++
		}
	}

	return res
}
