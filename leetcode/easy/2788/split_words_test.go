package _2788_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/split-strings-by-separator

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		sep   byte
		res   []string
	}{
		{
			name:  "1",
			words: []string{"one.two.three", "four.five", "six"},
			sep:   '.',
			res:   []string{"one", "two", "three", "four", "five", "six"},
		},
		{
			name:  "2",
			words: []string{"$easy$", "$problem$"},
			sep:   '$',
			res:   []string{"easy", "problem"},
		},
		{
			name:  "3",
			words: []string{"|||"},
			sep:   '|',
			res:   []string{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, splitWordsBySeparator(c.words, c.sep))
		})
	}
}

func splitWordsBySeparator(words []string, sep byte) []string {
	res := make([]string, 0, len(words))
	for _, w := range words {
		parts := strings.Split(w, string(sep))
		for _, s := range parts {
			if s == "" {
				continue
			}

			res = append(res, s)
		}
	}

	return res
}
