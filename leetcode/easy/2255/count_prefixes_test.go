package _2255_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-prefixes-of-a-given-string

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		prfx  string
		res   int
	}{
		{
			name:  "1",
			words: []string{"a", "b", "c", "ab", "bc", "abc"},
			prfx:  "abc",
			res:   3,
		},
		{
			name:  "2",
			words: []string{"a", "a"},
			prfx:  "aa",
			res:   2,
		},
		{
			name: "115",
			words: []string{
				"feh", "w", "w", "lwd", "c", "s", "vk", "zwlv", "n", "w", "sw", "qrd", "w", "w", "mqe", "w",
				"w", "w", "gb", "w", "qy", "xs", "br", "w", "rypg", "wh", "g", "w", "w", "fh", "w", "w", "sccy",
			},
			prfx: "w",
			res:  14,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, countPrefixes(c.words, c.prfx))
		})
	}
}

func countPrefixes(words []string, prfx string) int {
	var res int
	for _, w := range words {
		c := prfx
		if len(prfx) > len(w) {
			c = c[:len(w)]
		}
		if w == c {
			res++
		}
	}

	return res
}
