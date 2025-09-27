package _2942_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-words-containing-character

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		char  byte
		res   []int
	}{
		{
			name:  "1",
			words: []string{"leet", "code"},
			char:  'e',
			res:   []int{0, 1},
		},
		{
			name:  "2",
			words: []string{"abc", "bcd", "aaaa", "cbc"},
			char:  'a',
			res:   []int{0, 2},
		},
		{
			name:  "3",
			words: []string{"abc", "bcd", "aaaa", "cbc"},
			char:  'z',
			res:   []int{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, findWordsContaining(c.words, c.char))
		})
	}
}

func findWordsContaining(words []string, x byte) []int {
	var idxes []int
	for i, w := range words {
		if strings.ContainsRune(w, rune(x)) {
			idxes = append(idxes, i)
		}
	}

	return idxes
}
