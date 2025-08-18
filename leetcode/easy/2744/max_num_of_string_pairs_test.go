package _2744_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-maximum-number-of-string-pairs

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   int
	}{
		{
			name:  "1",
			words: []string{"cd", "ac", "dc", "ca", "zz"},
			res:   2,
		},
		{
			name:  "2",
			words: []string{"ab", "ba", "cc"},
			res:   1,
		},
		{
			name:  "3",
			words: []string{"aa", "ab"},
			res:   0,
		},
		{
			name:  "11",
			words: []string{"cd", "ac", "dc", "ca", "zz", "ca", "ac"},
			res:   3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumNumberOfStringPairs(c.words))
		})
	}
}

func maximumNumberOfStringPairs(words []string) int {
	s := make(map[string]struct{})
	var res int
	for _, w := range words {
		if _, ok := s[w]; ok {
			res++
			delete(s, w)
			continue
		}

		w = reverse(w)
		s[w] = struct{}{}
	}

	return res
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
