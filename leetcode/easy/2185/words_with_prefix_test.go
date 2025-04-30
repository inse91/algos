package _2185_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/counting-words-with-a-given-prefix/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		words  []string
		prefix string
		res    int
	}{
		{
			name:   "1",
			words:  []string{"pay", "attention", "practice", "attend"},
			prefix: "at",
			res:    2,
		},
		{
			name:   "2",
			words:  []string{"leetcode", "win", "loops", "success"},
			prefix: "code",
			res:    0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, prefixCount(c.words, c.prefix))
		})
	}
}

func prefixCount(words []string, pref string) int {
	var res int
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			res += 1
		}
	}

	return res
}
