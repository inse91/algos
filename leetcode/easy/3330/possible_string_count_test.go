package _3330_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-original-typed-string-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		word string
		res  int
	}{
		{
			name: "1",
			word: "abbcccc",
			res:  5,
		},
		{
			name: "2",
			word: "abcd",
			res:  1,
		}, {
			name: "3",
			word: "aaaa",
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, possibleStringCount(c.word))
		})
	}
}

func possibleStringCount(word string) int {
	res := 1
	for i := 1; i < len(word); i++ {
		cur := word[i]
		prev := word[i-1]
		if cur == prev {
			res++
		}
	}

	return res
}
