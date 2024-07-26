package _1160_test

import (
	"github.com/stretchr/testify/assert"
	"maps"
	"testing"
)

// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		chars string
		res   int
	}{
		{
			name:  "1",
			words: []string{"cat", "bt", "hat", "tree"},
			chars: "atach",
			res:   6,
		},
		{
			name:  "2",
			words: []string{"hello", "world", "leetcode"},
			chars: "welldonehoneyr",
			res:   10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countCharacters(c.words, c.chars))
		})
	}
}

func countCharacters(words []string, chars string) int {
	charMap := make(map[rune]int, len(chars))
	for _, r := range chars {
		charMap[r]++
	}

	var res int
	for _, w := range words {
		if !canBeFormed(charMap, w) {
			continue
		}

		res += len(w)
	}

	return res
}

func canBeFormed(m map[rune]int, word string) bool {
	chars := maps.Clone(m)

	for _, r := range word {
		v, ok := chars[r]
		if !ok || v == 0 {
			return false
		}

		chars[r]--
	}

	return true
}
