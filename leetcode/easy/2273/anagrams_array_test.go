package _2273_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   []string
	}{
		{
			name:  "1",
			words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			res:   []string{"abba", "cd"},
		},
		{
			name:  "2",
			words: []string{"a", "b", "c", "d", "e"},
			res:   []string{"a", "b", "c", "d", "e"},
		},
		{
			name:  "195",
			words: []string{"a", "b", "a"},
			res:   []string{"a", "b", "a"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, removeAnagrams(c.words))
		})
	}
}

func removeAnagrams(words []string) []string {
	var (
		prevKey [26]int
		res     []string
	)

	for _, w := range words {
		key := wordMap(w)
		if key == prevKey {
			continue
		}

		prevKey = key
		res = append(res, w)
	}

	return res
}

func wordMap(w string) [26]int {
	res := [26]int{}
	for _, r := range w {
		idx := int(r - 'a')
		res[idx] += 1
	}

	return res
}
