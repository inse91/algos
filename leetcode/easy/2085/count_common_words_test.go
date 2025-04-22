package _2085_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-common-words-with-one-occurrence/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		words1 []string
		words2 []string
		res    int
	}{
		{
			name:   "1",
			words1: []string{"leetcode", "is", "amazing", "as", "is"},
			words2: []string{"amazing", "leetcode", "is"},
			res:    2,
		},
		{
			name:   "2",
			words1: []string{"b", "bb", "bbb"},
			words2: []string{"a", "aa", "aaa"},
			res:    0,
		},
		{
			name:   "3",
			words1: []string{"a", "ab"},
			words2: []string{"a", "a", "a", "ab"},
			res:    1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countWords(c.words1, c.words2))
		})
	}
}

func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]struct{}, len(words1))
	banned := make(map[string]struct{})
	for _, w := range words1 {
		if _, ok := banned[w]; ok {
			delete(m1, w)
			continue
		}

		if _, ok := m1[w]; ok {
			banned[w] = struct{}{}
			delete(m1, w)

			continue
		}

		m1[w] = struct{}{}
	}

	m2 := make(map[string]struct{}, len(words1))
	for _, w := range words2 {
		if _, ok := m1[w]; !ok {
			continue
		}

		if _, ok := banned[w]; ok {
			delete(m1, w)

			continue
		}

		if _, ok := m2[w]; ok {
			banned[w] = struct{}{}
			delete(m2, w)

			continue
		}

		m2[w] = struct{}{}
	}

	return len(m2)
}
