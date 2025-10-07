package _2697_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lexicographically-smallest-palindrome

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "egcfe",
			res:  "efcfe",
		},
		{
			name: "2",
			s:    "abcd",
			res:  "abba",
		},
		{
			name: "3",
			s:    "seven",
			res:  "neven",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, makeSmallestPalindrome(c.s))
		})
	}
}

func makeSmallestPalindrome(s string) string {
	bts := []byte(s)
	for i := 0; i < len(bts)/2; i++ {
		cur := bts[i]
		pair := bts[len(bts)-1-i]

		if cur == pair {
			continue
		}

		bts[i] = min(cur, pair)
		bts[len(bts)-1-i] = bts[i]
	}

	return string(bts)
}
