package _1684_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		allowed string
		words   []string
		res     int
	}{
		{
			name:    "1",
			allowed: "ab",
			words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			res:     2,
		},
		{
			name:    "2",
			allowed: "abc",
			words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			res:     7,
		},
		{
			name:    "3",
			allowed: "cad",
			words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			res:     4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countConsistentStrings(c.allowed, c.words))
		})
	}
}

func countConsistentStrings(allowed string, words []string) int {
	allowedMap := make(map[byte]struct{}, len(allowed))
	for _, r := range allowed {
		allowedMap[byte(r)] = struct{}{}
	}

	var c int
	for _, w := range words {
		isValid := true
		for _, r := range w {
			_, ok := allowedMap[byte(r)]
			if ok {
				continue
			}

			isValid = false
			break
		}

		if isValid {
			c++
		}
	}

	return c
}
