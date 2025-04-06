package _1941_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		word string
		res  bool
	}{
		{
			name: "1",
			word: "abacbc",
			res:  true,
		},
		{
			name: "2",
			word: "aaabb",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, areOccurrencesEqual(c.word))
		})
	}
}

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	var q int
	for i, r := range s {
		m[r]++
		if i == len(s)-1 {
			q = m[r]
		}
	}

	for _, v := range m {
		if v == q {
			continue
		}

		return false
	}

	return true
}
