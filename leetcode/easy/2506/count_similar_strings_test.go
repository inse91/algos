package _2506_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-pairs-of-similar-strings

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   int
	}{
		{
			name:  "1",
			words: []string{"aba", "aabb", "abcd", "bac", "aabc"},
			res:   2,
		},
		{
			name:  "2",
			words: []string{"aabb", "ab", "ba"},
			res:   3,
		},
		{

			name:  "3",
			words: []string{"nba", "cba", "dba"},
			res:   0,
		},
		{
			name: "4",
			words: []string{
				"dcedceadceceaeddcedc", "dddcebcedcdbaeaaaeab", "eecbeddbddeadcbbbdbb",
				"decbcbebbddceacdeadd", "ccbddbaedcadedbcaaae", "dddcaadaceaedcdceedd",
				"bbeddbcbbccddcaceeea", "bdabacbbdadabbbddaea",
			},
			res: 16,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, similarPairs(c.words))
		})
	}
}

func similarPairs(words []string) int {
	type wordMap [26]bool
	m := make(map[wordMap]int, len(words))
	for _, word := range words {
		var wm wordMap
		for _, r := range word {
			wm[r-'a'] = true
		}
		m[wm]++
	}

	var res int
	for _, c := range m {
		res += c * (c - 1) / 2
	}

	return res
}
