package _2451_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/odd-string-difference

func Test(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		res   string
	}{
		{
			name:  "1",
			words: []string{"adc", "wzy", "abc"},
			res:   "abc",
		},
		{
			name:  "2",
			words: []string{"aaa", "bob", "ccc", "ddd"},
			res:   "bob",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, oddString(c.words))
		})
	}
}

func oddString(words []string) string {
	if len(words) <= 2 {
		panic("not enough words")
	}

	s0 := diff(words[0])
	s1 := diff(words[1])

	if slices.Equal(s0, s1) {
		for i := 2; i < len(words); i++ {
			cur := diff(words[i])
			if slices.Equal(cur, s0) {
				continue
			}

			return words[i]
		}
	}

	// s0 or s1 is the one we need
	s2 := diff(words[2])
	if slices.Equal(s0, s2) {
		return words[1]
	}

	return words[0]
}

func diff(s string) []rune {
	res := make([]rune, 0, len(s)-1)
	for i, r := range s {
		if i == 0 {
			continue
		}

		prev := rune(s[i-1])
		res = append(res, r-prev)
	}

	return res
}
