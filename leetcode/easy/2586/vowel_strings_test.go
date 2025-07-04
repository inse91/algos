package _2586_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		left  int
		right int
		res   int
	}{
		{
			name:  "1",
			words: []string{"are", "amy", "u"},
			left:  0,
			right: 2,
			res:   2,
		},
		{
			name:  "2",
			words: []string{"hey", "aeo", "mu", "ooo", "artro"},
			left:  1,
			right: 4,
			res:   3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, vowelStrings(c.words, c.left, c.right))
		})
	}
}

func vowelStrings(words []string, left int, right int) int {
	var res int
	for i := left; i <= right; i++ {
		word := words[i]

		firstByte := word[0]
		lastByte := word[len(word)-1]

		switch firstByte {
		case 'a', 'e', 'i', 'o', 'u':
		default:
			continue
		}

		switch lastByte {
		case 'a', 'e', 'i', 'o', 'u':
		default:
			continue
		}

		res++
	}

	return res
}
