package _1974_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		word string
		res  int
	}{
		{
			name: "1",
			word: "abc",
			res:  5,
		},
		{
			name: "2",
			word: "bza",
			res:  7,
		},
		{
			name: "3",
			word: "zjpc",
			res:  34,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minTimeToType(c.word))
		})
	}
}

func minTimeToType(word string) int {
	res := len(word)

	// first letter
	gap := word[0] - 'a'
	if gap > 13 {
		gap = 26 - gap
	}
	res += int(gap)

	for i := 1; i < len(word); i++ {
		if word[i] >= word[i-1] {
			gap = word[i] - word[i-1]
		} else {
			gap = word[i-1] - word[i]
		}

		if gap > 13 {
			gap = 26 - gap
		}
		res += int(gap)
	}

	return res
}
