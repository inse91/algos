package _2068_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent

func Test(t *testing.T) {
	cases := []struct {
		name string
		w1   string
		w2   string
		res  bool
	}{
		{
			name: "1",
			w1:   "aaaa",
			w2:   "bccb",
			res:  false,
		},
		{
			name: "2",
			w1:   "abcdeef",
			w2:   "abaaacc",
			res:  true,
		},
		{
			name: "3",
			w1:   "cccddabba",
			w2:   "babababab",
			res:  true,
		},
		{
			name: "189",
			w1:   "zzzyyy",
			w2:   "iiiiii",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkAlmostEquivalent(c.w1, c.w2))
		})
	}
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	m1 := make(map[rune]int, len(word1))
	m2 := make(map[rune]int, len(word1))

	for i := 0; i < len(word1); i++ {
		m1[rune(word1[i])]++
		m2[rune(word2[i])]++
	}

	for k, v := range m1 {
		diff := m2[k] - v
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 {
			return false
		}

		delete(m2, k)
	}

	for k, v := range m2 {
		diff := m1[k] - v
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 {
			return false
		}
	}

	return true
}
