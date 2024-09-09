package _1332_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-palindromic-subsequences/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "ababa",
			res:  1,
		},
		{
			name: "2",
			s:    "abb",
			res:  2,
		},
		{
			name: "3",
			s:    "baabb",
			res:  2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, removePalindromeSub(c.s))
		})
	}
}

func removePalindromeSub(s string) int {
	if isPal(s) {
		return 1
	}

	return 2
}

func isPal(s string) bool {
	bts := []byte(s)
	slices.Reverse(bts)

	return string(bts) == s
}
